package main

import (
	"fmt"
	"github.com/werbhelius/pilot/model"
)

func main() {
	fmt.Println("hello, world!")

	var a model.UnitTemp = 3.2
	fmt.Println(a.FormatTemp())
}
