package main

import (
	"flag"
	"github.com/werbhelius/pilot/ui"
)

func main() {

	// global flags
	city := flag.String(
		"city",
		"Beijing",
		"city must be request (you can find city name、id and coord in https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json) ")
	flag.StringVar(city, "c",
		"Beijing",
		"city must be request (you can find city name、id and coord in https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json) ")
	lang := flag.String("lang", "zh_ch", "weather desc language ")
	flag.StringVar(lang, "l", "zh_ch", "weather desc language ")
	flag.Parse()

	ui.Render(*city, *lang)

}
