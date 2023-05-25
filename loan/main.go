package main

func main() {
	age := 25
	isEmployed := true
	activeEmployementTime := 2
	salary := 100000
	if age < 22 {
		println("Not available for your age only 22+")
	} else if isEmployed == false {
		println("Only available to people who are employed")
	} else if activeEmployementTime <= 1 {
		println("Only available to those who are employed for 1 or more years")
	} else if salary >= 100000 {
		println("Your loan is not gonna have interests")
	} else {
		println("Loan granted")
	}

}
