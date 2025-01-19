package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewMetricService(NewLoggingService(&priceFetcher{}))
	server := NewJSONAPIServer(":3000", svc)
	server.Run()
	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
