package cracking_the_coding_interview

import (
	"fmt"
	"testing"
)

type AnimalShelf struct {
	animals []*Animal
}

type Animal struct {
	Index int
	Type  int
}

func Constructor2() AnimalShelf {
	return AnimalShelf{
		animals: make([]*Animal, 0),
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if len(animal) != 2 {
		return
	}

	this.animals = append(this.animals, &Animal{Index: animal[0], Type: animal[1]})
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.animals) == 0 {
		return []int{-1, -1}
	}

	res := []int{this.animals[0].Index, this.animals[0].Type}
	this.animals = this.animals[1:]
	return res
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.animals) == 0 {
		return []int{-1, -1}
	}

	for i, animal := range this.animals {
		if animal.Type == 1 {
			res := []int{animal.Index, animal.Type}
			this.animals = append(this.animals[:i], this.animals[i+1:]...)
			return res
		}
	}

	return []int{-1, -1}
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.animals) == 0 {
		return []int{-1, -1}
	}

	for i, animal := range this.animals {
		if animal.Type == 0 {
			res := []int{animal.Index, animal.Type}
			this.animals = append(this.animals[:i], this.animals[i+1:]...)
			return res
		}
	}

	return []int{-1, -1}
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */

func Test_AnimalShelf(t *testing.T) {
	obj := Constructor2()

	print := func() {
		fmt.Printf("obj: ")
		for _, animal := range obj.animals {
			fmt.Printf("%v ", animal)
		}
		fmt.Printf("\n")
	}

	{
		obj.Enqueue([]int{0, 0})
		obj.Enqueue([]int{1, 0})
		print()

		getCat1 := obj.DequeueCat()
		print()
		fmt.Printf("get cat1: %v, want: %v\n", getCat1, []int{0, 0})

		getDog1 := obj.DequeueDog()
		print()
		fmt.Printf("get dog1: %v, want: %v\n", getDog1, []int{-1, -1})

		getAny1 := obj.DequeueAny()
		print()
		fmt.Printf("get any1: %v, want: %v\n", getAny1, []int{1, 0})
	}
}

func Test_trimSlice(t *testing.T) {
	n := []int{1, 2, 3}

	for i, v := range n {
		if v == 1 {
			n = append(n[:i], n[i+1:]...)
		}
	}
	fmt.Println(n)
}
