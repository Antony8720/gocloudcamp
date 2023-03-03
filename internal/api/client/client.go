package client

import (
	"context"
	"log"
	"gocloudcamp/internal/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewApiClient(conn)
	res, err := client.Play(context.Background(), &proto.Request{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

}
