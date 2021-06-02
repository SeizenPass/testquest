package fast

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isLetter(byteVal byte) bool {
	if (byteVal >= 65 && byteVal <= 90) || (byteVal >= 97 && byteVal <= 122) {
		return true
	}
	return false
}

func readBytes(output chan []byte, input []byte) {
	var slice []byte
	for _, v := range input {
		if isLetter(v) {
			if v >= 65 && v <= 90 {
				v += 32
			}
			slice = append(slice, v)
		} else if len(slice) != 0 {
			output <- slice
			slice = []byte{}
		}
	}
	close(output)
}

type word struct {
	data    []byte
	counter int
}

func acceptBytes(input chan []byte, out io.Writer) {
	var words []*word
	for data := range input {
		if words == nil {
			words = append(words, &word{
				data:    data,
				counter: 1,
			})
		} else {
			found := false
			for i, v := range words {
				if bytes.Equal(v.data, data) {
					words[i].counter++
					found = true
					break
				}
			}
			if !found {
				words = append(words, &word{
					data:    data,
					counter: 1,
				})
			}
		}
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].counter > words[j].counter
	})
	for i := 0; i < 20; i++ {
		fmt.Fprintf(out, "%v %v\n", words[i].counter, string(words[i].data))
	}
}

func OurSolution(out io.Writer) {
	f, err := ioutil.ReadFile("mobydick.txt")
	check(err)
	output := make(chan []byte, 100)
	go readBytes(output, f)
	acceptBytes(output, out)
}
