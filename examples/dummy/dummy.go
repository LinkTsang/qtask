package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	seconds := flag.Int("seconds", 1, "Dummy task duration in seconds")
	flag.Parse()

	fmt.Printf("Dummy task for %d seconds...\n", *seconds)
	time.Sleep(time.Duration(*seconds) * time.Second)
	fmt.Println("Done.")
}
