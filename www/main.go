package main

import (
	"fmt"
	"net/http"
)

// create struct
type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d he has money equal: %d", u.name, u.age, u.money)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.name = "Alex"
	fmt.Fprintf(w, "User name is: "+bob.name)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8000", nil)
}

func main() {
	//	//bob1 := User{"Bob", 25, -50, 4.2, 0.8}

	handleRequest()
}
