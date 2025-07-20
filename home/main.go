package main

import "fmt"

func printItems[T any](arr []T){
for _, val := range arr{
		fmt.Println(val)
	}
}

func getSum(nums []int) int{
	var sum = 0

	for _,  val := range nums{
		sum += val
	}
	return sum
}


func main(){


	var movies = []string {"money ball" , "f1" , "ford vs ferrari"}
	var ratings = []int {10 , 10 , 10}

	var sum = getSum(ratings)
	
	// var name string

	// fmt.Scan(&name)


	fmt.Println(sum)
	// fmt.Println("YOUR na<e : " , name )
	movies = append(movies , "Fight club")
	printItems(movies)

	printItems(ratings)
	fmt.Println("hello world version 3")

	fmt.Println("bro what is this");
	//  MAPS

	var mpp = map[string] int {
		"mr" : 20,
		"mj" : 30,
		"ms" : 40,
	}
  
	for name , age := range mpp{
		fmt.Println(name , " " , age)
	}


	for i:= 0 ; i<  n ; i++{
		fmt.Println("CHANGES change chetshanna ga ra lamdiiiiiiiiiii DONE");
	}

	fmt.Println("HEAD -> master , origin/master , origin/HEAD")



	var mpp := map[string]int{
		"new tutorial" : 2,
	}

}	