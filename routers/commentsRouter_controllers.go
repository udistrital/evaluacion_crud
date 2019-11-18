package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ClasificacionPlantillaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:CondicionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EstiloPipeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:EvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:OpcionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:ResultadoEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_crud/controllers:TipoItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
