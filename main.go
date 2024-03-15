package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/percoguru/notes-api-fiber/database"
	"github.com/percoguru/notes-api-fiber/router"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

    http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Application is alive and healthy\n"))
    })

    // Start the HTTP server on port 8080 in a separate goroutine
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            panic(err)
        }
    }()
	// Run Prometheus metrics endpoint on port 2112
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	// Listen on PORT 3000
	app.Listen(":3000")
}
