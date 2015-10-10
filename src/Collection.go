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

	anotherSlice :=make([]int,100)
	fmt.Println(anotherSlice)

	myMap :=make(map[int] string)
	fmt.Println(myMap)

	myMap[0]= "Jalpesh"
	fmt.Println(myMap)
}