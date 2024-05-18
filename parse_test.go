package librmonitor

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_Parse(t *testing.T) {
	f, err := os.Open("./data/sebring.txt")
	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(f)

	start := time.Now()
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i++
		Parse(line)
	}
	end := time.Now()

	fmt.Println("Took " + end.Sub(start).String())
	fmt.Printf("Processed %d records\n", i)
}
