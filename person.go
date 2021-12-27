package test

import "fmt"

type Person struct {
	id   int
	name string
}

func (p *Person) ID() int        { return p.id }
func (p *Person) Name() string   { return p.name }
func (p *Person) String() string { return fmt.Sprintf("%d%s", p.id, p.name) }
