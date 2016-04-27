package main

import (
	"fmt"
	"time"
)

type Member interface {
	GetName()
	GetDetails()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method
func (p Person) GetName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method
func (p Person) GetDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

type Admin struct {
	Person //type embedding for composition
	Roles  []string
}

//overrides GetDetails
func (a Admin) GetDetails() {
	//Call person GetDetails
	a.Person.GetDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type User struct {
	Person //type embedding for composition
	Skills []string
}

//overrides GetDetails
func (u User) GetDetails() {
	//Call person GetDetails
	u.Person.GetDetails()
	fmt.Println("Skills:")
	for _, v := range u.Skills {
		fmt.Println(v)
	}
}

type Team struct {
	Name, Description string
	Members           []Member
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team :%s  - %s\n", t.Name, t.Description)
	fmt.Println("Deteails of the team members:")
	for _, v := range t.Members {
		v.GetName()
		v.GetDetails()
	}
}

func main() {
	alex := Admin{
		Person{"Alex", "John", time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC), "alex@email.com", "New York"},
		[]string{"Create Team", "Create Task"},
	}
	shiju := User{
		Person{"ravi", "kumar", time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC), "shiju@email.com", "Kochi"},
		[]string{"Go", "Docker", "Kubernetes"},
	}
	chris := User{
		Person{"Chris", "Martin", time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC), "chris@email.com", "Santa Clara"},
		[]string{"Go", "Docker"},
	}
	team := Team{
		"Go",
		"Golang CoE",
		[]Member{alex, shiju, chris},
	}
	//Get details of the Team
	team.GetTeamDetails()
}
