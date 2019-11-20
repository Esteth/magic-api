package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"esteth.net/magic/api/data"
	"esteth.net/magic/api/env"
	pb "esteth.net/magic/api/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMagicServer
	db *data.DB
}

func (s *server) GetWaitTimes(ctx context.Context, in *pb.GetWaitTimesRequest) (*pb.GetWaitTimesResponse, error) {
	log.Printf("GetWaitTimes: %v", in)
	waitTimes, err := s.db.GetWaitTimes(in.GetParkId())
	if err != nil {
		return nil, errors.Wrap(err, "could not get wait times")
	}
	return &pb.GetWaitTimesResponse{
		WaitTime: waitTimes,
	}, nil
}

func (s *server) GetAttractions(ctx context.Context, in *pb.GetAttractionsRequest) (*pb.GetAttractionsResponse, error) {
	log.Printf("GetAttractions: %v", in)
	attractions, err := s.db.GetAttractions()
	if err != nil {
		return nil, errors.Wrap(err, "could not get attractions")
	}
	return &pb.GetAttractionsResponse{
		Attraction: attractions,
	}, nil
}

func main() {
	port, err := strconv.Atoi(env.ReadEnv("PORT"))
	if err != nil {
		log.Fatalf("could not parse PORT as int: %v", err)
	}
	dbHost := env.ReadEnv("DB_HOST")
	dbPort, err := strconv.Atoi(env.ReadEnv("DB_PORT"))
	if err != nil {
		log.Fatalf("could not parse DB_PORT as int: %v", err)
	}
	dbUser := env.ReadEnv("DB_USER")
	dbPassword := env.ReadEnv("DB_PASSWORD")
	dbName := env.ReadEnv("DB_NAME")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", port, err)
	}
	s := grpc.NewServer()

	log.Println("connecting to database...")
	db, err := data.Dial(
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName))
	if err != nil {
		// TODO(adamcopp): Cope with database not being available temporarily.
		log.Fatalf("could not connect to database: %v", err)
	}
	pb.RegisterMagicServer(s, &server{
		db: db,
	})
	fmt.Printf("listening on port %d...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
