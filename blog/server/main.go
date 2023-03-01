package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MangoSteen0903/go-grpc-tour/calculator/server/util"
	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
	"google.golang.org/grpc"
)

const PORT = "3000"

func main() {
	port := fmt.Sprintf(":%s", PORT)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &util.CalculatorServer{})

	log.Printf("Server Listening on http://localhost:%s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
