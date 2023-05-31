package observer

import (
	"fmt"
	"time"
)

func ConsoleLogStrategy(userGmail string, BTCPrice float64, errors chan error) {
	time.Sleep(2 * time.Second)
	if userGmail == "" {
		errors <- fmt.Errorf("empty user")
		return
	}
	fmt.Printf("%s got gmail with price: %f\n", userGmail, BTCPrice)
}

func SendGmailLetterStrategy(userGmail string, BTCPrice float64, errors chan error) {
	panic("not implemented")
}
