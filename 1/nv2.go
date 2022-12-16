package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addNumbers(a string, b string) string {
	l1 := len(a)
	l2 := len(b)

	i := l1
	j := l2
	var sum []string
	carry := 0

	for {
		if i > 0 && j > 0 {
			num1, err := strconv.Atoi(string(a[i-1]))
			if err != nil {
				num1 = 0
			}

			num2, err := strconv.Atoi(string(b[j-1]))
			if err != nil {
				num2 = 0
			}

			// ASCII value of 0 :
			currSum := num1 + num2 + carry
			if currSum > 10 {
				carry = 1
			} else {
				carry = 0
			}
			lastdigit := currSum % 10
			sum = append(sum, strconv.Itoa(lastdigit))
			fmt.Println(lastdigit)
			i--
			j--
		} else if i > 0 && j <= 0 {
			fmt.Println("to be added :", string(a[0:i]))
			for i > 0 {
				sum = append(sum, string(a[i-1]))
				i--
			}
			break
		} else if i <= 0 && j > 0 {
			fmt.Println("to be added :", string(b[0:j]))
			for j > 0 {
				sum = append(sum, string(b[j-1]))
				j--
			}
			break
		} else if i <= 0 && j <= 0 {
			break
		}
		// } else if i>0 && j <= 0 {

		//  //  1323   5 ->
		//  num1 := a[0:i-1]
		//  currSum := num1+carry
		//  i= 0

		//  } else if i<=0 && j > 0 {
		//  num2 := b[0:i-1]
		//  currSum := num2+carry
		//     j=0
		// }
	}

	//
	fmt.Println(sum)

	return strings.Join(sum, "")

}

func main() {
	// var num1, num2, res string
	// fmt.Scanf("%v\n%v", &num1, &num2)
	res := addNumbers("3", "145")
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
	fmt.Println(res)
}

// a : 123456
// b : 456789
//=================
// ->  580245
