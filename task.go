package main

import (
	"fmt"
	"s3dump/directory"
	"s3dump/process"
	"s3dump/util/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	directory.DBConnect()
	dbase := process.CreateDB()
	if !dbase {
		return
	}
	table := process.CreateTable()
	if !table {
		return
	}
	for j := 1; j < 10; j++ {
		err := process.InsertData()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err := directory.RemoveContents()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Dump()
	if err != nil {
		fmt.Println(err, "761517365")
		return
	}

	err = directory.CopyFolders()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = directory.Compress()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = process.Upload()
	if err != nil {
		fmt.Println(err)
		return
	}
}
