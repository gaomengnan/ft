package routes

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gaomengnan/ft/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func RegisterRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return Render(c, views.Blog("blog"))
	})

	app.Get("/home/blog", func(c *fiber.Ctx) error {
		fmt.Println("blog")
		return Render(c, views.Blog("blog"))
	})

	app.Get("/home/:tab", func(c *fiber.Ctx) error {
		tab := c.Params("tab")
		return Render(c, views.Blog(tab))
	})
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
