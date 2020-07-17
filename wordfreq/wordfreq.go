package wordfreq

import (
	"bufio"
	"fmt"
	"os"
)

var text = make(map[string]int)
var all []string

func TxtRead() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	i := 0
	for scanner.Scan() {
		i++
		word := scanner.Text()
		all = append(all, word)

		if i%10 == 0 {
			equals(all[i-10 : i])
		}
	}

	fmt.Printf("Слово: \tвключений: \n")
	for i, v := range text {
		fmt.Printf("%v\t%v\n", i, v)
	}

}

func equals(dictionary []string) {
	for _, v := range dictionary {

		if !findWordInValuesArr(v) {
			text[v] = 1
		} else {
			text[v]++
		}
	}
}

func findWordInValuesArr(word string) bool {
	for i, _ := range text {
		if i == word {
			return true
		}
	}
	return false
}
