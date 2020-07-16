package directory

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"s3dump/util/context"
	"s3dump/util/db"
	"s3dump/util/json"
)

//Reads neccesary configuration data from config file
func DBConnect() {
	file, err := ioutil.ReadFile("config/conffile.json")
	if err != nil {
		log.Fatalf("Configuraiton '%s' file not found", "config/conffile.json")
		return
	}
	data := json.Parse(file)
	context.Instance().Set("host", data.GetString("host"))
	context.Instance().Set("port", data.GetString("port"))
	context.Instance().Set("user", data.GetString("user"))
	context.Instance().Set("password", data.GetString("password"))
	context.Instance().Set("database", data.GetString("database"))
	context.Instance().Set("bucket", data.GetString("bucket"))
	context.Instance().Set("key", data.GetString("key"))
	context.Instance().Set("region", data.GetString("region"))
	err = db.Connect()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

}

func CopyFolders() error {

	file, err := ioutil.ReadFile("config/folder.json")
	if err != nil {
		log.Fatalf("Configuraiton '%s' file not found", "config/folder.json")
		return err
	}
	data := json.Parse(file)
	//fmt.Println(data)

	folder := data.GetStringArray("folder")
	//fmt.Println(folder)
	for _, i := range folder {
		cmd := exec.Command("cp", "-r", i, "backup")
		err = cmd.Run()
		if err != nil {
			return err
		}
		//time.Sleep(5 * time.Second)
	}
	//time.Sleep(5 * time.Second)
	return nil
}

func RemoveContents() error {
	dir := "backup"
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	os.Remove("sample.sql")
	os.Remove("compress.sql")
	return nil
}

func Compress() error {
	cmd := exec.Command("tar", "-zcvf", "compress.zip", "backup", "sample.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
