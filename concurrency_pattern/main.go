package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mutext := sync.Mutex{}

	USD := 100000000
	files := []string{"user1.json", "user2.json", "user3.json", "user4.json"}

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			mutext.Lock()
			defer mutext.Unlock()
			fmt.Printf("translate %d USD to IDR -> %s\n", USD, file)
		}(files[i])
	}

	wg.Wait()
}
