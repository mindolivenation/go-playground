package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type userInfo struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	var user userInfo

	user.firstName = "Foo"
	user.lastName = "Bar"
	user.contactInfo.email = "foo.bar@baz.com"
	user.contactInfo.zipCode = 90210

	user2 := userInfo{
		firstName: "Barry",
		lastName:  "Boo",
		contactInfo: contactInfo{
			email:   "barry.boo@bar.com",
			zipCode: 90210,
		},
	}

	// & get memory address of the value this variable is pointing at.
	// * get the value this memory address is pointing at
	user.updateName("Jay", "Barr")

	user.print()
	user2.print()

	name := "Bill"
	fmt.Println(*&name)
	fmt.Println(&name)

	userSlice := []string{"jay", "barr"}
	updateUserSlice(userSlice)
	fmt.Println(userSlice)
}

func (ui *userInfo) updateName(newFirstName string, newLasttName string) {
	(*ui).firstName = newFirstName
	(*ui).lastName = newLasttName
}

func (ui userInfo) print() {
	fmt.Printf("%+v\n", ui)
}

func updateUserSlice(s []string) {
	s[0] = "brom"
}
