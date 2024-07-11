package main

import (
	"log"

	"github.com/barganakukuhraditya/BOILERPLATE_TUBES/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/barganakukuhraditya/BOILERPLATE_TUBES/url"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/swagger" // swagger handler
	// _ "github.com/barganakukuhaditya/BOILERPLATE_TUBES/docs"
)

// @title SWAGGER TUGAS BESAR
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/barganakukuhraditya
// @contact.email 714220013@std.ulbi.ac.id

// @host tb-parfume2024.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
