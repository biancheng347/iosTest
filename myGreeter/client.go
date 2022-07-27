package myGreeter

import "errors"

type Client struct {
}

func (c Client) GreeterForTime(hour int) (msg string, err error) {
	if hour < 0 || hour > 24 {
		err = errors.New("invalid hour")
		return
	}

	switch {
	case hour >= 6 && hour < 12:
		msg = "Good morning"
	case hour >= 12 && hour < 18:
		msg = "Good afternoon"
	default:
		msg = "Good evening"
	}
	return
}
