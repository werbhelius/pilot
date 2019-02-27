package main

import (
	"flag"
	"fmt"
	"github.com/werbhelius/pilot/api"
)

func main() {

	// global flags
	city := flag.String(
		"city",
		"Beijing",
		"city must be request (you can find city name、id and coord in https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json) ")

	lang := flag.String("lang", "zh_ch", "weather desc language ")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*lang)

	api.Request(*city, *lang)

}
