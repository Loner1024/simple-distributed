package main

import (
	"context"
	"fmt"
	slog "log"
	"simple-distributed/log"
	"simple-distributed/service"
)

const (
	host = "localhost"
	port = "4000"
)

func main() {
	log.Run("./distributed.log")
	ctx, err := service.Start(context.Background(), "Log Service", host, port, log.RegisterHandlers)
	if err != nil {
		slog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
