package brzHolidayCalendar

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	//test a randon christmas date. Expect to return true
	td, _ := StringToDate("25/12/2023")
	if IsHoliday(td) == false {
		t.Errorf("2023 christmas not identified as a holiday")
	}

	//test a randon work date. Expect to return false
	td, _ = StringToDate("11/09/2023")
	if IsHoliday(td) == true {
		t.Errorf("%v wrongly identified as a holiday", td)
	}
}

func TestStringToDate(t *testing.T) {
	//ensures the conversion is using brazilian format
	strDate := "12/01/2001"
	d, _ := StringToDate(strDate)
	if d.Month() == time.December {
		t.Errorf("Exepct January, found December. Original string date: %v. Converted: %v", strDate, d)
	}
}

func TestDateToString(t *testing.T) {
	//ensures the conversion is using brazilian format
	d := time.Date(2001, time.December, 1, 0, 0, 0, 0, time.UTC)
	if DateToString(d) == "12/01/2001" {
		t.Errorf("Error converting date to brazilian format. Expect 01/12/2001, found 12/01/2001")
	}
}

func TestNetWorkDays(t *testing.T) {
	//ensures proper work days calculation
	id := time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC)
	fd := time.Date(2070, time.April, 1, 0, 0, 0, 0, time.UTC)
	r, _ := NetWorkDays(id, fd)
	if r != 11889 {
		t.Errorf("Error in net work days calc. Expected 332. Found %v", r)
	}
}

func TestNextWorkDAy(t *testing.T) {
	//ensures proper next day identification
	wd, _ := StringToDate("11/09/2023")
	twd, _ := nextWorkDay(wd)
	satd, _ := StringToDate("09/09/2023")
	tsatd, _ := nextWorkDay(satd)
	sund, _ := StringToDate("10/09/2023")
	tsund, _ := nextWorkDay(sund)
	hd, _ := StringToDate("07/09/2023")
	thd, _ := nextWorkDay(hd)

	//common workday test
	if twd != "11/09/2023" {
		t.Errorf("Error estimating the next work day for a commomn work day. Expected %v, found %v", wd, twd)
	}

	//saturday test
	if tsatd != "11/09/2023" {
		t.Errorf("Error estimating the next work day for a saturday. Expected 11/09/2023, found %v", tsatd)
	}

	//sunday test
	if tsund != "11/09/2023" {
		t.Errorf("Error estimating the next work day for a saturday. Expected 11/09/2023, found %v", tsund)
	}

	//holiday test
	if thd != "08/09/2023" {
		t.Errorf("Error estimating the next work day for a saturday. Expected 08/09/2023, found %v", thd)
	}
}

func TestHolidayMapSize(t *testing.T) {
	//checks for unwanted holiday list change
	l := len(holidays)
	if l != 1187 {
		t.Errorf("Holiday list size error. Exepected: 1187 Founded:%v", l)
	}
}
