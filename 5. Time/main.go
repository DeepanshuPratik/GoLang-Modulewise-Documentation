package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My time is going great")
	currTime := time.Now()
	// We have to mention exact date and day to get correct output.
	// MM DD YYYY
	fmt.Println("Current Time : ", currTime.Format("01-02-2006 Monday 15:04:05"))
	createDate := time.Date(2025, time.August, 18, 13, 22, 05, 0, time.UTC)
	// self created date
	// DD MM YYYY
	fmt.Println("Date of completion of B.Tech degree : ", createDate.Format("02-01-2006 Monday 15:04:05"))
}

// command to build for different os is GOOS="windows/linux/darwin" go build
