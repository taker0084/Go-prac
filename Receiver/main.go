package main

import "log"

type myStruct struct{
	FirstName string
}

//関数名の前に構造体へのポインターを書くと、構造体に関連づけた関数を定義できる → receiverという
//receiverには、ポインター型(s *Struct)と値型(s Struct)がある
//(m *myStruct)は、myStructのポインター型を受け取るreceiver
//(m myStruct)は、myStructの値型を受け取るreceiver
func (m *myStruct) printFirstName() string{
	return m.FirstName
}

func main(){
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}