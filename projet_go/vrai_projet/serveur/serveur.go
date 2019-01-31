package main

import (
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Repondre(ctx context.Context, in *pb.msgEnvoye) (*pb.reponse, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.reponse{Message: "Hello " + in.Name}, nil
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
