package helper

import ("fmt"
		"reflect")

func Hello() {

	result := true
	for result {
	var name string
	fmt.Println("Hello, World!")
	fmt.Println("Enter your name (letters only):")
	fmt.Scanln(&name)
	inputResult := ValidateID(name)
	if inputResult == false {
		fmt.Println("Invalid ID")
		result = false
    } else {
		fmt.Println("Valid ID")

}
}
}

func ValidateID(id string) bool{

	length := len(id)>= 2 

	isString := reflect.TypeOf(id).Kind() == reflect.String 

	noNum := true
	for _,character := range id {
		if character >= '0' && character <= '9' {
			noNum = false
        }
	}

    return length && isString && noNum

}
