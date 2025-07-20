package main

import "fmt"


type Car struct{
	name string
	model int
	Engine
}

type Bike struct{
	name string
	Engine
}

type Engine struct{
	engName string
	cap int
}


func main(){
	fmt.Println("BRO")

	var benz Car = Car{"mercedes" , 100 , Engine{"e10" , 10}}


	

	fmt.Print(benz)

}

