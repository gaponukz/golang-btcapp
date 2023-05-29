package observer

import "fmt"

func ConsoleLogStrategy(userGmail string, BTCPrice float64) {
	fmt.Printf("%s got gmail with price: %f", userGmail, BTCPrice)
}

func SendGmailLetterStrategy(userGmail string, BTCPrice float64) {
	panic("not implemented")
}
