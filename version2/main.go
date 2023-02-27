package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var totalCredits float64
	var totalPoints float64

	var classes []string
	var gSlc []string
	var cSlc []string
	var cType []string
	var pSlc []float64

	fmt.Println(`usage: ["class name" "cp/honors/ap" "credits" "grade"]`)
	fmt.Println("after all classes have been entered, click enter twice.")
	//fmt.Println("if a class is not cp/honors/ap, then don't add it in here!")
	fmt.Println()

	for in.Scan() {
		if in.Text() == "" {
			//fmt.Println("Total Points:", totalPoints)
			//fmt.Println("Total Credits:", totalCredits)
			prettyPrint(classes, gSlc, cSlc, cType, pSlc)
			fmt.Printf("\nCalculated GPA: %0.3f\n", totalPoints/totalCredits)
			//fmt.Println("\nClasses Slice:", classes)
			//fmt.Println("Points Slice:", gSlc)
			//fmt.Println("Credits Slice:", cSlc)
			//fmt.Println("Class Type Slice:", cType)
			return
		}

		inputs := strings.Fields(in.Text())
		credit, cerr := strconv.ParseFloat(inputs[2], 64)
		if cerr != nil {
			log.Fatal()
		}

		grade, gerr := strconv.ParseFloat(inputs[3], 64)
		if gerr != nil {
			log.Fatal()
		}

		input := args{inputs[0], inputs[1], credit, grade}
		credits, points := input.calcGPA()

		classes = append(classes, inputs[0])
		gSlc = append(gSlc, inputs[3])
		cSlc = append(cSlc, inputs[2])
		cType = append(cType, inputs[1])
		pSlc = append(pSlc, points/credits)

		totalCredits += credits
		totalPoints += points

	}
}
