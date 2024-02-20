package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"backoffice/3.6.End/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := grpc.NewServer()

	var u Rides

	pb.RegisterRidesServer(srv, &u)

	reflection.Register(srv)

	log.Printf("info: server ready on %s", addr)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func (r *Rides) Start(ctx context.Context, req *pb.RideRequest) (*pb.RideStartResponse, error) {
	// TODO: Validate req

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no_metadata")
	}

	log.Printf("info: apikey %v", md["api-key"])

	resp := pb.RideStartResponse{
		Id: req.Id,
	}

	// TODO: Work (insert to database ...)
	return &resp, nil
}

func (r *Rides) End(ctx context.Context, req *pb.RideRequest) (*pb.RideEndResponse, error) {
	// TODO: Validate req

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no_metadata")
	}

	log.Printf("info: apikey %v", md["api-key"])

	resp := pb.RideEndResponse{
		Id: req.Id,
	}

	// TODO: Work (insert to database ...)
	return &resp, nil
}

func (r *Rides) Location(stream pb.Rides_LocationServer) error {
	count := int64(0)
	driverId := ""

	for {
		req, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return status.Errorf(codes.Internal, "can not read")
		}

		count++
		driverId = req.DriverId
	}

	response := pb.LocationResponse {
		Count: count,
		DriverId: driverId,
	}

	return stream.SendAndClose(&response)
}

type Rides struct {
	pb.UnimplementedRidesServer
}