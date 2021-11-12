package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main(){

		Port := os.Getenv("PORT")
		if Port == "" {
			Port = ":50005"
		}

		listen, err := net.Listen("tcp", Port)
		if err != nil {
			log.Fatalf("Couldnot listen @ %v :: %v", Port, err)
		}

		log.Println("Listening @"+Port)

		// gRPC server instance
		grpcserver := grpc.NewServer()

		//register ChatService
		/**
		cs := chatserver.ChatServer{}
		chatserver.RegisterServicesServer(grpcserver, &cs)
		*/

		//grpc listen and serve
		err = grpcserver.Serve(listen)
		if err != nil {
			log.Fatalf("Failed to start gRPC Server")
		}

}
