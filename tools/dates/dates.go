package dates

import "time"

func Generate() string {
	return time.Now().Format(time.RFC3339)
}
