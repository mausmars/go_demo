package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"go_demo/src/profile"
)

func readbyte2(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	defer profile.Start().Stop()
	defer profile.Start(profile.MemProfile).Stop()

	//filePath := "doc/file_test/moby.txt"
	filePath := "/mnt/hgfs/vm_share/go_demo/doc/file_test/moby.txt"

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file %q: %v", filePath, err)
	}

	b := bufio.NewReader(f)
	words := 0
	inword := false
	for {
		r, err := readbyte2(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", filePath, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", filePath, words)
}
