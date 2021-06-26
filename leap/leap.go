// Package leap provides methods to perform calculations related to leap years.
package leap

// IsLeapYear identifies if the given year is a leap year or not. Returns a bool
// as the result.
func IsLeapYear(year int) bool {
	/* A leap year should follow the following rules:
	 1. Should be divisible by 4, i-e: year % 4 == 0
	 2. Should either be:
		a. not divisible by 100, i-e: year % 100 != 0 or;
		b. divisible by 400 i-e: year % 400 == 0
	*/
	return (year%4 == 0) && (year%100 != 0 || year%400 == 0)
}
