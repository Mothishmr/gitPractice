package main
import "fmt"

func process(ch chan int){
	for i:= 0 ; i<10 ; i++{
		ch <- i
	} 
	close(ch)
}

func main(){
	ch := make(chan int)
	go process(ch)
for val := range ch{
	fmt.Print(val)
}

}