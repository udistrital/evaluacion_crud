# evaluacion_crud

evaluacion_crud, CRUD para la gestion de formatos de evaluación del sistema Agora. El proyecto está escrito en el lenguaje Go. generado mediante el **[framework beego](https://beego.me/)**.

## Especificación Técnica

### Requerimientos

- [Postgres](https://www.postgresql.org/)
- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [Beego y Bee](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)

### Instalación
Para instalar el proyecto de debe relizar lo siguientes pasos:

#### Opción 1
Ejecutar desde la terminal **go get repositorio**:
```bash
go get github.com/udistrital/evaluacion_crud
```
#### Opción 2
Clonar en el proyecto en la carpeta local **go/src/github.com/udistrital**:
```bash
cd ~go/src/github.com/udistrital
git clone https://github.com/udistrital/evaluacion_crud
# Ir a la carpeta del proyecto
cd evaluacion_crud
# Instalar dependencias del proyecto:
go get
```

### Variables de Entorno

- EVALUACIONES_CRUD__PGDB=[nombre de la base de datos]  
- EVALUACIONES_CRUD__PGPASS=[password del usuario]
- EVALUACIONES_CRUD__PGURLS=[direccion de la base de datos]
- EVALUACIONES_CRUD__PGUSER=[usuario con acceso a la base de datos]
- EVALUACIONES_CRUD__PGSCHEMA=[esquema donde se ubican las tablas]
- EVALUACIONES_HTTP_PORT=[puerto de ejecucion] bee run

### Ejecución del proyecto

```shell
EVALUACIONES_CRUD__PGDB=XXX EVALUACIONES_CRUD__PGPASS=XXX EVALUACIONES_CRUD__PGURLS=XXX EVALUACIONES_CRUD__PGUSER=XXX EVALUACIONES_CRUD__PGSCHEMA=XXX EVALUACIONES_HTTP_PORT=XXX bee run
```
- Com documentación swagger:

```shell
EVALUACIONES_CRUD__PGDB=XXX EVALUACIONES_CRUD__PGPASS=XXX EVALUACIONES_CRUD__PGURLS=XXX EVALUACIONES_CRUD__PGUSER=XXX EVALUACIONES_CRUD__PGSCHEMA=XXX EVALUACIONES_HTTP_PORT=XXX bee run -downdoc=true -gendoc=true
```

### EndPoints

Al ejecutar el swagger se puede tener mayor apreciacion de los diferentes metodos de peticion por cada endpoint cuales son los distinpos endpoint disponibles y como usarlos.

### Modelo de Datos
El modelo de datos se comparte en imagen y en **[.dbm](https://drive.google.com/open?id=1Td88yP3jA7Y_kbAzPaOxE2nn41Teu7tl)** para que pueda ser editado en pgModeler.
![](evaluacion_proveedores.png)

### SQL del Modelo de Datos
**[SQL evaluacion_proveedores](https://drive.google.com/open?id=1mZLmSuDIbQzwIidVCA29c9LCFWljMZoG)**.

Registros paramétricos a tener encuenta: **[Parametricas](https://drive.google.com/open?id=1gUK_4g_-vU1LwKsgO_xsMG0yUWGADm9OZRF3fiDRBXg)**.

## Licencia

This file is part of evaluacion_crud

evaluacion_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
