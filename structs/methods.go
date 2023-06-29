package main

import "fmt"

type Car struct {
	brand    string
	capacity int
}

func (car Car) DisplayCarInfo() string {
	return fmt.Sprintf("brand: %s capacity: %d", car.brand, car.capacity)
}

type User struct {
	FirstName, LastName, Email string
}

// DisplayUserInfo method definition on the 'User' struct type:
func (u User) DisplayUserInfo() string {
	return fmt.Sprintf("%s %s email is: %s", u.FirstName, u.LastName, u.Email)
}

func main() {
	spongeBob := User{
		FirstName: "SpongeBob",
		LastName:  "SquarePants",
		Email:     "spongebob@krustykrab.bb",
	}

	fmt.Println(spongeBob.DisplayUserInfo()) // here we call and print the method
}
