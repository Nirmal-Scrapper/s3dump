package db

import (
	"database/sql"
	"fmt"
	"s3dump/util/context"
	"strconv"

	"github.com/JamesStewy/go-mysqldump"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Connect to DB
func Connect() error {
	host := context.Instance().Get("host")
	portstring := context.Instance().Get("port")
	user := context.Instance().Get("user")
	password := context.Instance().Get("password")
	database := context.Instance().Get("database")
	port, _ := strconv.Atoi(portstring)
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
	fmt.Println(mysqlInfo)
	dbs, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return err
	}
	db = dbs
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("connected to database")
	return nil
}

//List multiple rows
func List(sqlStatement string) *sql.Rows {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return rows
}

//Read single row
func Read(sqlStatement string) *sql.Row {
	row := db.QueryRow(sqlStatement)
	fmt.Println(row)
	return row
}

//Exec insert,update and delete...
func Exec(sqlStatement string) (sql.Result, error) {
	//fmt.Println(sqlStatement)
	result, err := db.Exec(sqlStatement)
	if err != nil {
		return result, err
	}
	//fmt.Println(result)
	return result, nil
}

func Dump() error {
	host := context.Instance().Get("host")
	portstring := context.Instance().Get("port")
	user := context.Instance().Get("user")
	password := context.Instance().Get("password")
	database := "sampledb"
	port, _ := strconv.Atoi(portstring)
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
	fmt.Println(mysqlInfo)
	dbs, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dumper, err := mysqldump.Register(dbs, "./", "sample")
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return err
	}

	// Dump database to file
	a, err := dumper.Dump()
	if err != nil {
		fmt.Println(err, a)
		return err
	}
	//fmt.Println(err, a)
	return nil
}
