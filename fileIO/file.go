package fileIO

import (
	"io"
	"os"
	"io/ioutil"
	"log"
)

func WriteToNewFile(data []byte,filePath string)(error,error){
	f, err := os.Create(filePath)
	_,err2 :=io.WriteString(f,string(data))
	return err,err2
}
func AppendFile(data []byte,filePath string){
	fd,_:=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd.Write(data)
	fd.Close()
}
func ReadFile(filePath string) []byte{
	myFile, err := os.Open(filePath)
	if err != nil {
	}
	defer myFile.Close()
	data,_ := ioutil.ReadAll(myFile)
	return data
}
func DeleteFile(filePath string){
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("success delete")
	}
}
