package main

import "fmt"
import "flag"
import "os"
import "bufio"
import "strings"
import "math"

func main() {
	//Argument Parser
	ff := flag.String("fasta", " ", "path to fasta file")
	w := flag.Int("w", 11, "window size [1]")
	h := flag.Float64("h", 1.1, "entropy threshold [1.1]")
	lc := flag.Bool("lc", false, "mask with lower case")
	flag.Parse()
	
	//Read FASTA
	fh, err := os.Open(*ff)
	if err != nil {
		panic(err)
	}
	
	seqs := []string{}
	ids := []string{}
	seq := ""
	
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if len(ids) >= 1 {
				seqs = append(seqs,seq)
				seq = ""
			}
			ids = append(ids,line[1:])
		 } else {
		 	seq += line
		 }
	}
	seqs = append(seqs,seq)
	
	//Main
	for i := 0; i < len(ids); i++ {
		seq = seqs[i]
		for j := 0; j < len(seq) - *w + 1; j++ {
			a := 0
			c := 0
			g := 0
			t := 0
			for k:= 0; k < *w; k++ {
				if seq[j+k] == 'A' {a++}
				if seq[j+k] == 'C' {c++}
				if seq[j+k] == 'G' {g++}
				if seq[j+k] == 'T' {t++}
			}
			//Get entropy
			pa := float64(a) / float64(*w)
			pc := float64(c) / float64(*w)
			pg := float64(g) / float64(*w)
			pt := float64(t) / float64(*w)
			
			ent := 0.0
			if pa > 0 {ent -= pa * math.Log(pa)}
			if pc > 0 {ent -= pc * math.Log(pc)}
			if pg > 0 {ent -= pg * math.Log(pg)}
			if pt > 0 {ent -= pt * math.Log(pt)}
			ent = ent / math.Log(2)
			
			//Mask
			if ent < *h {
				if *lc == false {
					seq = seq[:j+*w/2] + "N" + seq[j+*w/2+1:]
				}
				if *lc == true {
					seq = seq[:j+*w/2] + strings.ToLower(string(seq[j+*w/2])) + seq[j+*w/2+1:]
				}
			}
		}
		fmt.Print(">",ids[i],"\n")
		fmt.Println(seq)
	}
	
	
	fh.Close()
	
}
