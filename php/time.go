package php

import (
	"strings"
	"time"
)

// Time - Return current Unix timestamp
func Time() int64 {

	return time.Now().Unix()
}


// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06",   // A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01",      // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1",       // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan",     // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2",  // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon",    // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3",  // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Date - Format a local time/date
func Date(format string, ts ...time.Time) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

// DateAdd -  Adds an amount of days, months, years, hours, minutes and seconds to a DateTime object
func DateAdd(t time.Time, years int, months int, days int) time.Time {
	return t.AddDate(years, months, days)
}

// Sleep - Delay execution
func Sleep(s int64) {
	time.Sleep(time.Duration(s) * time.Second)
}

// Usleep - Delay execution in microseconds
func Usleep(ms int64) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}