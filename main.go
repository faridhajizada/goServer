package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name     string
	age      uint16
	money    int16
	avgGrades, happiness float64
	hobbies  []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s age is: %d money is: %d average grades is: %.2f happiness is: %.2f hobbies are: %v", u.name, u.age, u.money, u.avgGrades, u.happiness, u.hobbies)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 1000, 4.2, 0.8, []string{"Football", "Guitar"}}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
	
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
