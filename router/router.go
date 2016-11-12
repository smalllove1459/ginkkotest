package router

import (
	c "component_project/controller"
	"github.com/Unknwon/macaron"
)

func SetRouter(m *macaron.Macaron) {
	m.Get("/getmethod", c.GetMethod)
	m.Post("/postmethod", c.PostMethod)
}
