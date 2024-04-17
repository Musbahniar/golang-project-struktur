package main

import (
	"time"

	_ "github.com/Musbahniar/golang-project-struktur/docs"
	"github.com/Musbahniar/golang-project-struktur/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title			Belajar Resful Api
// @version		1.0
// @description	Documentation for Belajar Resful Api
// @termsOfService	http://swagger.io/terms/
func main() {
	// Start a new fiber app
	app := fiber.New()

	// CORS config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
	}))

	// Helmet config
	app.Use(helmet.New())

	// Limiter config
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	}))

	// Swagger docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/healthcheck",
	}))

	router.SetupRoutes(app)

	// Logger Console config
	app.Use(logger.New())

	// Logger File config
	// file, err := os.OpenFile("./my_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer file.Close()
	// loggerConfig := logger.Config{Output: file}
	// app.Use(logger.New(loggerConfig))

	// Listen on PORT 3000
	errRun := app.Listen(":3000")
	if errRun != nil {
		panic(errRun)
	}
}
