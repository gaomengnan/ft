package routes

import (
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
		return Render(c, views.Home(""))
	})

	app.Get("/home/:tab", func(c *fiber.Ctx) error {
		tab := c.Params("tab")
		return Render(c, views.Home(tab))
	})

	app.Post("/home/blog", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})

	app.Post("/home/tool", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})
	app.Post("/home/road", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})
	app.Post("/home/work", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})
	app.Post("/home/learn", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})
	app.Post("/home/travel", func(c *fiber.Ctx) error {
		return Render(c, views.Blog())
		// templ.Handler(views.Blog(), )
	})
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
