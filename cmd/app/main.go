package main

import (
	"html/template"
	"log"
	"os"

	"github.com/opusdv/weather/internal/app/weather"

	"github.com/opusdv/weather/internal/app/myformat"
)

func main() {
	info, err := weather.GetWeatherInfo("Kazan", "metric")
	if err != nil {
		log.Fatal(err)
	}

	report, err := template.New("report").Funcs(template.FuncMap{
		"unixTime": myformat.UnixTime}).Parse(myformat.Templ)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdin, info); err != nil {
		log.Fatal(err)
	}

}
