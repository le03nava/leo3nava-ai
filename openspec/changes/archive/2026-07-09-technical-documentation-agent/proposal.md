# Proposal: Technical Documentation Agent

## Intent

Create a manual post-archive SDD utility, similar to `sdd-operational-doc`, that generates one Spanish technical documentation Markdown file from archived SDD evidence only. This addresses the need for a repeatable final technical-documentation deliverable without changing the active SDD DAG or requiring new mandatory phases.

## Scope

### In Scope
- Add `skills/sdd-technical-doc/SKILL.md` as a user-invocable archive-consuming utility.
- Add `skills/sdd-technical-doc/assets/technical-document-template.md` with the required Spanish document contract below.
- Add `agents/sdd/sdd-technical-doc.md` as the thin executor prompt.
- Update `AGENTS.md` and `README.md` to document the manual post-archive boundary.
- Add or update specs for the manual technical documentation utility.

### Out of Scope
- No changes to the SDD DAG/dependency graph, phase order, status routing, archive gates, or existing SDD flows.
- No implementation of the generated technical document itself in this change.
- No modification of `sdd-operational-doc` beyond possible docs references.

## Capabilities

### New Capabilities
- `sdd-technical-documentation-workflow`: Manual post-archive utility that generates a Spanish technical documentation Markdown file from archived SDD evidence.

### Modified Capabilities
- `sdd-execution-persistence-contracts`: Clarify that `sdd-technical-doc` is manual, post-archive, archive-consuming, and not a required DAG phase.

## Core Output Requirement

The utility MUST generate exactly one Markdown technical documentation file from an already archived SDD change. It MUST obtain everything from archived evidence and MUST NOT modify current SDD flows. If a specific point does not apply, write `No aplica.` If information is unavailable, write that the information is unavailable.

Required document structure:

```text
Identificacion del Producto —
    Tabla pipe: Producto, Dueno de aplicacion, Plataforma, Lider de Proyecto TI, Desarrollador.
    Referencias — Bullets - archivo - descripcion. Solo documentos funcionales y fuentes del desarrollo (FCTI, historia, .pks/.pkb/.sql del package/tabla creada o modificada). NO referenciar scripts del instalador.
1. Presentacion del Producto —
    Parrafo introductorio del cambio (que hace, por que se desarrollo, contexto funcional).
    1.1 Objetivo
        Una o dos lineas con el objetivo funcional del cambio.
    1.2 Alcance
        Alcance funcional + Inventario de objetos del cambio (tabla pipe obligatoria: | Tipo | Esquema.Objeto | Operacion | Descripcion funcional | donde Tipo ∈ {TABLE, INDEX, PACKAGE SPEC, PACKAGE BODY, PROCEDURE, FUNCTION, TRIGGER, JOB ControlM, VIEW, SEQUENCE, GRANT} y Operacion ∈ {CREATE, ALTER, REPLACE, DROP}). Una fila por objeto tocado por el desarrollo.
    1.3 Sistemas Involucrados
        Lista de sistemas operativos involucrados (servidor MOM, Oracle BD, capa de servicios, ControlM, integraciones externas si aplica).
    1.4 Calendarizacion
        Tabla pipe con la operacion en runtime del cambio (frecuencia de invocacion, ventanas batch, jobs ControlM nuevos del desarrollo). El despliegue del paquete NO se calendariza aqui.
    1.5 Definiciones, Acronimos y Abreviaciones
        Glosario corto: terminos tecnicos, nombres de packages/tablas, abreviaturas del dominio.

2. Modelo Arquitectura
    2.1 Vistas 4 + 1
        2.1.1 Vista Logica
        2.2.2 Vista de Desarrollo
        2.3.3 Vista de Proceso
        2.3.4 Vista Fisica
        2.3.5 Vista de Escenarios
    2.2 Secciones aplicadas para la Plataforma de BI
        Solo aplica cuando los cambios sean de la plataforma de BI
    2.3 Referencias de Estandares Usados
    2.4 Especificaciones de Hardware y Software
    2.5 Seguridad
    2.6 Componentes para reutilizar
    2.7 Desglose de procesos (aplica para POS)
        Solo aplica cuando los cambios sean de la plataforma del Punto de Venta

4 Requerimientos
    4.1 Modelado de la aplicacion
        Describir la manera en la que se pretende dar solucion a los requerimientos del desarrollo, anexar diagramas segun el cambio: Diagrama de clases, diagramas de flujo, diagramas de secuencia, diagramas de entidad relacion
    4.2 Requerimientos de Sistemas
        Detalle de como cubrir los requerimientos funcionales y no funcionales del aplicativo: ID, Nombre del requerimiento, Dependencias, Descripcion, Tipo de Componentes, Componente a utilizar
    4.3 FrontEnd (Aplica para BI)
        Solo aplica cuando los cambios sean de la plataforma de BI
    4.4 Restricciones de diseno
    4.5 Requerimientos de Licencia
        Especificar necesidad de licencias sobre productos asociados a la implementacion.
    4.6 Componentes Comprados
        Describe todos los componentes comprados a ser usados por el sistema, licencias aplicables, restricciones de uso, compatibilidad/interoperabilidad o estandares de interfaz.
    4.7 Interaccion con otros sistemas
        Describir la interaccion con otros sistemas y mapeo de informacion cuando la interaccion sea para intercambio de informacion.
    4.8 Requerimientos de base de datos, fuentes de informacion, destinos y procesamiento
        Detalle de como cubrir los requerimientos base de datos, fuentes de informacion, destinos y procesamiento: ID, Nombre del requerimiento, Dependencias, Descripcion, Tipo de Componentes, Componente a utilizar

5 Matriz de Pruebas Unitarias (Aplica para BI)
    Solo aplica cuando los cambios sean de la plataforma de BI

6 Especificaciones del Software (aplica para POS)
    Solo aplica cuando los cambios sean de la plataforma del Punto de Venta

7 Firmas de aprobacion
    Tabla estatica (por definir)

8 Control de revisiones
    Tabla estatica (por definir)
```

