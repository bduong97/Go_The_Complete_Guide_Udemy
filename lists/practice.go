package main

import "fmt"

func main() {
	// 1
	question1 := [3]string{"hobby1", "hobby2", "hobby3"}
	fmt.Println(question1)

	// -------------------
	//  2
	fmt.Println(question1[0])
	fmt.Println(question1[1:3])

	// -------------------
	//  3
	question3a := question1[0:2]
	question3b := question1[:2]
	fmt.Println(question3a)
	fmt.Println(question3b)

	// -------------------
	// 4
	question3a = question3a[1:3]
	fmt.Println(question3a)

	// -------------------
	// 5
	question5 := []string{"not sure", "trying"}
	fmt.Println(question5)

	// -------------------
	// 6
	question5[1] = "oh well"
	question5 = append(question5, "taco")
	fmt.Println(question5)

	// -------------------
	// 7

	products := []Product{}
	for i := 0; i < 2; i++ {
		products = append(products, Product{
			title: "queso",
			id: "cheese",
			price: float64(i),
		})
	}
	fmt.Println(products)
	products = append(products, Product{
		title: "taco",
		id: "asada",
	})
	fmt.Println(products)

}

type Product struct {
	title string
	id    string
	price float64
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
