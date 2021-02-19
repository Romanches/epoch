package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var u int
	flag.IntVar(&u, "u", 1613732386, "unixtime")
	flag.Parse()
	unixtime := flag.Arg(0)

	if unixtime == "" {
		fmt.Println("Usage: epoch [unixtime]")
		fmt.Println("Now: ", time.Now().Unix(), time.Now())
		flag.PrintDefaults()
		os.Exit(1)
	}

	i, _ := strconv.Atoi(unixtime)
	fmt.Printf("%v", time.Unix(int64(i), 0))
}
