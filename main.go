package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/students", getStudents)
	app.Get("/student/:id", getStudentById)
	app.Post("/student", createNewStudent)
	app.Put("/student/:id", updateStudent)
	app.Delete("/student/:id", deletedStudent)
	app.Post("/upload", uploadImage)
	app.Get("/test-html", testHTML)
	app.Get("/config", getEnv)
	app.Listen(":3000")
}

