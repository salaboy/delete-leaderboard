package function

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"os"
)

var redisHost = os.Getenv("REDIS_HOST") // This should include the port which is most of the time 6379
var redisPassword = os.Getenv("REDIS_PASSWORD")
var redisTLSEnabled = os.Getenv("REDIS_TLS")
var redisTLSEnabledFlag = false

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	if redisTLSEnabled != "" && redisTLSEnabled != "false" {
		redisTLSEnabledFlag = true
	}
	var client *redis.Client

	if !redisTLSEnabledFlag {
		client = redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		})
	}

	response, err := client.FlushDB().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(res, "Status: " + response)

}
