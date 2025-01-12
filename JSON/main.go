package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}
func main(){
	myJSON := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "Black",
			"has_dog": false
		}
	]`
	var unMarshalled []Person

	err :=  json.Unmarshal([]byte(myJSON), &unMarshalled)
	if err != nil{
		log.Println("Error unMarshalling json", err)
	}

	log.Printf("unMarshalled %v",unMarshalled)

	//write json from struct
	var mySlice []Person

	var m1 Person
	m1.FirstName="Wally"
	m1.LastName="West"
	m1.HairColor="Red"
	m1.HasDog=false

	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName="Diana"
	m1.LastName="Prince"
	m1.HairColor="Black"
	m1.HasDog=false

	mySlice = append(mySlice, m2)

	newJson,err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil{
		log.Println("Error marshalling", err)
	}
	fmt.Println(string(newJson))
}