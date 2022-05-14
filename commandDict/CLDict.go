package main

import (
	"Go_Project/commandDict/api"
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	//use system wait group to process 2 translation api together
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	api.CaiyunQuery(word, &waitGroup)
	api.Query360(word, &waitGroup)
	waitGroup.Wait()
}
