package main

import (
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
