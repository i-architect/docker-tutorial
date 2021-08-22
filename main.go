package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	countHandler := func(w http.ResponseWriter, req *http.Request) {
		result, err := rdb.Incr(context.Background(), "counter").Result()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(w, strconv.FormatInt(result, 10)+"\n")
	}

	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
