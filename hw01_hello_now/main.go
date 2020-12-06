package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const timeFormat = "2 Jan 2006 15:04:05"

	currentTime := time.Now().Round(0)

	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalln("error: ntp net error")
	}

	exactTime = exactTime.Round(0)

	fmt.Println("current time:", currentTime.Format(timeFormat))
	fmt.Println("exact time:", exactTime.Format(timeFormat))

	os.Exit(0)
}
