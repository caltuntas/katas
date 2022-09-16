package bufio

import (
	"fmt"
	"log"
	"os"
  "bufio"
)

func bufio() {
	data := []string{"an old falcon", "misty mountains",
		"a wise man", "a rainy morning"}

	f, err := os.Create("words.txt")

  if err !=nil {
    log.Fatal(err)
  }

  defer f.Close()

  wr := bufio.NewWriter(f)

  for _, line := range data {
    wr.WriteString(line + "\n")
  }

  wr.Flush()

  fmt.Println("data written")
}
