package main

import (
	"context"
	"fmt"
	"time"
	"tutor/go-mongodb/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("SERVICE START")

	Create()
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Age       string             `bson:"age" json:"age"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func Create() {
	// mongo connection
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		table = db.MgoCollection("users", client)
	)

	_, err := table.InsertOne(ctx, User{
		Name:      "Andi Wibowo",
		Age:       "26",
		CreatedAt: time.Now().In(time.UTC),
	})
	if err != nil {
		fmt.Println("error insert to mongodb : ", err)
	}

	fmt.Println("Successfully inserted")
}
