package observer

import (
	"btcapp/src/settings"
	"fmt"
	"net/smtp"
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

func SendGmailLetterStrategy(
	userGmail string,
	BTCPrice float64,
	errors chan error,
	settings settings.Settings,
) {
	message := fmt.Sprintf("To: %s\r\nSubject: BTC/UAH price\r\n\r\nBTC/UAH price: %f\r\n", userGmail, BTCPrice)

	errors <- smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			settings.Gmail,
			settings.GmailPassword,
			"smtp.gmail.com",
		),
		settings.Gmail,
		[]string{userGmail},
		[]byte(message),
	)
}
