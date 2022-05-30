package creational_petterns

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p *Person) setFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) setLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) setAge(age int) {
	p.age = age
}

func NewPersonaBuilder() *Person {
	p := &Person{}
	p.setFirstName("Amir")
	p.setLastName("Last")
	p.setAge(20)
	return p
}

// use this design pattern when building object is complex
