package database_set


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "lyk123"
)

var DB *sql.DB

func INIT_DB(){
	path :=strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	DB,_=sql.Open("mysql",path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)

	if err:= DB.Ping(); err != nil{
		fmt.Println("<System message>:open database fail!")
		return
	}
	fmt.Println("<System message>:connect success!")

}
func ADD_DB(DATE string,AUTHOR string,MESSAGE string){
	path :=strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	DB,_=sql.Open("mysql",path)
	if err:= DB.Ping(); err != nil{
		fmt.Println("<System message>:open database fail!")
		return
	}
	//fmt.Println("<System message>:connect success!")
	stmt, err:= DB.Prepare("INSERT bottle_info SET bottle_date=?,bottle_author=?,bottle_message=?")
	if err != nil{
		fmt.Println(err)
		fmt.Println("PREPARE ERR!")
		return
	}

	res,err:= stmt.Exec(DATE,AUTHOR,MESSAGE)
	id,err := res.LastInsertId()
	if err!=nil{
		panic(err)
	}

	fmt.Println("Bottle_id is:",id)
	defer DB.Close()
}

func Query_DB(w http.ResponseWriter) {
	path :=strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	DB,_=sql.Open("mysql",path)
	if err:= DB.Ping(); err != nil{
		fmt.Println("<System message>:open database fail!")
		return
	}

	//type member struct {
	//	bottle_id int `DB :"bottle_id"`
	//	bottle_date string `DB :"bottle_date"`
	//	bottle_author string `DB :"bottle_author"`
	//	bottle_message string `DB :"bottle_message"`
	//}
	defer DB.Close()

	//查询数据库
	query, err := DB.Query("select * from bottle_info")
	if err != nil {
		fmt.Println("查询数据库失败", err.Error())
		return
	}
	defer query.Close()

	//读出查询出的列字段名
	cols, _ := query.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	//最后得到的map
	results := make(map[int]map[string]string)
	i := 1
	for query.Next() { //循环，让游标往下推
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	//fmt.Println(i)
	rand.Seed(time.Now().UnixNano())
	rad := rand.Intn(i)
	//fmt.Println(rad)
	fmt.Println("A random data ",results[rad])
	fmt.Fprintf(w,"A random data:%v",results[rad])
	//查询出来的数组
	//for k, v := range results {
	//	fmt.Println(k, v)
	//
	//}

	DB.Close() //用完关闭


}

func DEL_DB(ID int,w http.ResponseWriter){
	path :=strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	DB,_=sql.Open("mysql",path)
	if err:= DB.Ping(); err != nil{
		fmt.Println("<System message>:open database fail!")
		return
	}
	fmt.Println("<System message>:connect success!")

	stmt,err:=DB.Prepare("DELETE FROM bottle_info WHERE bottle_id=?")
	check(err)

	res,err:= stmt.Exec(ID)
	check(err)

	num,err:=res.RowsAffected()
	check(err)
	if num!=0{
		fmt.Println("Delete success!")
		fmt.Println("Delete id:",ID)
		fmt.Fprintln(w,"Detele id:",ID)
	}
	stmt.Close()


}
func check(err error){
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
}
//
//func main(){
//
//}

