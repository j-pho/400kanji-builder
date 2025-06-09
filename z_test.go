package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	IN  = "./testdata.txt"
	OUT = "./testdata-out.txt"
)

func TestSkipper(t *testing.T) {
	fi, err := os.Open(IN)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	var txt string
	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		txt += sc.Text() + "\n"
	}

	counter := 0
	lines := strings.Split(txt, "\n")
	var trimmed []string
	for _, line := range lines {
		if counter < 4 {
			trimmed = append(trimmed, line)
		}
		counter++
		if counter == 10 {
			counter = 0
		}
	}

	fo, err := os.Create(OUT)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	buffer := bufio.NewWriter(fo)
	for _, line := range trimmed {
		x := line + "\n"
		_, e := buffer.WriteString(x)
		if e != nil {
			log.Fatal(e)
		}
	}
}
