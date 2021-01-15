package main

import (
	"VanGogh/conf"
	"VanGogh/router"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path.")
)

func main() {
	pflag.Parse()
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("mode"))
	g := router.Setup()

	// Check server response
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The server has no response, or it might took too long to start up.", err)
		}
		log.Print("The server has been deployed successfully.")
	}()

	// Start HTTP Server & Log
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	log.Printf(http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("port")), g).Error())
}

func pingServer() error {
	count := viper.GetInt("max_ping_count")
	for i := 0; i < count; i++ {
		// Ping the server by sending a GET request to `/health`
		resp, err := http.Get(viper.GetString("url") + ":" + viper.GetString("port") + "/sc/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the server, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the server")
}
