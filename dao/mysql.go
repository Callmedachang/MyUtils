package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
	"strings"
)
type MySqler struct {
	User string
	PassWord string
	Host string
	Port string
	DbName string
	CharSet string
}
func (mySqler *MySqler)getConnection()*sql.DB{
	db, err := sql.Open("mysql", mySqler.User+":"+mySqler.PassWord+"@tcp("+mySqler.Host+":"+mySqler.Port+")/"+mySqler.DbName+"?charset="+mySqler.CharSet)
	checkErr(err)
	return db
}
func checkErr(i error) {
	if i!=nil {
		fmt.Println(i)
	}
}
func dealStruct(bean interface{})([]string,[]interface{},string){
	var fieldNames =make([]string,2)
	var values =make([]interface{},2)
	object := reflect.ValueOf(bean)
	myref := object.Elem()
	typeOfType := myref.Type()
	structValues :=strings.Split(typeOfType.Name(),".")
	var tableName string
	if len(structValues)>1 {
		tableName = structValues[len(structValues)-1]
	}else{
		tableName = structValues[0]
	}
	tableName = strings.ToLower(tableName)
	for i:=0; i<myref.NumField(); i++{
		field := myref.Field(i)
		fieldNames[i] = typeOfType.Field(i).Name
		values[i] = field.Interface()
	}
	return fieldNames,values,tableName
}
func composeFields(fieldNames []string)string{
	var fields string
	for i:=0;i<len(fieldNames);i++  {
		if i!=len(fieldNames)-1 {
			fields = fields+strings.ToLower(fieldNames[i])+","
		}else{
			fields = fields+strings.ToLower(fieldNames[i])
		}
	}
	return fields
}
func composeValues(values []interface{})string{
	var pars string
	for i:=0;i<len(values);i++  {
		if i!=len(values)-1 {
			pars = pars+"?"+","
		}else{
			pars = pars+"?"
		}
	}
	return pars
}
func (mySqler *MySqler)Insert(bean interface{})int64{
	db:=mySqler.getConnection()
	fieldNames,values,tableName :=dealStruct(bean)
	fields := composeFields(fieldNames)
	pars := composeValues(values)
	sql :="INSERT into "+tableName+" ("+fields+") values ("+pars+")"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err :=stmt.Exec(values[0],values[1])
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}
func (mySqler *MySqler)Query(tableName string,condition string)*sql.Rows{
	db:=mySqler.getConnection()
	sql :="SELECT * from "+tableName+" where "+condition
	fmt.Println(sql)
	stmt, err := db.Prepare(sql)
	checkErr(err)
	resule ,err:=stmt.Query()
	return resule
}
func (mySqler *MySqler)Update(tableName string,change string,condition string){
	db:=mySqler.getConnection()
	//UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值
	sql :="UPDATE "+tableName+" SET "+change+" where "+condition
	stmt, err := db.Prepare(sql)
	checkErr(err)
	stmt.Exec()
}
func (mySqler *MySqler)Delete(bean interface{}){
	db:=mySqler.getConnection()
	var condition string
	fieldNames,values,tableName :=dealStruct(bean)
	//DELETE FROM 表名称 WHERE 列名称 = 值
	for i:=1;i<len(fieldNames) ;i++  {
		if i!=len(fieldNames)-1 {
			condition =condition+fieldNames[i]+" = ? AND"
		}else{
			condition =condition+fieldNames[i]+" = ?"
		}
	}
	sql :="DELETE FROM "+tableName+" WHERE "+condition
	fmt.Println(sql)
	stmt, err := db.Prepare(sql)
	checkErr(err)
	stmt.Exec(values[1])
}