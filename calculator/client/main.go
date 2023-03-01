package main

import (
	"log"

	"github.com/MangoSteen0903/go-grpc-tour/calculator/client/methods"
	pb "github.com/MangoSteen0903/go-grpc-tour/calculator/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	//addResult := methods.ExecuteAdd(c)

	//log.Printf("Caculate result : %d", addResult)

	//methods.ExecutePrime(c)

	//methods.ExecuteAvg(c)

	//methods.ExecuteMax(c)

	sqrtResult, err := methods.ExecuteSqrt(c, 10)

	if err != nil {
		log.Println("Error Occured")
	} else {
		log.Printf("Server: Sqrt Result : %f", sqrtResult)
	}

}
