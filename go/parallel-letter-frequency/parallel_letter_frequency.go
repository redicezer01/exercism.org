package letter

import "sync"
import _ "fmt"

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkSequentialFrequency-8             59113             18987 ns/op
// BenchmarkConcurrentFrequency-8             90918             13243 ns/op

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var m = FreqMap{}
var letterMap = map[rune]bool{}

type letterFreq struct {
	l   rune
	cnt int
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	lf := make(chan letterFreq, 100)

	var wg sync.WaitGroup
	for _, s := range l {
		for _, r := range s {
			if !letterMap[r] {
				wg.Add(1)
				// TODO: limit the count of goroutines
				// run  in parallel (concurrent) to count letter frequency.
				// one goroutine for one uniqe letter
				go countLetterFreq(l, r, &wg, lf)
				letterMap[r] = true
			}
		}
	}
	go func() {
		wg.Wait()
		close(lf)
	}()

	// gather from each goroutine and collect letters frequency in one map.
	for item := range lf {
		m[item.l] = item.cnt
	}
	return m
}

// countLetterFreq counts specified letter frequency
func countLetterFreq(l []string, rtw rune, wg *sync.WaitGroup, lf chan letterFreq) {
	defer wg.Done()
	lm := letterFreq{l: rtw, cnt: 0}
	for _, s := range l {
		for _, r := range s {
			if r == rtw {
				lm.cnt++
			}
		}
	}
	lf <- lm
}
