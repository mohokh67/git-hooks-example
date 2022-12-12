package notify

import "fmt"

func SendSMS(recipient, msg string) bool {
	fmt.Printf("sending text message to '%s' with message: '%s'\n", recipient, msg)
	return true
}
