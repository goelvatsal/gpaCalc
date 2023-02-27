package main

import (
	"log"
	"strconv"
	"strings"
)

func calcGPA(creds string, ctype string, grade string) (recreds, repoints float64) {
	var subtractor float64
	var gpa float64
	var totalCreds float64
	var totalPoints float64

	switch ctype {
	case strings.ToLower("cp"):
		subtractor = 1.0
	case strings.ToLower("honors"):
		subtractor = 0.5
	case strings.ToLower("ap"):
		subtractor = 0.0
	default:
		log.Panic("err on line 31")
	}

	credits, err := strconv.ParseFloat(creds, 64)
	if err != nil {
		log.Panic(err)
	}
	switch i, _ := strconv.ParseFloat(grade, 64); {
	case i >= 98 && i <= 100:
		gpa = 5.0 - subtractor
	case i >= 93 && i <= 100:
		gpa = 4.7 - subtractor
	case i >= 90 && i <= 92:
		gpa = 4.5 - subtractor
	case i >= 87 && i <= 89:
		gpa = 4.3 - subtractor
	case i >= 83 && i <= 86:
		gpa = 4.0 - subtractor
	case i >= 80 && i <= 82:
		gpa = 3.7 - subtractor
	case i >= 77 && i <= 79:
		gpa = 3.4 - subtractor
	case i >= 73 && i <= 76:
		gpa = 3.0 - subtractor
	case i >= 70 && i <= 72:
		gpa = 2.7 - subtractor
	case i >= 65 && i <= 69:
		gpa = 2.0 - subtractor
	case i >= 0 && i <= 64:
		gpa = 0
	}
	totalCreds += credits
	points := gpa * credits
	totalPoints += points

	//fmt.Println("class gpa:", gpa)
	//fmt.Println("points:", points)
	//fmt.Println("credits:", credits)
	return totalCreds, totalPoints
}
