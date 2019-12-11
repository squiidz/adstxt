package pkg

import (
	"testing"
)

func TestKey(t *testing.T) {
	s := &Seller{
		PublisherName: "test.com",
		Domain:        "seller.com",
		AccountID:     "123",
		TypeOfAccount: "DIRECT",
	}
	if string(s.key()) != "ads:test.com:seller:123" {
		t.Fail()
	}
}
