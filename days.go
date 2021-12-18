package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	day = 24 * time.Hour

	datefmt    = "2006-01-02"
	altDatefmt = "20060102"
)

func main() {
	days := flag.Int("d", 0, "add `d` days to today")
	date := flag.String("date", "", "return number of days since/to date")
	flag.Parse()

	if *date != "" {
		for _, f := range []string{datefmt, altDatefmt} {
			if dd, err := time.Parse(f, *date); err == nil {
				dd = time.Date(dd.Year(), dd.Month(), dd.Day(), 0, 0, 0, 0, time.Local)
				fmt.Println(int(time.Until(dd) / 24 / time.Hour))
				break
			}
		}
	} else {
		today := time.Now()

		if os.Getenv("DDATE") != "" {
			for _, f := range []string{datefmt, altDatefmt} {
				if dd, err := time.Parse(f, os.Getenv("DDATE")); err == nil {
					today = dd
					break
				}
			}
		}

		day := today.AddDate(0, 0, *days)
		fmt.Println(day.Format(datefmt))
	}
}
