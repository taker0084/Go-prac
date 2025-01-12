package main

import "log"

func main(){
	for i:=0;i<=10;i++ {
		log.Println(i)
	}
	//Goではforeach文はない → for...rangeでかく
	animals := []string{"dog","fish","horse","cow","cat"}
	//前の引数はloop数,後ろはアイテム
	for _,animal := range animals{
		log.Println(animal)
	}

	Animals := make(map[string]string)
	Animals["dog"] = "Fido"
	Animals["cat"] = "Fluffy"
	for AnimalType, animal := range Animals{
		log.Println(AnimalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"

	for i, l := range firstLine{
		log.Println(i,":",l)
	}

	type User struct{
		FirstName string
		LastName string
		Email string
		Age int
	}
	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 30})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 30})
	users = append(users, User{"Alex", "Anderson", "Alex@Anderson.com", 30})

	for _,l := range users{
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}