package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// create struct
type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

// Отображение пользователей
func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d he has money "+
		"equal: %d", u.Name, u.Age, u.Money)
}

// Вызов функции пользователей
func (u *User) setNewName(newName string) {
	u.Name = newName
}

// Домашняя страница
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())

	// Подключение шаблона HTML
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

// Подключение локалхоста, функции home_page и протокола
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8000", nil)
}

// Основная функция
func main() {
	//	//bob1 := User{"Bob", 25, -50, 4.2, 0.8}

	handleRequest()
}
