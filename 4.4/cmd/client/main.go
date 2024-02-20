package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"backoffice/3.6.End/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	addr := ":9292"

	creds := insecure.NewCredentials()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error: can't connect - %s", err)
	}

	defer conn.Close()

	log.Printf("info: connected to %s", addr)
	
	c := pb.NewRidesClient(conn)

	fmt.Println(c)

	req := pb.RideRequest {
		Id: "9ed1c06186cf462d8421004f7c49835d",
    	DriverId: "007",
    	Location: &pb.Location{Lat: 51.4871871, Lng : -0.1266743},
    	PassengersIds: []string{"M", "Q"},
    	Start: timestamppb.New(time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)), 
    	Type: pb.RideType_POOL,
	}

	reqCtx, reqCancel := context.WithTimeout(context.Background(), time.Second)

	defer reqCancel()

	reqCtx = metadata.AppendToOutgoingContext(reqCtx, "api-key", "s3cr3t")

	resp, err := c.Start(reqCtx, &req)
	if err != nil {
		log.Fatalf("error: can't retrieve response - %s", err)
	}

	fmt.Println(resp)

	req2Ctx, req2Cancel := context.WithTimeout(context.Background(), time.Second)

	defer req2Cancel()

	req2Ctx = metadata.AppendToOutgoingContext(req2Ctx, "api-key", "s2s2")

	resp2, err := c.End(req2Ctx, &req)
	if err != nil {
		log.Fatalf("error: can't retrieve response  2- %s", err)
	}

	fmt.Println(resp2)
}