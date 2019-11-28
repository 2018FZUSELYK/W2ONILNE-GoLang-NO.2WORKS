package server

import (
	"W2ONLINE/AssessmentROUND2/bottlehtml/btm/database_set"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)


//type MyMux struct {
//}

type MyForm struct {
	NAME string
	DATE string
	MESSAGE string
}

//func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	if r.URL.Path == "/"{
//		SayHelloName(w,r)
//		return
//	}
//	if r.URL.Path == "/about"{
//		About(w,r)
//		return
//	}
//	if r.URL.Path == "/login"{
//		Login(w,r)
//		return
//	}
//	http.NotFound(w,r)
//	return
//}

func SayHelloName(w http.ResponseWriter,r *http.Request){
	//解析url传递的参数
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path) // r.URL类里面有关于URL的相关方法和属性
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "欢迎来到漂流瓶!") //输出到客户端的信息


}

func About(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"The Test Of Server.\n")
}

func Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("bottle.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()


		Bottle := r.Form["bottle"]
		Date := r.Form["date"]
		Author :=	r.Form["author"]
		Message :=r.Form["message"]
		fmt.Println("Bottle:",Bottle)
		fmt.Println("Date:",Date)
		fmt.Println("Author:",Author)
		fmt.Println("Message:",Message)
		for i,v :=range Bottle{
			fmt.Println(i)
			fmt.Fprintf(w,"Bottle:%v\n",v)
		}
		for k,n :=range Date {
			fmt.Println(k)
			fmt.Fprintf(w, "Date:%v\n", n)
		}
		for i,v :=range Author{
			fmt.Println(i)
			fmt.Fprintf(w,"Name:%v\n",v)
		}
		for j,b :=range Message{
			fmt.Println(j)
			fmt.Fprintf(w,"Message:%v\n",b)
		}
		//database_set.ADD_DB("2019-11-11","LYK","TEST")
		database_set.ADD_DB(Date[0],Author[0],Message[0])

	}
}

func Query(w http.ResponseWriter,r *http.Request){
	database_set.Query_DB(w)
}

func Delete(w http.ResponseWriter,r *http.Request){
	rand.Seed(time.Now().UnixNano())
	database_set.DEL_DB(rand.Intn(15),w)
}

