package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/zahraakaraki/MyApp/database"
	pb "github.com/zahraakaraki/MyApp/protos"
	"google.golang.org/grpc"
)

var result []*pb.ReturnedType

type ClassServiceServer struct {
	pb.UnimplementedClassServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClassServiceServer(grpcServer, &ClassServiceServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8000: %v", err)
	}

}

func (s *ClassServiceServer) GetClasses(in *pb.Empty, stream pb.ClassService_GetClassesServer) error {
	log.Printf("recieved ")
	result = database.GetClassesFromDB()
	for _, res := range result {
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
