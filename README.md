# s3dump

can change the database configuration and aws configuration in config/conffile.json
{
    "host":"localhost",        //db host
    "user":"root",             //db user
    "password":"",             //db password
    "port":"3306",             //db port
    "database":"",             //db name
    "bucket":"filesqldumpzip",   //s3 bucket name
    "key":"compress.zip",       
    "region":"us-east-1"         //aws region
}

uses default aws iam user set in system

run the task.go file
