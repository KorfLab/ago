package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"math"
	"strings"
)

func dust(w int, h float64, n bool, seqs []string, ids []string){
	for i := 0; i < len(ids); i++ {
		seq := seqs[i]
		for j := 0; j < len(seq) - w + 1; j++ {
			nt := make(map[int]int)
			nt[65] = 0
			nt[67] = 0
			nt[71] = 0
			nt[84] = 0
			for k := 0; k < w; k++ {
				char := int(seq[j+k])
				nt[char] += 1
			}
			a := float64(nt[65]/w)
			c := float64(nt[67]/w)
			g := float64(nt[71]/w)
			t := float64(nt[84]/w)
			
			entropy := 0.0
			if nt[65] > 0 {entropy -= a * math.Log(a)}
			if nt[67] > 0 {entropy -= c * math.Log(c)}
			if nt[71] > 0 {entropy -= g * math.Log(g)}
			if nt[84] > 0 {entropy -= t * math.Log(t)}
			
			entropy /= math.Log(2)
			
			if entropy < h {
				if n == false {
					seq = seq[:j+w/2] + "N" + seq[j+w/2+1:]
				}
				if n == true {
					seq = seq[:j+w/2] + strings.ToLower(string(seq[j+w/2])) + seq[j+w/2+1:]
				}
			
			}
		}
		fmt.Print(">",ids[i],"\n")
		fmt.Println(seq)
	
	}

}
func main() {
	fs := flag.String("f", "", "path to file")
    w := flag.Int("w", 11, "window size")
    h := flag.Float64("h", 1.1, "enthropy threshold")
    n := flag.Bool("n", false, "N based masking")
    
    flag.Parse()
    
    ff, err := os.Open(*fs)
	if err != nil {
		panic(err)
	}
    
    seqs := []string{}
    ids := []string{}
    seq := ""
    
    fasta := bufio.NewScanner(ff)
	for fasta.Scan() {
		line := fasta.Text()
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
	ff.close()
	
	
    dust(*w, *h, *n, seqs, ids)
}