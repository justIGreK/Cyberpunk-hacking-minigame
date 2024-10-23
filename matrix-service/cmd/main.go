package main

import (
	"context"
	"fmt"
	"log"
	"matrix-service/cmd/config"
	"matrix-service/cmd/handler"
	mongorepo "matrix-service/internal/respository/mongo"
	"matrix-service/internal/service"
	"net/http"
	"os"
)

// @title Cyberpunk 2077 hacking mini game

// @host localhost:8080
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
