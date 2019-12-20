-- Datos para tabla parametrica tipo_item 
insert into evaluacion.tipo_item (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('label','lbl','item usado como label para la plantilla',TRUE,null,'2019-11-12','2019-11-12');
insert into evaluacion.tipo_item (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('input','inp','item usado como input para la plantilla',TRUE,null,'2019-11-12','2019-11-12');
insert into evaluacion.tipo_item (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('select','slt','item usado como slect para la plantilla',TRUE,null,'2019-11-12','2019-11-12');

-- Datos para tabla estilo_pipe
insert into evaluacion.estilo_pipe (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('ngxCapitalize','ngxcap','pipe que coloca la primeta letra en mayuscula y las demas en minuscula',TRUE,null,'2019-11-12','2019-11-12');
insert into evaluacion.estilo_pipe (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('uppercase','upper','pupper case a los string',TRUE,null,'2019-11-12','2019-11-12');
insert into evaluacion.estilo_pipe (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('lowercase','lower','lowercase a los string',TRUE,null,'2019-11-12','2019-11-12');
insert into evaluacion.estilo_pipe (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion) 
    values ('date','date','pipe para formato de fechas',TRUE,null,'2019-11-12','2019-11-12');

-- Datos para tabla clasificacion
insert into evaluacion.clasificacion (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion,limite_inferior,limite_superior) 
    values ('EXCELENTE','EX','calificacion para el proveedor que cumplio con todos los termnios',TRUE,null,'2019-11-12','2019-11-12',80,100);
insert into evaluacion.clasificacion (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion,limite_inferior,limite_superior) 
    values ('BUENO','BN','calificacion para el proveedor que cumplio parcialmentente con los terminos',TRUE,null,'2019-11-12','2019-11-12',46,79);
insert into evaluacion.clasificacion (nombre,codigo_abreviacion,descripcion,activo,numero_orden,fecha_creacion,fecha_modificacion,limite_inferior,limite_superior) 
    values ('MALO','ML','caloficacion para el proveedor que cumplio con pocos terminos',TRUE,null,'2019-11-12','2019-11-12',0,45);