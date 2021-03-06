package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	n := "test.jpg"
	fmt.Println(strings.Index(n, "png"))
	fmt.Println(strings.IndexAny(n, "e"))

	fmt.Println(n[strings.LastIndex(n, ".")+1 : len(n)])

	now := time.Now()
	y := now.Year()
	m := int(now.Month())
	d := now.Day()
	fmt.Println(y, m, d)

	path := "/Users/irving/itest"
	fmt.Println(os.MkdirAll(path, 0777))
}
