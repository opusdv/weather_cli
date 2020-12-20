package myformat

import (
	"time"
)

// UnixTime ...
func UnixTime(t int64) time.Time {
	ut := time.Unix(t, 0)
	return ut
}

// Templ ...
const Templ = `Щирота: {{.Coord.Lat}}
Долгоmа: {{.Coord.Lon}}
{{range .Weather}}-----------------------------
Main: {{.Main}}
Description: {{.Description}}
{{end}}
Температура: {{.Main.Temp}}
Минимальная температура: {{.Main.TempMin}}
Максимальная Температура: {{.Main.TempMax}}
-----------------------------------------------
Время рассвета: {{.Sys.Sunrise | unixTime}}
Время заката: {{.Sys.Sunset | unixTime}}
Время Dt: {{.Dt | unixTime}}
`
