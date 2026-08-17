=================================================================================
                       Información General 
=================================================================================
{{APP_BLOCKS}}

[For each app, the APP_BLOCK looks like:]
{N}. 	Aplicación: 	{{APP_NAME}}
	Versión:		{{APP_VERSION}}
	Repositorio:    {{REPO_URL}}     {{BRANCH}}
  
		| Ambiente          | Host                          | Contenedor                   |
        --------------------|-------------------------------|------------------------------|
{{ENVIRONMENT_TABLE}}

=================================================================================
                       Pasos de Instalación
=================================================================================
I.   Ingresar a la consola de administración WebLogic

II.  Apagar los servidores administrados:
{{MANAGED_SERVERS}}

III. Ingresar en modo edición de la consola a Deployments/App Deployments

IV.  Eliminar las aplicaciones:
{{APP_NAMES}}

V.   Darle commit a los cambios

VI.  Crear un nuevo despliegue por cada aplicación:
     Nombre:     {{APP_NAME}}
     Destinos:   {{CONTAINERS}}
     Despliegue: Iniciar aplicación
     Aplicativo:
         Carpeta:    /Install/AppDeploy/{{ARTIFACT_NAME}}

VII. Darle commit a los cambios

VIII. Iniciar los servidores administrados:
{{MANAGED_SERVERS}}

IX.  Validar que las aplicaciones se encuentren activas:
{{APP_NAMES}}

{{INSTALL_STEPS}}

=================================================================================
                       Control de Versiones del documento
=================================================================================
---Versión {{VERSION}}
---Fecha del cambio: {{DATE}}
---Autor del cambio: {{AUTHOR}}
---CHG: {{CHG_NUMBER}}
 	
 	/*HISTORIAS DE USUARIO*/	
{{USER_STORIES}}
 	
 	/*CORRECCIONES*/
{{CORRECTIONS}}

=================================================================================
                       Control de Revisiones de la Plantilla STTI
=================================================================================

---Revisión 0
---Fecha del cambio: 
---Autor del cambio: 
  * Emisión del documento
