package main

import (
	"buffio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}

		if txt < prev {
			panic("file not sorted")
		}

		prev = txt
		alreadySeen[txt] = true
		fmt.Println(txt)
	}

}
