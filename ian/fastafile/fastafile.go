package fastafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Entry struct {
	Def string
	Seq string
}

func NewEntry(def string, seq string) *Entry {
	p := Entry{Def: def, Seq: seq}
	return &p
}

type Reader struct {
	filename    string
	lastline    string
	scanner     *bufio.Scanner
	DoneReading bool
}

func NewReader(filename string) *Reader {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastline string
	if scanner.Scan() {
		lastline = scanner.Text()
	} else {
		fmt.Println("file is empty")
		os.Exit(1)
	}

	p := Reader{filename: filename, lastline: lastline,
		scanner: scanner, DoneReading: false}
	return &p
}

func NextEntry(reader *Reader) *Entry {
	if reader.DoneReading {
		fmt.Println("attempt to read from fully parsed file")
		os.Exit(1)
	}

	def := reader.lastline
	var seqs []string

	complete := true
	for reader.scanner.Scan() {
		line := reader.scanner.Text()
		if line[0] == '>' {
			reader.lastline = line
			complete = false
			break
		} else {
			seqs = append(seqs, line)
		}
	}

	if complete {
		reader.DoneReading = true
	}

	return NewEntry(def, strings.Join(seqs, ""))
}
