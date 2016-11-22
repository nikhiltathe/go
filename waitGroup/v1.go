package main

import (
	"log"
)

type Result struct {
	NumAlerts  int
	NumChassis int
	NumFans    int
}

func GetNumAlerts() (int, error) {
	return 10, nil
}

func GetNumChassis() (int, error) {
	return 1, nil
}

func GetNumFans() (int, error) {
	return 2, nil
}

func main() {
	var r Result

	na, err := GetNumAlerts()
	if err != nil {
		log.Fatal(err)
	}

	nc, err := GetNumChassis()
	if err != nil {
		log.Fatal(err)
	}

	nf, err := GetNumFans()
	if err != nil {
		log.Fatal(err)
	}

	r = Result{na, nc, nf}

	log.Println(r)
}
