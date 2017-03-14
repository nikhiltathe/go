package main

import (
	"os"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkedResult(rows []mysql.Row, res mysql.Result, err error) ([]mysql.Row,
	mysql.Result) {
	checkError(err)
	return rows, res
}

func main() {
	user := "testuser"
	pass := "TestPasswd9"
	dbname := "test"
	proto := "tcp"
	addr := "127.0.0.1:3306"

	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Println("Connect to %s:%s... ", proto, addr)
	checkError(db.Connect())
	fmt.Println("Connected ")	

	_, err := db.Start("drop table A")
	if err == nil {
	fmt.Println("Table deleted")	
	} else if e, ok := err.(*mysql.Error); ok {
		// Error from MySQL server
		fmt.Println(e)
	} else {
		checkError(err)
	}

	fmt.Println("Create A table... ")
	checkedResult(db.Query("CREATE TABLE A (name varchar(40), number int)"))
	
	fmt.Println("Insert into A... ")
	for i := 1; i <= 10; i++ {
		checkedResult(db.Query("INSERT A VALUES(%d,%d)", i,i*10))
	}

	rows, res := checkedResult(db.Query("SELECT * from A"))
	name := res.Map("name")
	number := res.Map("number")
/*	for ii, row := range rows {
		fmt.Printf(
			"Row: %d\n name:  %-10s \n number: %-8d  \n", ii,
			"'"+row.Str(name)+"'",
			row.Int(number), 
		)
	}
*/

	//Select and print value of 5 and 7
	//var myname string = "5"
	var query = "SELECT * from A WHERE name <= 5"

	rows,_,_ = db.Query(query)
//	fmt.Println(rows)

	//var number int
	for _, row := range rows {
		fmt.Println("\n Value of",row.Str(name)," is",row.Int(number))
	}


	fmt.Println("Close connection... ")
	checkError(db.Close())
	fmt.Println("Disconnected ")		
}