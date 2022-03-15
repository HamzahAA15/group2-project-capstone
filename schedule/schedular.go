package schedule

import (
	"sync"
	"time"
)

var ticker *time.Ticker = nil

func checkoutsTime() {
	t := time.Now()
	duration := t.AddDate(0, 0, 1).Sub(t)
	if ticker == nil {
		ticker = time.NewTicker(time.Duration(duration.Hours()) * time.Hour)
	}

	for {
		CheckoutAttendance()
		<-ticker.C
	}
}

func generateCalender() {
	t := time.Now()
	duration := t.AddDate(0, 1, 0).Sub(t)
	if ticker == nil {
		ticker = time.NewTicker(time.Duration(duration.Hours()) * time.Hour)
	}

	for {
		GenerateCalender()
		<-ticker.C
	}
}

func Schedular() {
	wg.Add(2)
	go time.AfterFunc(timeCheckouts(), checkoutsTime)
	go time.AfterFunc(timeGenerateCalender(), generateCalender)
	wg.Wait()
}

func timeCheckouts() time.Duration {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 00, 0, t.Location())
	if t.After(n) {
		n = n.AddDate(0, 0, 1)
	}

	d := n.Sub(t)

	return d
}

func timeGenerateCalender() time.Duration {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), 1, 00, 00, 00, 0, t.Location())
	if t.After(n) {
		n = n.AddDate(0, 1, 0)
	}

	d := n.Sub(t)

	return d
}

var wg sync.WaitGroup
