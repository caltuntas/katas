package bufio

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func bufio() {
	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
