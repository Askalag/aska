package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

func main() {
	var ch = make(chan string)
	//h <- "er"

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ch <- "lo"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()

	dir, _ := os.Getwd()

	feeds, err := obtainFeeds(dir + "/sandbox/feeds.json")

	print(feeds, err)

	fmt.Println("hello")
	close(ch)
}

func obtainFeeds(pathToFile string) ([]*Feed, error) {
	if pathToFile == "" {
		return nil, errors.New("error empty path to file")
	}

	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, errors.New("error during opening file")
	}

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, errors.New("error decode from json file")
	}

	return feeds, nil
}
