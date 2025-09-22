package main

type Child struct {
	name string
	age  int
}

func (c *Child) setName(name string) {
	c.name = name
}

func (c Child) getName() string {
	return c.name
}

func main() {
	c := &Child{}
	c.setName("Alice")
	println(c.name)

	c.getName()
}
