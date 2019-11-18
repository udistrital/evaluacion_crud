// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/evaluacion_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_item",
			beego.NSInclude(
				&controllers.TipoItemController{},
			),
		),

		beego.NSNamespace("/opciones",
			beego.NSInclude(
				&controllers.OpcionesController{},
			),
		),

		beego.NSNamespace("/opcion_item",
			beego.NSInclude(
				&controllers.OpcionItemController{},
			),
		),

		beego.NSNamespace("/seccion",
			beego.NSInclude(
				&controllers.SeccionController{},
			),
		),

		beego.NSNamespace("/condicion",
			beego.NSInclude(
				&controllers.CondicionController{},
			),
		),

		beego.NSNamespace("/clasificacion",
			beego.NSInclude(
				&controllers.ClasificacionController{},
			),
		),

		beego.NSNamespace("/plantilla",
			beego.NSInclude(
				&controllers.PlantillaController{},
			),
		),

		beego.NSNamespace("/clasificacion_plantilla",
			beego.NSInclude(
				&controllers.ClasificacionPlantillaController{},
			),
		),

		beego.NSNamespace("/estilo_pipe",
			beego.NSInclude(
				&controllers.EstiloPipeController{},
			),
		),

		beego.NSNamespace("/item",
			beego.NSInclude(
				&controllers.ItemController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/resultado_evaluacion",
			beego.NSInclude(
				&controllers.ResultadoEvaluacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
