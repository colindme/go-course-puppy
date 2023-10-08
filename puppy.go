package puppy

import (
	"fmt"

	dog "github.com/colindme/go-course-dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("This is from version 0.1.0")
}
