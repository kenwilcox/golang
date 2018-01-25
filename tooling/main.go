package main

import "fmt"

func main() {
	return
	var name string
	name = "Ken"
	fmt.Println("Hello ", name)
}

/*
package main

func main() {
data := struct {
Message string `json:"message"`
}{Message: "Hello world"}
err := json.NewEncoder(os.Stdout).Encode(data)
if err != nil {
log.Fatalln(err)
}
}
*/
