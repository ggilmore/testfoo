package main // import "whoami.com/foo"

import "whoami.com/foo/bar"

func main() {
	b := bar.Baz{
		Contents: []string{
			"Ham",
			"Eggs",
		},
	}

	b.PrintContents()
}