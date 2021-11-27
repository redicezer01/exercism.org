// package luhn implements Luhn algorithm
package luhn

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkValid-8        54861319                21.89 ns/op

// TODO: find better way
// Valid check a string using Luhn algorithm
func Valid(id string) bool {
	id_len := len(id)

	if id_len <= 1 {
		return false
	}

	sum := 0
	fi := 0
	digit := 0
	var sym uint8

	// loop over string and processing
	for i := id_len - 1; i >= 0; i-- {
		sym = id[i]

		// check symbol if not allowed return false (digit or space)
		if (sym < '0' || '9' < sym) && sym != ' ' {
			return false
		}

		// skip spaces
		if sym == ' ' {
			continue
		}

		// convert to int
		digit = int(sym - '0')

		// multiply by 2 every second digit from right
		// nasted if  is worse
		if (fi+1)%2 == 0 {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}

		sum += digit
		fi++
	}

	// valid number ?
	if sum%10 != 0 || fi == 1 {
		return false
	}
	return true
}
