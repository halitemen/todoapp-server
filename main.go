package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"server/handler"
	"server/repository"
	"server/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//os.Setenv("MONGO_URI", "")
	e := echo.New()
	e.Use(middleware.CORS())

	todoHandler := getService()

	e.GET("/", todoHandler.SayHello)
	e.GET("/todo/alltodo", todoHandler.GetTodos)
	e.POST("/todo/addtodo", todoHandler.CreateTodo)

	e.Logger.Fatal(e.Start(":16500"))

	fmt.Println("Server Started.")
}

func getService() handler.Handler {

	collection := getCollection("mydb", "todos")
	if collection == nil {
		fmt.Println("collection cannot be nil")
		os.Exit(-1)
	}

	repository := repository.NewTodoRepository(getCollection("mydb", "todos"))
	todoService := service.NewTodoService(repository)
	todoHandler := handler.NewTodoHandler(todoService)
	return todoHandler
}

func getCollection(dbname, collname string) *mongo.Collection {
	client := getClient()
	if client != nil {
		return client.Database(dbname).Collection(collname)
	}
	return nil
}

func getClient() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil
	}

	fmt.Println(uri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
