package preprocess

import "testing"

func TestVirtualPhone(t *testing.T) {
	if !IsVirtualPhoneNumber("16211451419") {
		t.Fatalf("it should be virtual phone number")
	}
}
