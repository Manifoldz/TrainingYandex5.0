package main

import (
	"fmt"
	"time"
)

func main() {
	var holidayN, year int
	_, err := fmt.Scanf("%d\n%d\n", &holidayN, &year)
	if err != nil {
		fmt.Print("Read-mistake holidayN number")
		return
	}
	if holidayN < 0 || holidayN > 366 || year < 1800 || year > 2100 {
		fmt.Print("Wrong input value")
		return
	}

	daysOfWeek := map[string]int{
		"Monday":    52 + holidayN,
		"Tuesday":   52 + holidayN,
		"Wednesday": 52 + holidayN,
		"Thursday":  52 + holidayN,
		"Friday":    52 + holidayN,
		"Saturday":  52 + holidayN,
		"Sunday":    52 + holidayN,
	}

	for i := 0; i < holidayN; i++ {
		var day int
		var month string
		var date_layout = "2-January-2006"

		fmt.Scanf("%d %s", &day, &month)
		date, err := time.Parse(date_layout, fmt.Sprintf("%d-%s-%d", day, month, year))
		if err != nil {
			fmt.Print(err)
			return
		}
		daysOfWeek[date.Weekday().String()] -= 1

	}
	var firstWeekDay string
	fmt.Scan(&firstWeekDay)

	// Проверка на наличие ключа в словаре
	_, ok := daysOfWeek[firstWeekDay]
	if !ok {
		fmt.Print(err)
		return
	}
	daysOfWeek[firstWeekDay] += 1

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {

		nextDay := map[string]string{
			"Monday":    "Tuesday",
			"Tuesday":   "Wednesday",
			"Wednesday": "Thursday",
			"Thursday":  "Friday",
			"Friday":    "Saturday",
			"Saturday":  "Sunday",
			"Sunday":    "Monday",
		}

		// Проверка на наличие ключа в словаре
		_, ok := nextDay[firstWeekDay]
		if !ok {
			fmt.Print(err)
			return
		}
		daysOfWeek[nextDay[firstWeekDay]] += 1
	}

	max_key := "Monday"
	max := daysOfWeek[max_key]
	min := max
	min_key := max_key

	for str, val := range daysOfWeek {
		if val > max {
			max = val
			max_key = str
		} else if val < min {
			min = val
			min_key = str
		}
	}

	fmt.Printf("%s %s", max_key, min_key)
}
