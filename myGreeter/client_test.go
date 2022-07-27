package myGreeter

import (
	"fmt"
	"testing"
	"time"
)

const (
	morning   = "Good morning"
	afternoon = "Good afternoon"
	evening   = "Good evening"
)

func validClientGreeter(t *testing.T, timeStr, key string) {
	fmt.Println("初始：", timeStr)
	timer, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	msg, err := Client{}.GreeterForTime(timer.Hour())
	if err != nil {
		t.Fatal("invalid time")
	}
	if msg != key {
		t.Fatal("testing failed")
	}
}

func TestClientMorning(t *testing.T) {
	timeStr := "2022-07-27 06:15:05"
	validClientGreeter(t, timeStr, morning)
}

func TestClientAfternoon(t *testing.T) {
	timeStr := "2022-07-27 16:15:05"
	validClientGreeter(t, timeStr, afternoon)
}

func TestClientEvening(t *testing.T) {
	timeStr := "2022-07-27 23:15:05"
	validClientGreeter(t, timeStr, evening)
}
