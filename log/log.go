package log

import (
	"os"
	"fmt"
	"log"
)

func InitLog(logFile string)*log.Logger{
	//"C:/Users/Administrator/Desktop/data.log"
	logfile,err:=os.OpenFile(logFile,os.O_RDWR|os.O_APPEND,0666)
	if err!=nil{
		fmt.Printf("%s\r\n",err.Error())
		os.Exit(-1)
	}
	logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
	return logger
}