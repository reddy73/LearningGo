package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	userID    int
}

func main() {
	user := User{"John", "Doe", 1}
	fmt.Println("first name", user.firstName)
	fmt.Println("last name", user.lastName)
	user.lastName = "DoesWork"
	fmt.Println("last name", user.lastName)

	emp := struct {
		fName, lName string
		salary       float32
	}{
		fName:  "Jane",
		lName:  "Doe",
		salary: 999999.99,
	}
	fmt.Println(emp.salary)

	fmt.Println(user.userID)
	modifyMystruct(user)
	fmt.Println(user.userID)
	modifyMyStructViaPointer(&user)
	fmt.Println(user.userID)

}

func modifyMystruct(user User) User {
	user.userID = 22
	return user
}

func modifyMyStructViaPointer(user *User) {
	user.userID = 229
}
