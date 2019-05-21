package main

import (
	"time"
	"log"
	"net/http"
	"os"
	"fmt"
)

func main() {
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("please define $URL")
	}
	rawDuration := os.Getenv("TICK")
	if rawDuration == "" {
		log.Fatal("please define $TICK")
	}

	duration, err := time.ParseDuration(rawDuration)
	if err != nil {
		log.Fatal("$TICK is not a valid format: See https://golang.org/pkg/time/#ParseDuration")
	}

	ticker := time.NewTicker(duration)

	go func() {
		for _ = range ticker.C {
			log.Printf("triggering post to %s", url)
			err := callUrl(url)
			if err != nil {
				log.Printf("error for requst: %v", err)
			}
		}
	}()

	select {}
}

func callUrl(url string) error {
	resp, err := http.Post(url, "application/json", nil)
    if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got non-ok from %s: %v", url, resp.Status)
	}

	return nil
}
