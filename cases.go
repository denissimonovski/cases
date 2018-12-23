package main

import (
	"fmt"
	"time"
)

var Day time.Duration = 24 * time.Hour
var WorkDay time.Duration = 8 * time.Hour

func checkWorkTime(d time.Time) time.Time {
	if d.Weekday() == time.Weekday(6) {
		afterTwoDays := d.Add(Day * 2)
		return time.Date(afterTwoDays.Year(), afterTwoDays.Month(), afterTwoDays.Day(), 8, 0, 0, 0, time.UTC)
	} else if d.Weekday() == time.Weekday(0) {
		afterOneDay := d.Add(time.Hour * 24)
		return time.Date(afterOneDay.Year(), afterOneDay.Month(), afterOneDay.Day(), 8, 0, 0, 0, time.UTC)
	} else if d.Hour() >= 16 {
		afterOneDay := d.Add(Day)
		return time.Date(afterOneDay.Year(), afterOneDay.Month(), afterOneDay.Day(), 8, 0, 0, 0, time.UTC)
	} else if d.Hour() < 8 {
		return time.Date(d.Year(), d.Month(), d.Day(), 8, 0, 0, 0, time.UTC)
	}
	return d
}

func odzemi(s, e time.Time) time.Duration {
	startTime := checkWorkTime(s)
	endTime := checkWorkTime(e)
	if startTime.Month() == endTime.Month() && startTime.Day() == endTime.Day() {
		return e.Sub(s)
	}
	return razlicenDen(s, e) + (time.Duration(denovi(s, e)) * WorkDay)
}

func denovi(s, e time.Time) int {
	denovi := 0
	den1 := time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.UTC)
	den2 := time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, time.UTC)
	for i := 0; i < 345; i++ {
		if den1.Add(Day) == den2 {
			break
		} else {
			if den1.Add(Day).Weekday() != time.Weekday(0) && den1.Add(Day).Weekday() != time.Weekday(6) {
				denovi += 1
				den1 = den1.Add(Day)
			} else {
				den1 = den1.Add(Day)
			}
		}
	}
	return denovi
}

func razlicenDen(s, e time.Time) time.Duration {
	vreme1 := time.Date(s.Year(), s.Month(), s.Day(), 16, 0, 0, 0, time.UTC).Sub(s)
	vreme2 := e.Sub(time.Date(e.Year(), e.Month(), e.Day(), 8, 0, 0, 0, time.UTC))
	return vreme1 + vreme2

}
func main() {
	startDate := time.Date(2018, 11, 27, 13, 15, 0, 0, time.UTC)
	endDate := time.Date(2018, 11, 28, 9, 15, 0, 0, time.UTC)
	vreme := odzemi(startDate, endDate)
	fmt.Println(vreme)
	startDate = time.Date(2018, 11, 27, 13, 15, 0, 0, time.UTC)
	endDate = time.Date(2018, 11, 29, 9, 15, 0, 0, time.UTC)
	vreme = odzemi(startDate, endDate)
	fmt.Println(vreme)
	startDate = time.Date(2018, 11, 27, 13, 15, 0, 0, time.UTC)
	endDate = time.Date(2018, 11, 30, 10, 15, 0, 0, time.UTC)
	vreme = odzemi(startDate, endDate)
	fmt.Println(vreme)
	startDate = time.Date(2018, 11, 27, 13, 15, 0, 0, time.UTC)
	endDate = time.Date(2018, 12, 4, 12, 15, 0, 0, time.UTC)
	vreme = odzemi(startDate, endDate)
	fmt.Println(vreme)

}
