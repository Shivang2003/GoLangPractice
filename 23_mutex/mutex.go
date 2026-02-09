package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
}

func (p *Post) incViews(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views++

}

func main() {

	var wg sync.WaitGroup
	post := Post{views: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go post.incViews(&wg)
	}

	wg.Wait()
	fmt.Println("Total views: ", post.views)
}
