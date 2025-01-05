package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	api "simpleServer/api/grpc"
)

func main() {
	host := "localhost"
	port := "8080"
	id := int32(1)
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := api.NewPersonServiceClient(conn)
	personReq, err := client.AddPerson(context.Background(), &api.Person{})
	log.Println(personReq)
	if err != nil {
		log.Println(err)
		log.Println(personReq)
	}
	personReq, err = client.AddPerson(context.Background(), &api.Person{Id: &id, Name: "test"})
	log.Println(personReq)
	if err != nil {
		log.Println(err)
	}
	personReq, err = client.AddPerson(context.Background(), &api.Person{Id: &id, Name: "test"})
	log.Println(personReq)
	if err != nil {
		log.Println(err)
	}
	person, err := client.GetPerson(context.Background(), &api.GetPersonRequest{Id: &id})
	log.Println(person)
	if err != nil {
		log.Println(err)
	}
	id++
	person, err = client.GetPerson(context.Background(), &api.GetPersonRequest{Id: &id})
	log.Println(person)
	if err != nil {
		log.Println(err)
	}
}
