package main

type Trup struct {
	name string
}

func NewTrup(name string) *Trup {
	return &Trup{name: name}
}

func (u *Trup) GetName() string {
	return u.name
}
