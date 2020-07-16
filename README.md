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

creates db
creates table
insert dummy value
creates dump file
copies the content in json(config/folder.json) into backup dir
compress backup folder and dump file into a zip file
uploads to aws s3 bucket


run the task.go file
