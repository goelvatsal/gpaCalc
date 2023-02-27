package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var totalCredits float64
	var totalPoints float64

	fmt.Println(`usage: ["class name" "cp/honors/ap" "credits" "grade"]`)
	fmt.Println("if a class is not cp/honors/ap, then don't add it in here!")
	fmt.Println()

	for in.Scan() {
		if in.Text() == "" {
			fmt.Printf("Total Points: %g\n", totalPoints)
			fmt.Printf("Total Credits: %g\n", totalCredits)
			fmt.Printf("Calculated GPA: %0.3f\n", totalPoints/totalCredits)
			return
		}

		inputs := strings.Fields(in.Text())
		credits, points := calcGPA(inputs[2], inputs[1], inputs[3])
		totalCredits += credits
		totalPoints += points
		//fmt.Println()
	}
}
