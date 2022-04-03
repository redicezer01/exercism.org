package encode

import "strings"
import "strconv"
import _ "fmt"

func RunLengthEncode(input string) string {
	res := strings.Builder{}
	length := len(input)
	if length == 0 || length == 1 {
		return input 
	}
	var (
		curr byte
		cnt int
		cnt_str string
	)
	
	for i := 0; i < length; i++ {
		curr = input[i]
		cnt++
		if cnt > 1 {
			cnt_str = strconv.Itoa(cnt)
		}
		if i == length-1 || curr != input[i+1] {
			res.WriteString(cnt_str+string(curr))
			cnt = 0
			cnt_str = ""
		}
	}

	return res.String()
}

func RunLengthDecode(input string) string {
	length = len(input)
	for i := 0; i < length; i++ {
			
	}
}
