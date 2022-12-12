package notify

import "testing"

func TestSendSMS(t *testing.T) {

	r := SendSMS("+44-123", "changes committed")
	if r != true {
		t.Errorf("SendSMS(...) = %t; want true", r)
	}
}
