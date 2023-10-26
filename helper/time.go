package helper

import "time"

func GetCurrentTimestamp() string {
	currentDate := time.Now()
	return currentDate.Format("2006-01-02 15:04:05")
}
