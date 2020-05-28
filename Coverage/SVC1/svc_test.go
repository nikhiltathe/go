package SVC1

import "testing"

func TestSVC1(t *testing.T) {
	a := WhoAMI()
	if a != "SVC1" {
		t.Error("Failed")
	}
}
