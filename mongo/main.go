package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name     string
	Surename string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// get session first, then acess DB and Collections
func main() {
	session, err := mgo.Dial("localhost")
	check(err)
	defer session.Close()

	c := session.DB("Hana").C("People")
	check(c.Insert(&Person{"gampol", "thiti"}, &Person{"kulrawee", "thiti"}))

	var results []Person
	check(c.Find(bson.M{"surename": "thiti"}).All(&results))
	fmt.Println("Results All: ", results)

}
