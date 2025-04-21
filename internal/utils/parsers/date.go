package parsers

import (
	"time"

	"github.com/gabrielteiga/startup-rush/configs"
)

// Parsing a string date to the time.Time type using the layout configured in the config.go file
func StringDateToTime(date string) (time.Time, error) {
	return time.Parse(configs.DATE_LAYOUT, date)
}
