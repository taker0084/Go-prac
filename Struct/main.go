package main

import (
	"log"
	"time"
)


var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

//Goはオブジェクト指向言語でないため、publicとprivateがない → 最初を大文字にすると外部から参照可能(関数や変数など)
var Special string

func main(){
	user := User{
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 555 555-1212",
		Age: 30,
		BirthDate: time.Now(),
	}
	log.Println(user.FirstName, user.LastName, "BirthDate", user.BirthDate, "Age", user.Age, "PhoneNumber", user.PhoneNumber)
}