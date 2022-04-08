package main

import (
	"flag"
	"fmt"
	"korflab/entropy"
	"korflab/fastafile"
	"os"
)

var print = fmt.Println

func main() {

	// CLI
	arg_i := flag.String("i", "", "fasta file (default stdin)")
	arg_w := flag.Int("w", 11, "window size")
	arg_t := flag.Float64("t", 1.5, "entropy threshold")
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	w := *arg_w
	t := *arg_t

	reader := fastafile.NewReader(*arg_i)
	for {
		if reader.DoneReading {
			break
		}
		entry := fastafile.NextEntry(reader)

		// mask - not centered over window
		mask := make([]rune, len(entry.Seq))
		for i := 0; i < len(entry.Seq)-w+1; i++ {
			h := entropy.SeqEntropy(entry.Seq[i : i+w])
			if h < t {
				mask[i] = 'N'
			} else {
				mask[i] = rune(entry.Seq[i])
			}
		}

		// last unmasked window
		for i := len(entry.Seq) - w; i < len(entry.Seq); i++ {
			mask[i] = rune(entry.Seq[i])
		}

		// output as fasta - not wrapping
		print(entry.Def)
		print(string(mask))
	}

}
