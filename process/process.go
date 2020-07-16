package process

import (
	"fmt"
	"s3dump/util/aws"
	"s3dump/util/context"
	"s3dump/util/db"
	"strings"
)

func CreateDB() bool {
	_, err := db.Exec("CREATE DATABASE sampledb")
	if err != nil {
		if strings.Contains(err.Error(), "database exists") {
			fmt.Println("database already exists")
		} else {
			return false
		}
	}
	_, err = db.Exec("use sampledb")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CreateTable() bool {
	_, err := db.Exec("CREATE table firsttable(name varchar(30),city varchar(30),phone varchar(20),age int);")
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("table already exists")
		} else {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func InsertData() error {
	_, err := db.Exec("insert into firsttable(name,city,phone,age) values ('ross geller','new york','+1944774874674',32);")
	if err != nil {
		return err
	}
	return nil
}

func Upload() error {
	err := aws.Upload("compress.zip", context.Instance().Get("bucket"), context.Instance().Get("key"), context.Instance().Get("region"))
	return err
}
