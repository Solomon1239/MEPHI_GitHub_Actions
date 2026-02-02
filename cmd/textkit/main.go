package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/pipeline-go-demo/pkg/textkit"
)

func main() {
	mode := flag.String("mode", "slug", "mode: slug or dist")
	a := flag.String("a", "", "first string")
	b := flag.String("b", "", "second string")
	flag.Parse()

	switch *mode {
	case "slug":
		if *a == "" {
			fmt.Fprintln(os.Stderr, "provide -a")
			os.Exit(2)
		}
		fmt.Println(textkit.Slugify(*a))
	case "dist":
		if *a == "" || *b == "" {
			fmt.Fprintln(os.Stderr, "provide -a and -b")
			os.Exit(2)
		}
		fmt.Println(textkit.LevenshteinDistance(*a, *b))
	default:
		fmt.Fprintln(os.Stderr, "unknown -mode; use slug or dist")
		os.Exit(2)
	}
}
