package main

import (
	"fmt"
	"time"

	"utils/timeutil"
)

func main() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		loc = time.FixedZone("KST", 9*60*60)
	}

	// Parse a local time string with an explicit location (avoid time.Local).
	t, err := timeutil.ParseTimeIn("2026-04-26 19:30:33", loc)
	if err != nil {
		panic(err)
	}

	midnight, err := timeutil.MidnightIn(t, loc)
	if err != nil {
		panic(err)
	}

	truncatedHour, err := timeutil.TruncateToHourIn(t, loc)
	if err != nil {
		panic(err)
	}

	fmt.Println("original:", t.Format(timeutil.DefaultTimeLayout), t.Format(timeutil.LogTimestampLayout))
	fmt.Println("midnight:", midnight.Format(timeutil.DefaultTimeLayout), midnight.Format(timeutil.LogTimestampLayout))
	fmt.Println("hour:", truncatedHour.Format(timeutil.DefaultTimeLayout), truncatedHour.Format(timeutil.LogTimestampLayout))

	start, _ := timeutil.ParseDateIn("2026-04-01", loc)
	end, _ := timeutil.ParseDateIn("2026-04-30", loc)
	fmt.Println("valid_range:", timeutil.IsValidDateRange(start, end))

	fmt.Println("utc_instant:", timeutil.ToUTC(t).Format(time.RFC3339Nano))
}
