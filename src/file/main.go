package main

import (
	"fmt"
	"os"
)

func main()  {

	file,err := os.Create("test/test.txt")

	fmt.Println(err)
	fmt.Println(file)
}
