package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/currency/protos/currency"
	"github.com/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	cs := server.NewCurrency(log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, cs)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to listen", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}

// test with grpcurl
// grpcurl -d '{\"base\": \"GBP\", \"destination\": \"USD\"}' -plaintext  localhost:9092 Currency/GetRate
