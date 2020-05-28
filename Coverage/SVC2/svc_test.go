package SVC2

import "testing"

func TestSVC2(t *testing.T) {
	a := WhoAMI()
	if a != "SVC2" {
		t.Error("Failed")
	}
}
