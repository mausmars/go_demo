package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"fmt"
)

const (
	//DSN = "root:tongzhen@tcp(127.0.0.1:3306)"
	DSN = "root:@tcp(10.0.106.205:3306)"
	DB_SETTLE = "test"
)

//扣费结算
func Settle(userID int, planID int, charge int)  int{
	db, err := sql.Open("mysql", DSN + "/" + DB_SETTLE)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handle, err := db.Prepare("CALL test.transfer(?, ?, ?, @outStatus)")
	if err != nil {
		panic(err.Error())
	}
	defer handle.Close()

	//call procedure
	var result sql.Result
	result, err = handle.Exec(userID, planID, charge)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)

	var sql = "SELECT @outStatus as ret_status"
	selectInstance, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer selectInstance.Close()

	var ret_status int
	err = selectInstance.QueryRow().Scan(&ret_status)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ret_status)
	return ret_status
}

func main()  {
	result:=Settle(1,2,200)
	fmt.Println(result)
}