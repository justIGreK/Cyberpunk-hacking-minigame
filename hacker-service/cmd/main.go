package main

import (
	"context"
	"fmt"
	"hacker-service/cmd/config"
	"hacker-service/cmd/handler"
	mongorepo "hacker-service/internal/respository/mongo"
	"hacker-service/internal/service"
	"log"
	"net/http"
	"os"
)

// @title Utility to solve matrices of matrix_service

// @host localhost:8081
// @BasePath /

func main() {
	config.LoadEnv()
	ctx := context.Background()
	db := mongorepo.CreateMongoClient(ctx)
	repo := mongorepo.NewMatrixRepo(db)
	service := service.NewMatrixService(repo)
	handler := handler.NewHandler(service)
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Printf("hacker-service starts listening :%v \n", port)
	http.ListenAndServe(port, handler.InitRoutes())
}