## Approach

Implement a sibling manual utility to `sdd-operational-doc`: an English skill/agent instruction contract plus a Spanish template asset. The skill resolves an archived change, reads only archived evidence, excludes installer scripts from final references, validates object inventory enum values, and refuses to invent missing product, ownership, runtime, database, security, or integration details.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `skills/sdd-technical-doc/SKILL.md` | New | Archive-only technical documentation workflow. |
| `skills/sdd-technical-doc/assets/technical-document-template.md` | New | Mandatory Spanish document template. |
| `agents/sdd/sdd-technical-doc.md` | New | Manual executor prompt. |
| `AGENTS.md` | Modified | Manual utility boundary. |
| `README.md` | Modified | Manual utility listing. |
| `openspec/specs/sdd-technical-documentation-workflow/spec.md` | New | Capability requirements. |
| `openspec/specs/sdd-execution-persistence-contracts/spec.md` | Modified | Non-DAG manual boundary. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Generated doc invents missing archive data | Medium | Require explicit unavailable-information text. |
| Technical doc is mistaken for a required SDD phase | Medium | Document manual post-archive status in skill, agent, README, and AGENTS. |
| Installer scripts leak into references | Medium | Explicit final-reference exclusion rule. |

## Rollback Plan

Delete the new `sdd-technical-doc` skill, template, agent, and spec; revert README/AGENTS and the persistence-contract spec note. No DAG migration or state recovery is needed because active SDD flows remain unchanged.

## Dependencies

- Existing `sdd-operational-doc` pattern for manual archive-consuming utilities.
- Existing skill style guide and opencode agent conventions.
- Archived SDD evidence availability for generated documents.

## Success Criteria

- [ ] New utility is documented as manual/post-archive and not part of required SDD DAG.
- [ ] Template preserves the required Spanish structure, conditional BI/POS sections, static approval/revision placeholders, and inventory table enums.
- [ ] Skill requires `No aplica.` and unavailable-information markers instead of inference.
- [ ] Final references include only functional/development sources and exclude installer scripts.
