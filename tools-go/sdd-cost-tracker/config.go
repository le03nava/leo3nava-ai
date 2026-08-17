package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	defaultPort = 7438
	portEnv     = "SDD_COST_PORT"
	dbPathEnv   = "SDD_COST_DB_PATH"
)

type Config struct {
	Port    int
	DBPath  string
	MCPMode bool
}

func ParseConfig(args []string) (Config, error) {
	return parseConfig(args, os.Getenv)
}

func parseConfig(args []string, getenv func(string) string) (Config, error) {
	fs := flag.NewFlagSet("sdd-cost-tracker", flag.ContinueOnError)

	portFlag := fs.Int("port", 0, "HTTP port")
	dbPathFlag := fs.String("db-path", "", "SQLite database path")
	mcpFlag := fs.Bool("mcp", false, "Run MCP stdio mode")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}

	flagSet := map[string]bool{}
	fs.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})

	port, err := resolvePort(flagSet["port"], *portFlag, getenv)
	if err != nil {
		return Config{}, err
	}

	dbPath, err := resolveDBPath(flagSet["db-path"], *dbPathFlag, getenv)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:    port,
		DBPath:  dbPath,
		MCPMode: *mcpFlag,
	}, nil
}

func resolvePort(flagProvided bool, flagValue int, getenv func(string) string) (int, error) {
	if flagProvided {
		if flagValue <= 0 {
			return 0, fmt.Errorf("--port must be > 0")
		}
		return flagValue, nil
	}

	if envPort := strings.TrimSpace(getenv(portEnv)); envPort != "" {
		parsed, err := strconv.Atoi(envPort)
		if err != nil || parsed <= 0 {
			return 0, fmt.Errorf("%s must be a positive integer", portEnv)
		}
		return parsed, nil
	}

	return defaultPort, nil
}

func resolveDBPath(flagProvided bool, flagValue string, getenv func(string) string) (string, error) {
	raw := ""

	if flagProvided {
		raw = strings.TrimSpace(flagValue)
	} else {
		raw = strings.TrimSpace(getenv(dbPathEnv))
	}

	if raw == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve user home: %w", err)
		}
		raw = filepath.Join(home, ".sdd-cost-tracker", "db.sqlite")
	}

	return expandHome(raw)
}

func expandHome(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("db path cannot be empty")
	}

	if path == "~" || strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "~\\") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve user home: %w", err)
		}
		if path == "~" {
			return home, nil
		}
		suffix := strings.TrimPrefix(strings.TrimPrefix(path, "~/"), "~\\")
		path = filepath.Join(home, suffix)
	}

	return filepath.Clean(path), nil
}
