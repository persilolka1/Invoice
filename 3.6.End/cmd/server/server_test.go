package main

import (
	"backoffice/3.6.End/pb"
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestEnd(t *testing.T) {
	req := pb.RideRequest {
		Id: "9ed1c06186cf462d8421004f7c49835d",
    	DriverId: "007",
    	Location: &pb.Location{Lat: 51.4871871, Lng : -0.1266743},
    	PassengersIds: []string{"M", "Q"},
    	Start: timestamppb.New(time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)), 
    	Type: pb.RideType_POOL,
	}

	var srv Rides

	ctx := context.Background()

	resp, err := srv.End(ctx, &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Id != req.Id {
		t.Fatalf("respomse id %#v is not equal to request id %#v", resp.Id, req.Id)
	}
}

func TestStartE2E(t *testing.T) {
	list, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	server := createServer()

	go server.Serve(list) 

	port := list.Addr().(*net.TCPAddr).Port

	address := fmt.Sprintf("localhost:%v", port)

	creds := insecure.NewCredentials()

	c, err := grpc.DialContext(context.Background(), address, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	client := pb.NewRidesClient(c)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "api-key", "none")

	req := pb.RideRequest {
		Id: "9ed1c06186cf462d8421004f7c49835d",
    	DriverId: "007",
    	Location: &pb.Location{Lat: 51.4871871, Lng : -0.1266743},
    	PassengersIds: []string{"M", "Q"},
    	Start: timestamppb.New(time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)), 
    	Type: pb.RideType_POOL,
	}

	resp, err := client.Start(ctx, &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Id != req.Id {
		t.Fatalf("respomse id %#v is not equal to request id %#v", resp.Id, req.Id)
	}
}