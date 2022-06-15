package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year < 1918 {
		return year%4 == 0
	} else {
		return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	}

}

func dayOfTheProgrammer(year int) int {
	// The year of the transition from Julian to gregorian was 1918 and Julian Calendar had gained quiete a few years
	// the problem statements says that the 14th of February was the 32nd day of the year in gregorian calendar
	// so normally when the year is not leap from january to august we have 243 so:
	// 256th day of the year is the day of the programmer
	// 243 is number of days from january 1 fst to august 31, 1918 is not a leap year
	// 13 is the number of days the calendar have won so far
	// and when we do 256 - (243 - 13) we get the date of the programmers day in russia in 1918 which is 26
	if year == 1918 {
		return 26
	}
	// 256th of the year is always on september
	// these are the number of days up until august
	nDays := 243

	if isLeapYear(year) {
		return 256 - (nDays + 1)
	}

	return 256 - nDays

}

func main() {

	var year int
	fmt.Scanf("%d\n", &year)

	fmt.Printf("%d.09.%d\n", dayOfTheProgrammer(year), year)
}
