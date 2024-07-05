package main

import (
	"api_gateway/api"
	"api_gateway/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Config{}
	conn, err := grpc.NewClient(":" + cfg.URL_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	r := api.NewRouter(conn)

	r.Run()
}
