package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/zahraakaraki/MyApp/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewClassServiceClient(conn)

	runGetClasses(client)
}

func runGetClasses(client pb.ClassServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetClasses(ctx, req)

	if err != nil {
		log.Fatalf("%v.GetClasses() = %v", client, err)
	}

	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetClasses() = %v", client, err)
		}
		log.Printf("Class: %v", row)
	}
}
