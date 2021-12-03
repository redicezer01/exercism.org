// package twelve returns sentences of song 'The Twelve days of Christmas'
package twelve

import "strings"
import "fmt"
import "bytes"

const giftsStr = ` 
a Partridge in a Pear Tree
two Turtle Doves
three French Hens
four Calling Birds
five Gold Rings
six Geese-a-Laying
seven Swans-a-Swimming
eight Maids-a-Milking
nine Ladies Dancing
ten Lords-a-Leaping
eleven Pipers Piping
twelve Drummers Drumming
`
const daysStr = `
first
second
third
fourth
fifth
sixth
seventh
eighth
ninth
tenth
eleventh
twelfth
`

// Verse returns one line for specified days (i int)
func Verse(i int) string {
	giftsS := strings.Split(giftsStr, "\n")[1:]
	daysS := strings.Split(daysStr, "\n")[1:]
	var res bytes.Buffer
	tail := ""
	format := "On the %s day of Christmas my true love gave to me: %s."
	// constructing recursive tail of the sentence
	for day := 0; day < i; day++ {
		and := ""
		if day == 1 {
			and = "and "
		}
		if day == 0 {
			tail = giftsS[day]
		} else {
			tail = fmt.Sprintf("%v, %v%v", giftsS[day], and, tail)
		}

	}
	// composing a sentence
	fmt.Fprintf(&res, format, daysS[i-1], tail)
	return res.String()
}

// Song returns whole text of the song
func Song() string {
	var song bytes.Buffer
	// get all lines and building result text
	for i := 1; i <= 12; i++ {
		newLine := "\n"
		if i == 12 {
			newLine = ""
		}
		fmt.Fprintf(&song, "%s%s", Verse(i), newLine)
	}
	return song.String()
}
