package main
import (
	"fmt"
)

func main()  {

	myArray := [3] int{55,50,45}
	fmt.Println(myArray)
	fmt.Println(len(myArray))

	mySlice :=myArray[:]
	mySlice = append(mySlice,200)
	fmt.Println(mySlice)
}