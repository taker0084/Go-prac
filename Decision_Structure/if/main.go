package main

import "log"

func main(){
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	}else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is cat")
	}else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTrue2 := false
	if myNum >99 && !isTrue2{
		log.Println("myNum is greater than 99 and isTrue is set to true")
	}else if myNum<100 && isTrue{
		log.Println(1)
	}else if myNum == 101 || !isTrue2{
		log.Println("2")
	}else if myNum > 100 && isTrue==false{
		log.Println("3")
	}
}