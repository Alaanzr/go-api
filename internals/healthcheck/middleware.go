package middleware

import (
	"fmt"
	"time"
)

func LogRequestStart(f func()) {
	fmt.Println(time.Now())
	f()
}
