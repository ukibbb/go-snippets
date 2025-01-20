package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ukibbb/go-snippets/proto"
)

func main() {
	svc := NewMetricService(NewLoggingService(&priceFetcher{}))

	grpcClient, err := NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.FetchPrice(context.Background(), &proto.PriceRequest{Ticker: "ETH"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
			<-ticker.C

		}
	}()
	go MakeGRPCServerAndRun(":4000", svc)
	server := NewJSONAPIServer(":3000", svc)
	server.Run()

}
