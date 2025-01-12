package main

import (
	"log"
)

//slice = 配列
func main(){
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)

	log.Println(numbers[6:9])


	name := []string{"one","seven","fish", "cat"}
	log.Println(name)
}