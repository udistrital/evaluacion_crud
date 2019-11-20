-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2-beta
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: evaluacion_proveedores | type: DATABASE --
-- -- DROP DATABASE IF EXISTS evaluacion_proveedores;
-- CREATE DATABASE evaluacion_proveedores;
-- -- ddl-end --
-- 

-- object: evaluacion | type: SCHEMA --
-- DROP SCHEMA IF EXISTS evaluacion CASCADE;
CREATE SCHEMA evaluacion;
-- ddl-end --
ALTER SCHEMA evaluacion OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,evaluacion;
-- ddl-end --

-- object: evaluacion.evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.evaluacion CASCADE;
CREATE TABLE evaluacion.evaluacion (
	id serial NOT NULL,
	proveedor_id integer NOT NULL,
	contrato_suscrito integer NOT NULL,
	vigencia integer,
	cotizacion_id integer,
	plantilla_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	aprobado boolean,
	activo boolean NOT NULL,
	CONSTRAINT evaluacion_pk PRIMARY KEY (id),
	CONSTRAINT uq_proveedor_contrato_vigencia UNIQUE (proveedor_id,contrato_suscrito,vigencia)

);
-- ddl-end --
COMMENT ON COLUMN evaluacion.evaluacion.id IS 'pk de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion.evaluacion.proveedor_id IS 'identificador, ya sea nit o cedula del proveedor';
-- ddl-end --
COMMENT ON COLUMN evaluacion.evaluacion.contrato_suscrito IS 'guarda el numero del contrato suscrito';
-- ddl-end --
COMMENT ON COLUMN evaluacion.evaluacion.plantilla_id IS 'hace referencia a la plantilla que se uso para realizar la evaluacion guardada';
-- ddl-end --
COMMENT ON CONSTRAINT uq_proveedor_contrato_vigencia ON evaluacion.evaluacion  IS 'UQ con el fin de que no se diplique una evaluacion para un proveedor y mismo contrato';
-- ddl-end --
ALTER TABLE evaluacion.evaluacion OWNER TO postgres;
-- ddl-end --

