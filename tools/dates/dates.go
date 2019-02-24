package dates

import (
	"strings"
	"time"
)

func Generate() string {
	return time.Now().Format(time.RFC3339)
}

func PickDate(datetime string) string {
	return strings.Split(datetime, "T")[0]
}
