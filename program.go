package main

import (
	"fmt"
	"gobasic/calculator"
)

// type Porduct struct {
// 	name     string
// 	price    float64
// 	category string
// 	discount int
// }

func main() {
	// var name string
	// fmt.Print("Enter Your Name : ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("My name is :", name)
	// var number = []int{1, 2, 3}
	// number = append(number, 4)
	// var size = len(number)
	// fmt.Println(number, " Size ", size)
	// fmt.Println(number[:2], " Size ", size)
	// product := Porduct{name: "pen", price: 10.5, category: "ssss", discount: 5}
	// fmt.Println(product.name)
	fmt.Println(calculator.Add(10, 20))

}
