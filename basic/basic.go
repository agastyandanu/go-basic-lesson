package main

import "fmt"

func main() {

	// NUMBER
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 1)
	fmt.Println("tiga = ", 3.5)

	//BOOLEAN
	fmt.Println("true", true)
	fmt.Println("false", false)

	//STRING
	fmt.Println("Danoe")
	fmt.Println(len("Danoeee")) // len() function is used to count total character/data (length in javascript)
	fmt.Println("Danoe"[1]) // to take character(as binary) based on its index


	//VARIABLE
	var name string
	name = "Danu Agastyan" // set value of the variable
	var age int
	age = 23

	var hobby = "Music, Football, Travel" // make variable without declare its data-type
	country := "Indonesia" //make variable without declare "var"

	var ( // declare some variables once
		firstName = "Danu"
		lastName = "Agastyan"
		currentAge = 23
	)
	fmt.Println("My name is", name)
	fmt.Println("I am ", age, "years old")
	fmt.Println("My hobby:", hobby)
	fmt.Println("Country:", country)
	fmt.Println("Name:", firstName, lastName, "Age:", currentAge)

	

	//CONSTANT (UNEDITABLE DATA VARIABLE)
	const (
		constName = "Danu Agastyan"
		constAge = 23
	)
	fmt.Println(constName, constAge, "years old")



	//DATA CONVERTION
	var stringName = "Danu"
	var e = stringName[3]
	var eString = string(e)
	fmt.Println("This is already converted to be string", eString)



	//DATA TYPE ALIAS
	type stringDanoe string
	type intDanoe int
	type boolDanoe bool	
	var married boolDanoe = false
	fmt.Println("Marriage Status:", married)



	//ARITHMETIC
	var a = 7
	var b = 10
	var i = 1
	i += 2 //add itself to new value
	fmt.Println(i)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)



	//ARRAY
	users := [...]string { //[...] defines that array has unknown/unlimited capacity(length)
		"Danu",
		"Danoe",
		"Agastyan",
	}
	months := [12]string {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"Actober",
		"November",
		"December",
	}

	fmt.Println(users)
	fmt.Println("Length of array", len(users)) // get length of the array
	fmt.Println("Capacity of the array", cap(months)) // get capacity of the array
	fmt.Println("slice", months[3:7]) //slice (get array-data starts from index 3rd to before 7th = data-index 3, 4, 5 & 6)


	

}