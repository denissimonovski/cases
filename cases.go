package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

const Day time.Duration = 24 * time.Hour
const WorkDay time.Duration = 8 * time.Hour
const endOfWork int = 16
const startOfWork int = 8

func checkWorkTime(d time.Time) time.Time {
	if d.Weekday() == time.Weekday(6) {
		afterTwoDays := d.Add(Day * 2)
		return time.Date(afterTwoDays.Year(), afterTwoDays.Month(), afterTwoDays.Day(), startOfWork, 0, 0, 0, time.UTC)
	} else if d.Weekday() == time.Weekday(0) {
		afterOneDay := d.Add(Day)
		return time.Date(afterOneDay.Year(), afterOneDay.Month(), afterOneDay.Day(), startOfWork, 0, 0, 0, time.UTC)
	} else if d.Hour() >= endOfWork {
		afterOneDay := d.Add(Day)
		if afterOneDay.Weekday() == time.Weekday(6) {
			afterThreeDays := d.Add(Day * 3)
			return time.Date(afterThreeDays.Year(), afterThreeDays.Month(), afterThreeDays.Day(), startOfWork, 0, 0, 0, time.UTC)
		}
		return time.Date(afterOneDay.Year(), afterOneDay.Month(), afterOneDay.Day(), startOfWork, 0, 0, 0, time.UTC)
	} else if d.Hour() < startOfWork {
		return time.Date(d.Year(), d.Month(), d.Day(), startOfWork, 0, 0, 0, time.UTC)
	}
	return d
}

func odzemi(s, e time.Time) time.Duration {
	startTime := checkWorkTime(s)
	endTime := checkWorkTime(e)
	if startTime.Month() == endTime.Month() && startTime.Day() == endTime.Day() {
		return endTime.Sub(startTime)
	}
	return razlicenDen(startTime, endTime) + (time.Duration(denovi(startTime, endTime)) * WorkDay)
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
	var VkupnoVreme time.Duration
	xlFileName := "/root/go/src/cases/dookt.xlsx"
	xlFile, err := xlsx.OpenFile(xlFileName)
	if err != nil {
		fmt.Println(err)
	}
	AdventSheet := xlFile.Sheet["AdventNetReport"]
	for _, row := range AdventSheet.Rows {
		startTime := parseTime(row.Cells[2].String())
		var endTime time.Time
		if row.Cells[3].String() == "-" {
			//endTime = time.Date(2018, 31, 12, 16, 0, 0, 0, time.UTC)
			continue
		} else {
			endTime = parseTime(row.Cells[3].String())
		}
		vremeCase := odzemi(startTime, endTime)
		VkupnoVreme += vremeCase
		fmt.Println(row.Cells[1:4], vremeCase.String())
	}
	fmt.Println(VkupnoVreme/time.Duration(len(AdventSheet.Rows)), len(AdventSheet.Rows))
}

func parseTime(s string) time.Time {
	vreme, err := time.Parse("02-01-2006 15:04", s)
	if err != nil {
		fmt.Println(err)
	}
	return vreme
}
