# sdd-technical-documentation-workflow Specification

## Purpose

Define the manual post-archive `sdd-technical-doc` utility that generates one Spanish technical documentation Markdown file from an already archived SDD change.

## Requirements

### Requirement: Manual Post-Archive Utility

`sdd-technical-doc` MUST remain a manual, user-invocable, archive-consuming utility. It MUST NOT be a required SDD DAG, status, dependency, verify, or archive phase. It MUST generate exactly one Markdown technical documentation file in neutral professional Spanish from an already archived SDD change.

#### Scenario: Archive flow is unchanged

- GIVEN an SDD change satisfies verify and archive criteria
- WHEN archive runs
- THEN completion MUST NOT require `sdd-technical-doc`
- AND no DAG or status successor MUST be added for it.

#### Scenario: Manual generation succeeds

- GIVEN the user invokes `sdd-technical-doc` for an archived change
- WHEN archived evidence is readable
- THEN exactly one Spanish `.md` technical documentation file MUST be produced.

### Requirement: Archived Evidence Only

The utility MUST use archived SDD evidence only and MUST NOT invent product, owner, functional, runtime, database, security, integration, object, or reference details. Inapplicable sections MUST contain `No aplica.`. Applicable but unavailable information MUST use a clear marker such as `Información no disponible en la evidencia archivada.`

#### Scenario: Section is inapplicable

- GIVEN archived evidence shows a section does not apply
- WHEN the document is generated
- THEN that section MUST say `No aplica.`

#### Scenario: Evidence is missing

- GIVEN a required detail is absent from archived evidence
- WHEN the document is generated
- THEN the document MUST mark the detail as unavailable
- AND it MUST NOT infer or fabricate the value.

### Requirement: Mandatory Spanish Document Structure

The generated document MUST preserve the requested structure: Identificación del Producto with references; 1 Presentación del Producto with 1.1-1.5; 2 Modelo Arquitectura with 2.1 4+1 views, BI-only 2.2, standards, hardware/software, security, reusable components, POS-only process breakdown; 4 Requerimientos with 4.1-4.8; BI-only unit-test matrix section 5; POS-only software specification section 6; static approval-signature section 7; and static revision-control section 8.

#### Scenario: Conditional platform section

- GIVEN archived evidence does not identify a BI or POS change
- WHEN BI-only or POS-only sections are rendered
- THEN each non-applicable conditional section MUST contain `No aplica.`

#### Scenario: Static placeholders are preserved

- GIVEN the document is generated
- WHEN sections 7 and 8 are rendered
- THEN they MUST retain static approval and revision placeholders.

### Requirement: References and Object Inventory

Final references MUST be restricted to functional/development sources: FCTI, story, and `.pks`, `.pkb`, or `.sql` package/table sources created or modified. Final references MUST NOT include installer scripts. Section 1.2 Alcance MUST include an object inventory table with columns `Tipo`, `Esquema.Objeto`, `Operacion`, and `Descripcion funcional`; `Tipo` MUST be one of `TABLE`, `INDEX`, `PACKAGE SPEC`, `PACKAGE BODY`, `PROCEDURE`, `FUNCTION`, `TRIGGER`, `JOB ControlM`, `VIEW`, `SEQUENCE`, `GRANT`; `Operacion` MUST be one of `CREATE`, `ALTER`, `REPLACE`, `DROP`.

#### Scenario: Installer script appears in archive

- GIVEN archived evidence contains installer scripts
- WHEN final references are written
- THEN installer scripts MUST be excluded from references.

#### Scenario: Inventory value is invalid

- GIVEN an object inventory row would require a `Tipo` or `Operacion` outside the allowed enums
- WHEN the document is generated
- THEN the row MUST use the unavailable-information marker or be corrected from archived evidence only.
