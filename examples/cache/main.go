package main

import (
	"context"
	"fmt"
	"gosdk/cache"
	"log"
	"reflect"
	"time"
)

func main() {
	cache := cache.NewSimpleRedisCache("redis://localhost:6379")
	data := map[string]any{"someKey": "value"}
	err := cache.Set(context.Background(), "k1:test1", data, 1*time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]any
	err = cache.Get(context.Background(), "k1:test1", &res)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(data, res) {
		log.Fatal("Not equal")
	}
	fmt.Println(data)
	fmt.Println(res)
}
