package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	form := "Jan 02 2006 15:04"

	input := "Wed Dec 31 18:37"

	fmt.Println("input : ", input)

	// Split current time and add year to make valid time objec
	parts := strings.Split(input, " ")
	//fmt.Println("parts len :", len(parts))
	//fmt.Printf("parts %+v\n", parts)

	input = parts[1] + " " + parts[2] + " 2021 " + parts[3]
	fmt.Println("new input : ", input)

	t2, e := time.Parse(form, input)
	//fmt.Println("t2 :", t2)
	//fmt.Println("t2 UnixNano:", t2.UnixNano())
	//fmt.Println("t2 Unix:", t2.Unix())

	// Learning : If no year mentioned it take 0000

	//	fmt.Println("t2 String:", t2.String())
	if e != nil {
		fmt.Println("e : ", e)
	}
	//	fmt.Println("t2 form : ", t2.Format(form))

	//	nowTime := time.Now()
	location := time.UTC
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("err : ", err)
	}
	nowTime := time.Date(2021, 01, 01, 9, 15, 3, 0, location)
	//fmt.Println("nowTime:", nowTime)
	//fmt.Println("nowTime UnixNano:", nowTime.UnixNano())
	//fmt.Println("nowTime Unix:", nowTime.Unix())

	fmt.Println("current formated time: ", nowTime.Format(form))

	if t2.After(nowTime) {
		fmt.Println("input is after nowTime")
	} else {
		fmt.Println("nowTime is after t2")
	}

	requiredTime := "Dec 23 2020 20:20"
	t3, e := time.Parse(form, requiredTime)
	fmt.Println("t3 :", t3)

	if t3.After(t2) {
		fmt.Println("required is after input")
	} else {
		fmt.Println("input is after required")
	}

}
