# evaluacion_crud

API CRUD para la gestion de formatos de evaluación del sistema Agora.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
EVALUACION_CRUD_DB_HOST = [descripción]
EVALUACION_CRUD_DB_NAME = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con EVALUACION_CRUD...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/evaluacion_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/evaluacion_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
EVALUACION_CRUD_PORT=8080 EVALUACION_CRUD_DB_HOST=127.0.0.1:27017 EVALUACION_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=evaluacion_crud . --no-cache
# docker run -p 80:80 evaluacion_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/evaluacion_crud

#2. Moverse a la carpeta del repositorio
cd evaluacion_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_crud)  | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_crud) |


## Modelo de Datos
[Modelo de Datos Relacional](https://drive.google.com/open?id=1Td88yP3jA7Y_kbAzPaOxE2nn41Teu7tl)  
[SQL evaluacion_proveedores](https://drive.google.com/open?id=1mZLmSuDIbQzwIidVCA29c9LCFWljMZoG)  
[Registros Paramétricos](https://drive.google.com/open?id=1gUK_4g_-vU1LwKsgO_xsMG0yUWGADm9OZRF3fiDRBXg)


## Licencia

This file is part of evaluacion_crud.

evaluacion_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

evaluacion_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with evaluacion_crud. If not, see https://www.gnu.org/licenses/.
