package helper

import (
	"fmt"
	"time"

	"github.com/Alihanc/web-develoment/model"
)

func Maintance_3600() {

	start_time, err := time.Parse("20-01-2023", model.InitialData.MAINTANCE_DATE)
	if err != nil {
		panic(err)
	}
	finish_time := start_time.Add(3600 * time.Hour)
	fmt.Println("Tahmini 3600 saat  bakım tarihi:", finish_time)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(finish_time)

		if start_time == finish_time {
			fmt.Println("Bakım zamanı geldi.")
			break
		} else {

			//fmt.Printf("Bakıma kalan süre: %02d\n", timeRemaining)
			fmt.Printf("Days: %02d Hours: %02d Minutes: %02d Seconds: %02d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)

		}

	}
}
func Maintance_7200() {

	start_time, err := time.Parse("20-01-2023", model.InitialData.MAINTANCE_DATE)
	if err != nil {
		panic(err)
	}
	finish_time := start_time.Add(7200 * time.Hour)
	fmt.Println("Tahmini 7200 saat  bakım tarihi:", finish_time)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(finish_time)

		if start_time == finish_time {
			fmt.Println("Bakım zamanı geldi.")
			break
		} else {

			//fmt.Printf("Bakıma kalan süre: %02d\n", timeRemaining)
			fmt.Printf("Days: %02d Hours: %02d Minutes: %02d Seconds: %02d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)

		}

	}
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
