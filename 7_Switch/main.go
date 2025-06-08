package main

import (
	"fmt"
	"time"
)

// go does not require a break statement in switch cases as it automatically breaks after each case

func simpleSwitch() {
	var day int
	fmt.Println("Enter a number (1-7) for the day of the week:")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input, please enter a number between 1 and 7.")
	}
}

func multiConditionSwitch() {
	var score int
	fmt.Println("Enter your score:")
	fmt.Scanln(&score)

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}

func dateTimeSwitch() {
	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Unknown day")
	}
}

func typeSwitch() {
	whoAmi := func(i interface{}) {
		switch i := i.(type) {
		case int:
			fmt.Println("I am an integer:", i)
		case string:
			fmt.Println("I am a string:", i)
		case bool:
			fmt.Println("I am a boolean:", i)
		default:
			fmt.Println("I am of unknown type")
		}
	}

	whoAmi(42)
	whoAmi("Hello")
	whoAmi(true)
	whoAmi(3.14)
}

func main() {
	simpleSwitch()
	multiConditionSwitch()
	dateTimeSwitch()
	typeSwitch()
}
