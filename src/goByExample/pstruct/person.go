package pstruct

import "fmt"

type  Person struct {
	Name string
	Age int
}
func (p *Person) Eat(){
	fmt.Println(p.Name+"eating...")
}
