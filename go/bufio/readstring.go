package bufio

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func bufio() {
	words := []string{}
	data := "A foggy mountain\nAn old falcon.\nA wise man."
	sc := bufio.NewScanner(strings.NewReader(data))
	sc.Split(bufio.ScanWords)

	n := 0
	for sc.Scan() {
		words = append(words, sc.Text())
		n++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# of words: %d\n", n)

	for _, word := range words {
		fmt.Println(word)
	}
}
