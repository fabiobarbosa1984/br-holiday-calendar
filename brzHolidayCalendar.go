package brzHolidayCalendar

import (
	"strconv"
	"strings"
	"time"
)

// method for checking if a specific date is a holiday
func IsHoliday(date time.Time) bool {
	return holidays[DateToString(date)]
}

// method for simple conversion from time to brazilian date format
func DateToString(d time.Time) string {
	return d.Format("02/01/2006")
}

// function to convert a string date in brazilian tipical format (DD/MM/YYYY) to a time.Time var
func StringToDate(ds string) (time.Time, error) {
	if ds != "" {
		parts := strings.Split(ds, "/")
		day, err := strconv.Atoi(parts[0])
		if err != nil {
			return time.Time{}, err
		}
		month, err := strconv.Atoi(parts[1])
		if err != nil {
			return time.Time{}, err
		}
		year, err := strconv.Atoi(parts[2])
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, nil
}

// Method for calculating working days between to dates, excluding weekends and holidays
func NetWorkDays(id, fd time.Time) (int, error) {
	wd := 0
	finalDate := fd.AddDate(0, 0, -1)

	for date := id; !date.After(finalDate); date = date.AddDate(0, 0, 1) {
		if !IsHoliday(date) && date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			wd++
		}
	}
	return wd, nil
}

// method to identify the first work day from the date, including itself
func nextWorkDay(d string) (string, error) {
	date, err := StringToDate(d)
	if err != nil {
		return "", err
	}

	fd := date.AddDate(0, 0, 6)

	for ; !date.After(fd); date = date.AddDate(0, 0, 1) {
		if !IsHoliday(date) && date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			return DateToString(date), nil
		}
	}
	return "", nil
}
