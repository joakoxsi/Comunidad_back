package main

import (
	"context"
	"os"
	"time"

	"github.com/FelipeMarchantVargas/Prueba/controllers"
	"github.com/FelipeMarchantVargas/Prueba/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://felipemarchantv:212217204@cluster0.phuhme5.mongodb.net/?retryWrites=true&w=majority"))
    if err != nil {
        panic(err)
    }
    defer client.Disconnect(ctx)

	uc := controllers.NewUserController(client)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	routes.Setup(app, uc)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":"+port)

}