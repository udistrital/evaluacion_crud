package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterResultadoEvaluacion_20191231_181012 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterResultadoEvaluacion_20191231_181012{}
	m.Created = "20191231_181012"

	migration.Register("AlterResultadoEvaluacion_20191231_181012", m)
}

// Run the migrations
func (m *AlterResultadoEvaluacion_20191231_181012) Up() {
	// use m.SQL("ALTER TABLE ...") to make table update
	file, err := ioutil.ReadFile("../scripts/20191231_181012_alter_resultado_evaluacion.sql")

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
