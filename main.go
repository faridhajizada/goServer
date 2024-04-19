package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name     string
	Age      uint16
	Money    int16
	AvgGrades, Happiness float64
	Hobbies  []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s age is: %d money is: %d average grades is: %.2f happiness is: %.2f hobbies are: %v", u.Name, u.Age, u.Money, u.AvgGrades, u.Happiness, u.Hobbies)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 1000, 4.2, 0.8, []string{"Football", "Guitar", "Programming"}}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	// fmt.Fprintf(w, "<h1>Welcome to the Home Page !</h1>")	

	tmpl, _ := template.ParseFiles("templates/homePage.html")
	tmpl.Execute(w, bob)
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Contact Page !")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the About Page !")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/about", aboutPage)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
