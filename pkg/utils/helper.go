package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

// Return local timezone
func LocalTime() time.Time{
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

// Convert string time to defind layout
func ConvertStringTimeToTimelayout(t string) time.Time{
	layout := "2015-03-02 12:00:00.999 -0700 MST"
	result , err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parst string time failed: %s", err.Error())
	}
	return result
}