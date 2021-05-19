package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/moz5691/bookings/pkg/config"
	"github.com/moz5691/bookings/pkg/handlers"
	"github.com/moz5691/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// template cache


	// change this to true when in production
	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Running at port %s\n", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}


//import (
//	"encoding/json"
//	"fmt"
//	"log"
//)
//
//type Person struct {
//	FirstName string `json:"first_name"`
//	LastName string `json:"last_name"`
//}
//
//func main() {
//	myJson := `
//[
//	{
//		"first_name": "Kerry",
//		"last_name": "Johns"
//	},
//	{
//		"first_name": "Jerry",
//		"last_name": "Foreman"
//	}
//
//]
//`
//	var unmarshalled []Person
//
//	err := json.Unmarshal([]byte(myJson), &unmarshalled)
//
//	if err != nil {
//		log.Println("Error unmarshall json", err)
//	}
//	log.Printf("unmarshalled %v", unmarshalled)
//
//	var mySlice []Person
//
//	var m1 Person
//	m1.LastName = "Diana"
//	m1.FirstName = "Princess"
//	mySlice = append(mySlice, m1)
//
//	var m2 Person
//	m2.LastName = "Diana"
//	m2.FirstName = "Princess"
//	mySlice = append(mySlice, m2)
//
//	newJson, err := json.MarshalIndent(mySlice, "", "    ")
//
//	fmt.Printf("marshalled: %v", string(newJson))
//
//}


//import (
//	"github.com/moz5691/go-web/helpers"
//	"log"
//)
//
//const numPool = 100
//
//func CalculateValue(intChan chan int) {
//	randomNumber := helpers.RandomNumber(numPool)
//	intChan <- randomNumber
//}
//
//
//
//func main() {
//	intChan := make(chan int)
//	defer close(intChan)
//
//	go CalculateValue(intChan)
//
//	num := <- intChan
//
//	log.Println(num)
//}



//import (
//	"github.com/moz5691/go-web/helpers"
//	"log"
//)
//
//type Animal interface {
//	Says() string
//	NumberOfLegs() int
//}
//
//type Dog struct {
//	Name string
//	Breed string
//}
//
//type Gorilla struct {
//	Name string
//	Color string
//	NumberOfTeeth int
//}
//
//
//
//func main() {
//	dog := Dog{
//		Name: "Samson",
//		Breed: "German Sepherd",
//	}
//
//	PrintInfo(dog)
//
//	gorilla := Gorilla{
//		Name: "Big boy",
//		Color: "Black",
//		NumberOfTeeth: 36,
//	}
//
//	PrintInfo(gorilla)
//
//	var myVar helpers.SomeType
//	myVar.TypeName = "something"
//
//	log.Println(myVar)
//}
//
//
//func (d Dog) Says() string {
//	return "woof"
//}
//
//func (d Dog) NumberOfLegs() int {
//	return 4
//}
//
//func (g Gorilla) Says() string {
//	return "Shoooo"
//}
//
//func (g Gorilla) NumberOfLegs() int {
//	return 2
//}
//
//func PrintInfo(a Animal) {
//	log.Println("This animal says ", a.Says(), "and has ", a.NumberOfLegs())
//}
//
