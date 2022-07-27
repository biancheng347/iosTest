package main

import (
	"iOSTestCode/myGreeter"
	"time"
)

func main() {
	myGreeter.Client{}.GreeterForTime(time.Now().Hour())
}
