package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(2000, 01, 23, 10, 12, 45, 233223, time.Local)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	fmt.Println(time.Now().Format("2006-01-02"))
}
