package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("lang.json")
	if err != nil {
		log.Fatal(err)
	}
	// just dump it
	//io.Copy(os.Stdout, input)
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		//fmt.Printf("%v\n", lang)
		//fmt.Printf("%+v\n", lang)
		f(lang)
	}
}

//func count(name, url string) {
//	start := time.Now()
//	r, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("%s: %s", name, err)
//		return
//	}
//	n, _ := io.Copy(ioutil.Discard, r.Body)
//	r.Body.Close()
//	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
//}

func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s\n", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}

//func main() {
//	do(func(lang Lang) {
//		fmt.Printf("%v\n", lang)
//	})
//}

//func main() {
//	do(func(lang Lang) {
//		data, err := xml.MarshalIndent(lang, "", "  ")
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%s\n", data)
//	})
//}

//func main() {
//	start := time.Now()
//	do(func(lang Lang) {
//		count(lang.Name, lang.URL)
//	})
//	fmt.Printf("%2fs total\n", time.Since(start).Seconds())
//}

// parallel
//func main() {
//	do(func(lang Lang) {
//		go count(lang.Name, lang.URL)
//	})
//	time.Sleep(10 * time.Second)
//}

// concurrent
//func main() {
//	start := time.Now()
//	c := make(chan string)
//	n := 0
//	do(func(lang Lang) {
//		n++
//		go count(lang.Name, lang.URL, c)
//	})
//
//	for i := 0; i < n; i++ {
//		fmt.Print(<-c)
//	}
//	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
//}

// concurrent with timeout
func main() {
	start := time.Now()
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})
	timeout := time.After(2 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Print("Timed out\n")
			return
		}
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
