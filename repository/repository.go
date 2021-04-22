package repository

import (
	"context"
	"fmt"
	"log"

	"server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetTodos() interface{}
	CreateTodo(i interface{}) interface{}
}

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) Repository {
	return TodoRepository{collection: collection}
}

func (r TodoRepository) GetTodos() interface{} {

	var results []*model.Todo
	findOptions := options.Find()

	cur, err := r.collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for cur.Next(context.TODO()) {
		var elem model.Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	cur.Close(context.TODO())
	return results
}

func (r TodoRepository) CreateTodo(i interface{}) interface{} {
	req, _ := i.(model.Todo)
	todo := model.Todo{
		Title:     req.Title,
		Completed: false,
	}
	insertResult, err := r.collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}
