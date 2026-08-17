# SQL Classification Rules

Classify every `.sql` file using this exact precedence order. Stop at the first matching rule.

## 1) Packages

- Match when filename contains `PKS` or `PKB`.
- Destination: `Install/DBObjects/Packages/`
- Examples:
  - `PKG_CUSTOMER_PKS.sql`
  - `INV_SYNC_PKB.sql`

## 2) Procedures

- Match when filename contains `_PRC`, `SP_`, `PROC`, or `_PROC`.
- Destination: `Install/DBObjects/Procedures/`
- Examples:
  - `CUSTOMER_PRC.sql`
  - `SP_REBUILD_INDEX.sql`
  - `INVENTORY_PROC.sql`

## 3) Types

- Match when filename contains `_TYPE` or `TYPE_`.
- Destination: `Install/DBObjects/Types/`
- Examples:
  - `ORDER_TYPE.sql`
  - `TYPE_PRICE_RULE.sql`

## 4) Tables

- Match when filename contains `_TAB` or `_TABLE`.
- Fallback match: file content begins with `CREATE TABLE` (ignoring comments/spacing).
- Destination: `Install/DBObjects/Tables/`
- Examples:
  - `CUSTOMER_TAB.sql`
  - `CREATE_SALES_TABLE.sql`
  - `new_entity.sql` (if first statement is `CREATE TABLE`)

## 5) Data

- Match when filename is exactly `execution.sql`.
- Fallback match: first non-comment SQL statement is `INSERT`, `UPDATE`, or `MERGE`.
- Destination: `Install/DBObjects/Data/`
- Examples:
  - `execution.sql`
  - `seed_master_data.sql` (starts with `INSERT`)
  - `sync_delta.sql` (starts with `MERGE`)

## 6) Unclassified

- If no prior rule matches, classify as unclassified.
- Destination: `Install/DBObjects/Unclassified/`
- Action: log warning in summary report.
- Example:
  - `misc_patch.sql`

## Numeric Prefix Rule

- Preserve existing `NN_` prefix (two digits + underscore), e.g. `07_SP_FOO.sql`.
- If missing, assign next sequential `NN_` inside the destination subfolder:
  - `SP_ALPHA.sql` -> `01_SP_ALPHA.sql`
  - `SP_BETA.sql` -> `02_SP_BETA.sql`
