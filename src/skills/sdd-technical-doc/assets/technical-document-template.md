# Plantilla de Documentación Técnica

## Reglas obligatorias

- El documento final debe estar en español neutral y profesional.
- Si una sección no aplica según la evidencia archivada, escribir exactamente: `No aplica.`.
- Si una sección aplica pero falta información en la evidencia archivada, escribir exactamente: `Información no disponible en la evidencia archivada.`.
- No inferir ni inventar producto, responsables, objetos, operaciones, seguridad, runtime, integraciones, licencias, hardware, software, pruebas o referencias.
- Las referencias finales solo pueden incluir FCTI, historia, `.pks`, `.pkb` y `.sql` de packages/tablas creadas o modificadas por el desarrollo. No incluir scripts instaladores.
- Las secciones condicionales BI/POS solo aplican si la evidencia archivada lo demuestra; si no, usar `No aplica.`.
- Las vistas 4 + 1 deben incluir diagramas Mermaid cuando la evidencia archivada o el contexto explícito del documento final lo soporten. No inventar actores, clases, componentes, infraestructura ni relaciones; si no hay evidencia suficiente para un diagrama, escribir exactamente `Información no disponible en la evidencia archivada.` debajo del título del diagrama correspondiente.

## Identificación del Producto

| Producto | Dueño de aplicación | Plataforma | Líder de Proyecto TI | Desarrollador |
| --- | --- | --- | --- | --- |
| Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. |

## Referencias

- `archivo` — descripción.

> Incluir únicamente FCTI, historia, `.pks`, `.pkb` y `.sql` de packages/tablas creadas o modificadas por el desarrollo. Excluir scripts instaladores.

# 1. Presentación del Producto

Información no disponible en la evidencia archivada.

## 1.1 Objetivo

Información no disponible en la evidencia archivada.

## 1.2 Alcance

Información no disponible en la evidencia archivada.

| Tipo | Esquema.Objeto | Operacion | Descripcion funcional |
| --- | --- | --- | --- |
| Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. |

Valores permitidos para `Tipo`: TABLE, INDEX, PACKAGE SPEC, PACKAGE BODY, PROCEDURE, FUNCTION, TRIGGER, JOB ControlM, VIEW, SEQUENCE, GRANT.

Valores permitidos para `Operacion`: CREATE, ALTER, REPLACE, DROP.

Debe existir una fila por cada objeto tocado por el desarrollo cuando la evidencia archivada lo identifique. Los valores fuera de los enums permitidos deben corregirse solo con evidencia archivada o marcarse como `Información no disponible en la evidencia archivada.`.

## 1.3 Sistemas Involucrados

Información no disponible en la evidencia archivada.

## 1.4 Calendarización

| Operación runtime | Frecuencia | Ventana batch | Job ControlM | Observaciones |
| --- | --- | --- | --- | --- |
| Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. |

> Calendarizar la operación runtime del cambio: frecuencia de invocación, ventanas batch y jobs ControlM nuevos del desarrollo. No calendarizar el despliegue del paquete.

## 1.5 Definiciones, Acrónimos y Abreviaciones

Información no disponible en la evidencia archivada.

# 2. Modelo Arquitectura

## 2.1 Vistas 4 + 1

### 2.1.1 Vista Lógica

Información no disponible en la evidencia archivada.

#### Diagrama de clases

Información no disponible en la evidencia archivada.

### 2.2.2 Vista de Desarrollo

Información no disponible en la evidencia archivada.

#### Diagrama de componentes

Información no disponible en la evidencia archivada.

### 2.3.3 Vista de Proceso

Información no disponible en la evidencia archivada.

#### Diagrama de secuencia

Información no disponible en la evidencia archivada.

### 2.3.4 Vista Física

Información no disponible en la evidencia archivada.

#### Diagrama de despliegue

Información no disponible en la evidencia archivada.

### 2.3.5 Vista de Escenarios

Información no disponible en la evidencia archivada.

#### Diagrama de casos de uso

Información no disponible en la evidencia archivada.

## 2.2 Secciones aplicadas para la Plataforma de BI

No aplica.

## 2.3 Referencias de Estándares Usados

Información no disponible en la evidencia archivada.

## 2.4 Especificaciones de Hardware y Software

Información no disponible en la evidencia archivada.

## 2.5 Seguridad

Información no disponible en la evidencia archivada.

## 2.6 Componentes para reutilizar

Información no disponible en la evidencia archivada.

## 2.7 Desglose de procesos (aplica para POS)

No aplica.

# 4. Requerimientos

## 4.1 Modelado de la aplicación

Información no disponible en la evidencia archivada.

> Describir la solución del desarrollo y anexar diagramas según aplique: clases, flujo, secuencia o entidad-relación.

## 4.2 Requerimientos de Sistemas

| ID | Nombre del requerimiento | Dependencias | Descripción | Tipo de Componentes | Componente a utilizar |
| --- | --- | --- | --- | --- | --- |
| Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. |

## 4.3 FrontEnd (Aplica para BI)

No aplica.

## 4.4 Restricciones de diseño

Información no disponible en la evidencia archivada.

## 4.5 Requerimientos de Licencia

Información no disponible en la evidencia archivada.

## 4.6 Componentes Comprados

Información no disponible en la evidencia archivada.

## 4.7 Interacción con otros sistemas

Información no disponible en la evidencia archivada.

## 4.8 Requerimientos de base de datos, fuentes de información, destinos y procesamiento

| ID | Nombre del requerimiento | Dependencias | Descripción | Tipo de Componentes | Componente a utilizar |
| --- | --- | --- | --- | --- | --- |
| Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. | Información no disponible en la evidencia archivada. |

# 5. Matriz de Pruebas Unitarias (Aplica para BI)

No aplica.

# 6. Especificaciones del Software (aplica para POS)

No aplica.

# 7. Firmas de aprobación

| Rol | Nombre | Firma | Fecha |
| --- | --- | --- | --- |
| Dueño de aplicación | Por definir | Por definir | Por definir |
| Líder de Proyecto TI | Por definir | Por definir | Por definir |
| Desarrollador | Por definir | Por definir | Por definir |

# 8. Control de revisiones

| Versión | Fecha | Autor | Descripción |
| --- | --- | --- | --- |
| 1.0 | Por definir | Por definir | Versión inicial. |
