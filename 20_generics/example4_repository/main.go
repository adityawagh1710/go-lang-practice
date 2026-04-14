package main

import "fmt"

type Repository[T any] struct {
	data []T
}

func (r *Repository[T]) Add(item T) {
	r.data = append(r.data, item)
}

func (r *Repository[T]) GetAll() []T {
	return r.data
}

func main() {
	userRepo := Repository[string]{}
	userRepo.Add("Aditya")
	userRepo.Add("Rahul")
	fmt.Println(userRepo.GetAll())

}
