// timers

package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers #########
	// if you want to do something at some point in future, use "timer" fo that
	// gives us a channel which will be notified about the time
	// when this timer will finish/will be fired
	timer := time.NewTimer(time.Second * 3)

	// awaiting for timer firing
	<-timer.C
	fmt.Println("timer fired")

	// Tickers ########
	// when you want to do something repeatedly in regular intervals

	// exaample of a ticker that ticks in regular interval

	ticker := time.NewTicker(time.Second * 4)
	done := make(chan bool)

	go func() {
		for{
			select{
				case 
			}
		}
	}()

}
