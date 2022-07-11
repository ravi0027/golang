package main

//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"time"
//)
//
//const uri = "mongodb+srv://jayakumar:<J@yakumar23>@cluster1.swfaa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
//
//func main() {
//	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
//	clientOptions := options.Client().
//		ApplyURI("mongodb+srv://jayakumar:<J@yakumar23>@cluster1.swfaa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
//		SetServerAPIOptions(serverAPIOptions)
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	client, err := mongo.Connect(ctx, clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = client.Connect(ctx)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(client.Connect(ctx))
//}
