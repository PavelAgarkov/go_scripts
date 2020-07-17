package slice

import (
	"fmt"
	"sort"
	"unicode"
)

const DELETED = ""

func StartReverse() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j/2; i++ {
		(*s)[i], (*s)[j-i] = (*s)[j-i], (*s)[i]
	}
}

func StartRotate() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v\n", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Printf("%v\n", s)
}

func DeleteAdjacentRepeatElement() {
	str := []string{"one", "one", "one", "one", "two", "two", "two", "two", "two", "five", "three", "three", "go"}
	deleteDublication(&str)
	fmt.Print(str)
}

func deleteDublication(str *[]string) {
	for i, v := range *str {

		if !findIndex(i+1, str) || len(*str) <= i {
			break
		} else {

			if v == (*str)[i+1] || v == (*str)[i-1] {
				(*str)[i] = DELETED
			} else {
				continue
			}

		}
	}

	sort.Slice(*str, func(i, j int) bool {
		return (*str)[i] < (*str)[j]
	})

	counter := 0
	for _, v := range *str {
		if v == DELETED {
			counter++
		}
	}

	*str = (*str)[counter:]
}

func findIndex(index int, str *[]string) bool {
	for i, _ := range *str {
		if index == i {
			return true
		}
	}
	return false
}

func UnicodeToASCII() []rune {

	strings := []rune{' ', ' ', ' ', ' ', 0, 123, ' ', ' ', 1, 2, 3, ' ', ' '}

	var link [][]int

	lastCount := 0

	for i, v := range strings {

		if unicode.IsSpace(v) {

			if lastCount > 0 && v == strings[i] && len(strings)-1 == i {
				link = append(link, []int{i - lastCount, i})
				lastCount = 0
				break
			}

			if v == strings[i+1] {
				lastCount++
			} else {
				link = append(link, []int{i - lastCount, i})
				lastCount = 0
			}
		}

	}

	function := func(index int) bool {
		for _, externalValue := range link {
			if index > externalValue[0] && index <= externalValue[1] {
				return true
			}
		}
		return false
	}

	indexCounter := 0
	for ind, _ := range strings {
		if function(ind) {
			strings = append(strings[:ind-indexCounter], strings[ind+1-indexCounter:]...)
			indexCounter++
		}
	}

	return strings
}
