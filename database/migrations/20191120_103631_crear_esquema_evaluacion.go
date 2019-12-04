package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaEvaluacion_20191120_103631 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaEvaluacion_20191120_103631{}
	m.Created = "20191120_103631"

	migration.Register("CrearEsquemaEvaluacion_20191120_103631", m)
}

// Run the migrations
func (m *CrearEsquemaEvaluacion_20191120_103631) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191120_103631_crear_esquema_evaluacion_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearEsquemaEvaluacion_20191120_103631) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191120_103631_crear_esquema_evaluacion_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
