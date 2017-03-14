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

	_, err := db.Start("drop table Student")
	if err == nil {
	fmt.Println("Table deleted")	
	} else if e, ok := err.(*mysql.Error); ok {
		// Error from MySQL server
		fmt.Println(e)
	} else {
		checkError(err)
	}

	fmt.Println("Create A table... ")
	checkedResult(db.Query("CREATE TABLE Student (rollno int, name varchar(40), marks int)"))
	
	fmt.Println("Insert into A... ")

	checkedResult(db.Query("INSERT Student VALUES(1,'ABC',50)"))
	checkedResult(db.Query("INSERT Student VALUES(2,'PQR',55)"))
	checkedResult(db.Query("INSERT Student VALUES(3,'XYZ',5)"))
	checkedResult(db.Query("INSERT Student VALUES(4,'JKL',45)"))
	checkedResult(db.Query("INSERT Student VALUES(5,'LMN',25)"))
	checkedResult(db.Query("INSERT Student VALUES(6,'QWR',56)"))


	rows, res := checkedResult(db.Query("SELECT * from Student"))
	rollno := res.Map("rollno")
	name := res.Map("name")
	marks := res.Map("marks")

	fmt.Printf("\n------------------------")
	fmt.Printf("\nRollno: Name    Marks   ")
	fmt.Printf("\n------------------------\n")

	for _, row := range rows {
		fmt.Println(row.Int(rollno),"\t"+row.Str(name)+"\t",row.Int(marks))
	}

/*	for ii, row := range rows {
		fmt.Printf(
			"Row: %d Roll no : %d \n name:  %-10s \n number: %-8d  \n", ii,
			row.Int(rollno), 
			"'"+row.Str(name)+"'",
			row.Int(marks), 
		)
	}
*/


	var query = "SELECT * from Student WHERE marks <= 10"
	rows,_,_ = db.Query(query)

	fmt.Printf("\n-------Failed------------")
	fmt.Printf("\nRollno: Name    Marks    ")
	fmt.Printf("\n-------------------------\n")
	
	for _, row := range rows {
		fmt.Println(row.Int(rollno),"\t"+row.Str(name)+"\t",row.Int(marks))
	}


	fmt.Println("\nClose connection... ")
	checkError(db.Close())
	fmt.Println("Disconnected ")		
}