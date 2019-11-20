package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"esteth.net/magic/api/env"
	pb "esteth.net/magic/api/proto"
)

func main() {
	address := env.ReadEnv("HOST_ADDRESS")

	fmt.Printf("Connecting to %v...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10000000000))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMagicClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetAttractions(ctx, &pb.GetAttractionsRequest{})
	if err != nil {
		log.Fatalf("could not get attractions: %v", err)
	}

	for _, attraction := range r.GetAttraction() {
		fmt.Printf("%s\n", attraction.GetName())
	}
}
