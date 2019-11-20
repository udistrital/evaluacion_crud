package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertarParametricasEvaluacion_20191120_143348 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarParametricasEvaluacion_20191120_143348{}
	m.Created = "20191120_143348"

	migration.Register("InsertarParametricasEvaluacion_20191120_143348", m)
}

// Run the migrations
func (m *InsertarParametricasEvaluacion_20191120_143348) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191120_143348_insertar_parametricas_evaluacion_up.sql")

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
func (m *InsertarParametricasEvaluacion_20191120_143348) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191120_143348_insertar_parametricas_evaluacion_down.sql")

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
