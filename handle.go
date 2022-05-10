package function

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"os"
)

var redisHost = os.Getenv("REDIS_HOST") // This should include the port which is most of the time 6379
var redisPassword = os.Getenv("REDIS_PASSWORD")

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: redisPassword,
		DB:       0,
	})

	response, err := client.FlushDB().Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(res, "Status: " + response)

}
