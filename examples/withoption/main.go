package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/gotcha"
	"github.com/bxcodec/gotcha/cache"
)

func main() {
	cache := gotcha.NewWithOption(
		gotcha.NewOption().SetAlgorithm(cache.LFUAlgorithm).
			SetExpiryTime(time.Minute * 10).SetMaxSizeItem(100),
	)
	err := cache.Set("Kue", "Nama")
	if err != nil {
		log.Fatal(err)
	}
	val, err := cache.Get("Kue")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
