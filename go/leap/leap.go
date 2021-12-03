// package leap provides functional to work with leap year
package leap

// IsLeapYear determine leap year
func IsLeapYear(year int) bool {
	divBy400 := year%400 == 0
	divBy100 := year%100 == 0
	divBy4 := year%4 == 0
	return divBy4 && !(divBy100 && !divBy400)
}