-- object: evaluacion.plantilla | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.plantilla CASCADE;
CREATE TABLE evaluacion.plantilla (
	id serial NOT NULL,
	activo boolean NOT NULL,
	descripcion varchar(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	usuario varchar(50) NOT NULL,
	CONSTRAINT plantilla_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.plantilla IS 'guarda las plantillas';
-- ddl-end --
COMMENT ON COLUMN evaluacion.plantilla.activo IS 'plantilla activa';
-- ddl-end --
ALTER TABLE evaluacion.plantilla OWNER TO postgres;
-- ddl-end --

-- object: evaluacion.seccion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.seccion CASCADE;
CREATE TABLE evaluacion.seccion (
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	id_plantilla integer NOT NULL,
	seccion_hija_id integer,
	activo boolean NOT NULL,
	CONSTRAINT seccion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.seccion IS 'representara com plantilla dinamica a los card o filas dento de ese card';
-- ddl-end --
COMMENT ON COLUMN evaluacion.seccion.nombre IS 'nombre en el cual se visualizara en el label de la aplicacion';
-- ddl-end --
ALTER TABLE evaluacion.seccion OWNER TO postgres;
-- ddl-end --

-- object: plantilla_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.seccion DROP CONSTRAINT IF EXISTS plantilla_fk CASCADE;
ALTER TABLE evaluacion.seccion ADD CONSTRAINT plantilla_fk FOREIGN KEY (id_plantilla)
REFERENCES evaluacion.plantilla (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: rel_seccion_seccion_hija | type: Generic SQL Object --
ALTER TABLE evaluacion.seccion ADD CONSTRAINT rel_seccion_seccion_hija
   FOREIGN KEY (seccion_hija_id)
   REFERENCES evaluacion.seccion (id)  
   ON DELETE  RESTRICT 
   ON UPDATE  CASCADE 
   NOT DEFERRABLE 
   INITIALLY IMMEDIATE
;
-- ddl-end --

-- object: evaluacion.item | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.item CASCADE;
CREATE TABLE evaluacion.item (
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	valor varchar(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	id_seccion integer NOT NULL,
	id_tipo_item integer NOT NULL,
	activo boolean NOT NULL,
	id_estilo_pipe integer,
	CONSTRAINT item_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN evaluacion.item.nombre IS 'nombre del label
';
-- ddl-end --
COMMENT ON COLUMN evaluacion.item.valor IS 'valor usado en caso de item de tipo input o similar';
-- ddl-end --
ALTER TABLE evaluacion.item OWNER TO postgres;
-- ddl-end --

-- object: evaluacion.tipo_item | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.tipo_item CASCADE;
CREATE TABLE evaluacion.tipo_item (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	codigo_abreviacion varchar(20),
	descripcion varchar(250),
	activo boolean NOT NULL,
	numero_orden decimal(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT tipo_item_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.tipo_item IS 'se almacenaran tipos de item , como input, label, select, entre otros';
-- ddl-end --
ALTER TABLE evaluacion.tipo_item OWNER TO postgres;
-- ddl-end --

-- object: evaluacion.opciones | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.opciones CASCADE;
CREATE TABLE evaluacion.opciones (
	id serial NOT NULL,
	nombre varchar(20) NOT NULL,
	descripcion varchar(250),
	valor integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT opciones_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.opciones IS 'opciones para los item tipo select o similar';
-- ddl-end --
COMMENT ON COLUMN evaluacion.opciones.nombre IS 'nombre que se mostrara en el select';
-- ddl-end --
COMMENT ON COLUMN evaluacion.opciones.descripcion IS 'descripcion de la opcion';
-- ddl-end --
COMMENT ON COLUMN evaluacion.opciones.valor IS 'valor numerico para la evaluacion';
-- ddl-end --
ALTER TABLE evaluacion.opciones OWNER TO postgres;
-- ddl-end --

-- object: seccion_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.item DROP CONSTRAINT IF EXISTS seccion_fk CASCADE;
ALTER TABLE evaluacion.item ADD CONSTRAINT seccion_fk FOREIGN KEY (id_seccion)
REFERENCES evaluacion.seccion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: tipo_item_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.item DROP CONSTRAINT IF EXISTS tipo_item_fk CASCADE;
ALTER TABLE evaluacion.item ADD CONSTRAINT tipo_item_fk FOREIGN KEY (id_tipo_item)
REFERENCES evaluacion.tipo_item (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion.opcion_item | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.opcion_item CASCADE;
CREATE TABLE evaluacion.opcion_item (
	id serial NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	id_item integer NOT NULL,
	id_opciones integer NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT opcion_item_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.opcion_item IS 'tabla de rompimiento entre item y opcion';
-- ddl-end --
ALTER TABLE evaluacion.opcion_item OWNER TO postgres;
-- ddl-end --

-- object: item_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.opcion_item DROP CONSTRAINT IF EXISTS item_fk CASCADE;
ALTER TABLE evaluacion.opcion_item ADD CONSTRAINT item_fk FOREIGN KEY (id_item)
REFERENCES evaluacion.item (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: opciones_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.opcion_item DROP CONSTRAINT IF EXISTS opciones_fk CASCADE;
ALTER TABLE evaluacion.opcion_item ADD CONSTRAINT opciones_fk FOREIGN KEY (id_opciones)
REFERENCES evaluacion.opciones (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion.condicion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.condicion CASCADE;
CREATE TABLE evaluacion.condicion (
	id serial NOT NULL,
	seccion_dependencia_id integer NOT NULL,
	opcion_item_id integer NOT NULL,
	id_seccion integer,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT condicion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.condicion IS 'referencia de la opcion que hara qe aparezca o no la opcio, esto hacer referencia entre secciones hermanas';
-- ddl-end --
COMMENT ON COLUMN evaluacion.condicion.seccion_dependencia_id IS 'hace refenrencia al seccion de la cual segun su opcion seleccionada, la presente opcion aparecera o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion.condicion.opcion_item_id IS 'hace referencia a la opcion seleccionada en la seccion que causa la dependencia';
-- ddl-end --
ALTER TABLE evaluacion.condicion OWNER TO postgres;
-- ddl-end --

-- object: seccion_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.condicion DROP CONSTRAINT IF EXISTS seccion_fk CASCADE;
ALTER TABLE evaluacion.condicion ADD CONSTRAINT seccion_fk FOREIGN KEY (id_seccion)
REFERENCES evaluacion.seccion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion.clasificacion_plantilla | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.clasificacion_plantilla CASCADE;
CREATE TABLE evaluacion.clasificacion_plantilla (
	id serial NOT NULL,
	id_clasificacion integer NOT NULL,
	id_plantilla integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT clasificacion_plantilla_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.clasificacion_plantilla IS 'tabla de rompimiento entre tipo de clasificacion y la plantilla';
-- ddl-end --
ALTER TABLE evaluacion.clasificacion_plantilla OWNER TO postgres;
-- ddl-end --

-- object: evaluacion.clasificacion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.clasificacion CASCADE;
CREATE TABLE evaluacion.clasificacion (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	codigo_abreviacion varchar(20),
	descripcion varchar(250),
	activo boolean NOT NULL,
	numero_orden decimal(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	limite_inferior integer NOT NULL,
	limite_superior integer NOT NULL,
	CONSTRAINT clasificacion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.clasificacion IS 'guarda la los tipos de clasificacion del resultado final de la evaluacion';
-- ddl-end --
ALTER TABLE evaluacion.clasificacion OWNER TO postgres;
-- ddl-end --

-- object: clasificacion_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.clasificacion_plantilla DROP CONSTRAINT IF EXISTS clasificacion_fk CASCADE;
ALTER TABLE evaluacion.clasificacion_plantilla ADD CONSTRAINT clasificacion_fk FOREIGN KEY (id_clasificacion)
REFERENCES evaluacion.clasificacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: plantilla_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.clasificacion_plantilla DROP CONSTRAINT IF EXISTS plantilla_fk CASCADE;
ALTER TABLE evaluacion.clasificacion_plantilla ADD CONSTRAINT plantilla_fk FOREIGN KEY (id_plantilla)
REFERENCES evaluacion.plantilla (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion.estilo_pipe | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.estilo_pipe CASCADE;
CREATE TABLE evaluacion.estilo_pipe (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	codigo_abreviacion varchar(20),
	descripcion varchar(250),
	activo boolean NOT NULL,
	numero_orden decimal(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT estilo_pipe_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.estilo_pipe IS 'almacenara los tipos de pipes aplicables al item, esto dependera si el pipe existe en el cliente';
-- ddl-end --
ALTER TABLE evaluacion.estilo_pipe OWNER TO postgres;
-- ddl-end --

-- object: estilo_pipe_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.item DROP CONSTRAINT IF EXISTS estilo_pipe_fk CASCADE;
ALTER TABLE evaluacion.item ADD CONSTRAINT estilo_pipe_fk FOREIGN KEY (id_estilo_pipe)
REFERENCES evaluacion.estilo_pipe (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion.resultado_evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion.resultado_evaluacion CASCADE;
CREATE TABLE evaluacion.resultado_evaluacion (
	id serial NOT NULL,
	resultado_evaluacion json NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo timestamp NOT NULL,
	id_evaluacion integer NOT NULL,
	CONSTRAINT resultado_evaluacion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion.resultado_evaluacion IS 'almacena los distintos resultados de la evaluacion, tanto la primera evaliuacion realizada como , si llegan a haber correcciones sobre la evaluacion';
-- ddl-end --
ALTER TABLE evaluacion.resultado_evaluacion OWNER TO postgres;
-- ddl-end --

-- object: evaluacion_fk | type: CONSTRAINT --
-- ALTER TABLE evaluacion.resultado_evaluacion DROP CONSTRAINT IF EXISTS evaluacion_fk CASCADE;
ALTER TABLE evaluacion.resultado_evaluacion ADD CONSTRAINT evaluacion_fk FOREIGN KEY (id_evaluacion)
REFERENCES evaluacion.evaluacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- -- Permisos de usuario
-- GRANT USAGE ON SCHEMA evaluacion TO test;
-- GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA evaluacion TO test;
-- GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA evaluacion TO test;

-- -- Permisos de usuario
-- GRANT USAGE ON SCHEMA evaluacion TO desarrollooas;
-- GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA evaluacion TO desarrollooas;
-- GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA evaluacion TO desarrollooas;

-- Permisos de usuario
GRANT USAGE ON SCHEMA evaluacion TO postgres;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA evaluacion TO postgres;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA evaluacion TO postgres;
