package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	redisApm "go.elastic.co/apm/module/apmgoredisv8/v2"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
	"log"
	"net/http"
	"os"
)

var redisClient = redis.NewClient(&redis.Options{
	DB:   0,
	Addr: os.Getenv("REDIS_URL"),
})

func serverHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Printf("ERROR [%+v] an error occurred", apm.TraceFormatter(ctx))
	redisClient.Set(ctx, "key", "value", 0)
	_, err := fmt.Fprint(w, "Hello")
	if err != nil {
		return
	}
}

func main() {
	redisClient.AddHook(redisApm.NewHook())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), apmhttp.Wrap(http.HandlerFunc(serverHandler)))
	if err != nil {
		return
	}
}
