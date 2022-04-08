package main

import "fmt"
import "flag"



func main() {
	var win = flag.Int("w", 12, "window size")
	flag.Parse()
	fmt.Println("hello world")
	fmt.Println(*win)
	fmt.Println(flag.Args())
}
