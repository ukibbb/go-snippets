package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

//func grpcf() {
//	svc := NewMetricService(NewLoggingService(&priceFetcher{}))
//
//	grpcClient, err := NewGRPCClient(":4000")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	go func() {
//		ticker := time.NewTicker(time.Second)
//		for {
//			time.Sleep(3 * time.Second)
//			resp, err := grpcClient.FetchPrice(context.Background(), &proto.PriceRequest{Ticker: "ETH"})
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			fmt.Printf("%+v\n", resp)
//			<-ticker.C
//
//		}
//	}()
//	go MakeGRPCServerAndRun(":4000", svc)
//	server := NewJSONAPIServer(":3000", svc)
//	server.Run()
//}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s := NewStore(NewRedisCache(client))

	s.cache.Set(1, "value")
	val, err := s.Get(1)

	fmt.Println(val, err)
}
