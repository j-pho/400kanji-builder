package main

import (
	"bufio"
	"log"
	"os"
)

// $ go build
// $ ./400kanji-builder

func main() {
	const (
		F1  = "merge1.txt"
		F2  = "merge2.txt"
		F3  = "merge3.txt"
		F4  = "merge4.txt"
		F5  = "merge5.txt"
		OUT = "merged-kanji-cards.txt"
	)

	files := []string{F1, F2, F3, F4, F5}
	ss := make([][]string, len(files))

	for i := range files {
		a, err := os.Open(files[i])
		if err != nil {
			panic(err)
		}
		defer a.Close()
		sc := bufio.NewScanner(a)
		for sc.Scan() {
			ss[i] = append(ss[i], sc.Text())
		}
	}

	// create file
	f, err := os.Create(OUT)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	buffer := bufio.NewWriter(f)
	for i := 0; i < len(ss[0]); i++ {
		for j := 0; j < len(ss); j++ {
			if i < len(ss[j]) {
				x := ss[j][i] + "\n"
				_, e := buffer.WriteString(x)
				if e != nil {
					log.Fatal(e)
				}
			}
		}
	}
}
