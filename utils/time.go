package utils

import "time"

func IsValidUTC(str string) bool {
	layout := time.RFC3339 // UTC format
	_, err := time.Parse(layout, str)
	return err == nil
}
