package main

import (
	//"shangqu-finance/router"
	"component_project/router"
	"github.com/Unknwon/macaron"
)

func main() {
	m := macaron.New()
	router.SetRouter(m)
	m.Run(8000)
}
