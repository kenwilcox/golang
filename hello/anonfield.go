package main

// borrowed from here:
// http://spf13.com/post/is-go-object-oriented/
import "fmt"

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Citizen struct {
	Country string
	Person  // all person's properties are included in the Citizen "object"
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (c *Citizen) Talk() {
	fmt.Println("Hello, my name is", c.Name, "and I'm from", c.Country)
}

type Human interface {
	Talk()
}

func SpeakTo(h Human) {
	h.Talk()
}

func main() {
	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()
	c.Person.Talk()
	c.Nationality()

	p := Person{Name: "Dave"}
	SpeakTo(&p)
	SpeakTo(&c)
}
