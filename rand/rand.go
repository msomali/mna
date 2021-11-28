package rand

import (
	"github.com/techcraftlabs/mna"
	"math/rand"
	"time"
)

const (
	numbersCharSet = "0123456789"
	suffixLength = 7
)

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func prefix() string {
	prefixes := []string{
		"071", "065", "067","074", "075", "076",
		"073", "078", "068", "069", "077", "066",
		"072", "061", "062", "063", "064",

	}
	randomPos := seededRand.Intn(len(prefixes))
    return prefixes[randomPos]
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func suffix() string {
	return stringWithCharset(suffixLength, numbersCharSet)
}

func Generate() string {
    return prefix() + suffix()
}

func GenerateN(len int) []string {
	// return string of generated random numbers using Generate()
	var numbers []string
	for i := 0; i < len; i++ {
        numbers = append(numbers, Generate())
    }
	return numbers
}

// GenerateNWithFilters generates a list of random numbers of length len,
// the number have to pass the filters to be added to the list
// if the number doesn't pass the filters, it will be skipped
func GenerateNWithFilters(n int, f1 mna.FilterPhoneFunc, f2 mna.FilterOperatorFunc) []string {
    var numbers []string
	i := 0
	for i < n {
		number := Generate()
		op,_ := mna.Get(number)
		if f1!=nil && f2!=nil {
			if f1(number) && f2(op) {
				numbers = append(numbers, number)
				i++
			}
		}else if f1!=nil {
            if f1(number) {
                numbers = append(numbers, number)
                i++
            }
		}else {
			if f2(op) {
                numbers = append(numbers, number)
                i++
            }
		}

	}
    return numbers
}