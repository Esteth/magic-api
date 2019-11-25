package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"esteth.net/magic/api/data"
	"esteth.net/magic/api/env"
	pb "esteth.net/magic/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMagicServer
	db *data.DB
}

func (s *server) GetWaitTimes(_ context.Context, in *pb.GetWaitTimesRequest) (*pb.GetWaitTimesResponse, error) {
	log.Printf("GetWaitTimes: %v", in)
	waitTimes, err := s.db.GetWaitTimes(in.GetParkId())
	if err != nil {
		return nil, fmt.Errorf("could not get wait times: %w", err)
	}
	return &pb.GetWaitTimesResponse{
		WaitTime: waitTimes,
	}, nil
}

func (s *server) GetAttractions(_ context.Context, in *pb.GetAttractionsRequest) (*pb.GetAttractionsResponse, error) {
	log.Printf("GetAttractions: %v", in)
	attractions, err := s.db.GetAttractions()
	if err != nil {
		return nil, fmt.Errorf("could not get attractions: %w", err)
	}
	return &pb.GetAttractionsResponse{
		Attraction: attractions,
	}, nil
}

func main() {
	port, err := strconv.Atoi(env.ReadEnvOr("PORT", "5002"))
	if err != nil {
		log.Fatalf("could not parse PORT as int: %v", err)
	}
	dbHost := env.ReadEnvOr("PGHOST", "localhost")
	dbPort, err := strconv.Atoi(env.ReadEnvOr("PGPORT", "5432"))
	if err != nil {
		log.Fatalf("could not parse DB_PORT as int: %v", err)
	}
	dbUser := env.ReadEnvOr("PGUSER", "postgres")
	dbPasswordFile := env.ReadEnv("PGPASSFILE")
	dbPassword, err := ioutil.ReadFile(dbPasswordFile)
	if err != nil {
		log.Fatalf("could not read PGPASSFILE %s: %v", dbPasswordFile, err)
	}
	dbName := env.ReadEnvOr("PGDATABASE", "postgres")

	hasReflection, err := strconv.ParseBool(env.ReadEnvOr("ALLOW_REFLECTION", "false"))
	if err != nil {
		log.Fatalf("could not parse ALLOW_REFLECTION environment variable as boolean: %v", err)
	}

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
	if hasReflection {
		reflection.Register(s)
	}
	fmt.Printf("listening on port %d...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
