package heyjay_test

import (
	"os"
	"testing"

	"github.com/whaangbuu/heyjay"
)

type User struct {
	Fullname string
	Age      int
	Address  string
}

func TestHeyJay(t *testing.T) {
	user := User{Fullname: "Richard Burk", Age: 22, Address: "Cebu City"}
	err := heyjay.JSON(os.Stdout, user)

	if err != nil {
		t.Errorf("Error due to: %v", err)
	}
}
