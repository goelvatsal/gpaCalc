package main

import (
	"fmt"
	"log"
	"strings"
)

type args struct {
	class   string
	ctype   string
	credits float64
	grade   float64
}

func (a args) calcGPA() (totalCreds, totalPoints float64) {
	var subtractor, gpa float64

	switch a.ctype {
	case strings.ToLower("cp"):
		subtractor = 1.0
	case strings.ToLower("h"):
		subtractor = 0.5
	case strings.ToLower("ap"):
		subtractor = 0.0
	default:
		log.Panic("err on line 31")
	}

	switch {
	case a.grade >= 98 && a.grade <= 100:
		gpa = 5.0 - subtractor
	case a.grade >= 93 && a.grade <= 100:
		gpa = 4.7 - subtractor
	case a.grade >= 90 && a.grade <= 92:
		gpa = 4.5 - subtractor
	case a.grade >= 87 && a.grade <= 89:
		gpa = 4.3 - subtractor
	case a.grade >= 83 && a.grade <= 86:
		gpa = 4.0 - subtractor
	case a.grade >= 80 && a.grade <= 82:
		gpa = 3.7 - subtractor
	case a.grade >= 77 && a.grade <= 79:
		gpa = 3.4 - subtractor
	case a.grade >= 73 && a.grade <= 76:
		gpa = 3.0 - subtractor
	case a.grade >= 70 && a.grade <= 72:
		gpa = 2.7 - subtractor
	case a.grade >= 65 && a.grade <= 69:
		gpa = 2.0 - subtractor
	case a.grade >= 0 && a.grade <= 64:
		gpa = 0
	}
	totalCreds += a.credits
	points := gpa * a.credits
	totalPoints += points

	return totalCreds, totalPoints
}

func prettyPrint(classes, grade, credits, classType []string, points []float64) {
	fmt.Printf("%-10s %-7s %-10s %-8s %-10s\n", "Class", "CP/H", "Credits", "Grade", "GPA")
	fmt.Println(strings.Repeat("-", 45))

	for i := 0; i < 6; i++ {
		class := strings.Title(classes[i])
		cType := strings.ToUpper(classType[i])

		//fmt.Printf("%-9s %-7s %-4s %s\n", class, cType, credits[i], grade[i])
		fmt.Printf("%-10s %-7s %-10s %-8s %-10g \n", class, cType, credits[i], grade[i], points[i])
	}
}
