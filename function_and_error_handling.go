// Error Handling In Fuction
package main

import "fmt"

func divide(a, b int) int { // The issue is that our answer should be in float but comes in int, now solve
	return a / b
}

// Now by Solving float value issue
func dvided(a, b float64) float64 {
	return a / b
}

// Now making function that is divided by zero
func division(a, b float64) float64 { // if number is divided by 0 answer comes in infinity
	return a / b
}

// We don't want to get answer in infinity now we handle this
func inf_division(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// Now we work for error handling so
func err_division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator Must Not Be Zero:")
	}
	return a / b, nil // Therefore we use nil because this part is completly correct
}

func main() {
	fmt.Println("Started Error Handling In Function:")
	fmt.Println("The Division Of Two Number Is:", divide(10, 4))
	fmt.Println("Now The Answer would become in float after solving issue:", dvided(10, 4))
	fmt.Println("The Answer After Divided By '0':", division(10, 0))  // Answer Comes In Infinity
	fmt.Println("By Handling Infinity Problem:", inf_division(10, 4)) // Function Works Properly
	// The Lower Given Part Is A Part Of err_division function
    // The Lower Given "_" represents the Declaration Of "error" 
    // we use "_" on any place where function returns some value and we want to ignore them
    // you can use "_" where you want to ignore the output somewhere
	ans, _ := err_division(10, 0)
	fmt.Println("The Answer after solving by error", ans)
}
