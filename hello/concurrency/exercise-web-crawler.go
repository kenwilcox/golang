package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeUrlMap struct {
	v   map[string]string
	mux sync.Mutex
}

func (c *SafeUrlMap) Set(key string, body string) {
	c.mux.Lock()
	c.v[key] = body
	c.mux.Unlock()
}

func (c *SafeUrlMap) Value(key string) (string, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.v[key]
	return val, ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlMap SafeUrlMap) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	urlMap.Set(url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if _, ok := urlMap.Value(u); !ok {
			wg.Add(1)
			Crawl(u, depth-1, fetcher, urlMap)
		}
	}
	return
}

var wg sync.WaitGroup

func main() {
	urlMap := SafeUrlMap{v: make(map[string]string)}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, urlMap)
	wg.Wait()

	for url := range urlMap.v {
		body, _ := urlMap.Value(url)
		fmt.Printf("found: %s %q\n", url, body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
