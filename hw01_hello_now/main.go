package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeFormat := "2 Jan 2006 15:04:05"

	currentTime := time.Now()
	currentTime = currentTime.Round(0)

	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err == nil {
		exactTime = exactTime.Round(0)

		fmt.Println("current time:", currentTime.Format(timeFormat))
		fmt.Println("exact time:", exactTime.Format(timeFormat))

		os.Exit(0)
	}

	fmt.Println("error: ntp net error")
	os.Exit(1)
}
