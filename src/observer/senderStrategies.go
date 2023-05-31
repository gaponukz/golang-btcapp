package observer

import (
	"btcapp/src/settings"
	"fmt"
	"net/smtp"
)

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
