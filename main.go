package main

import (
	"W2ONLINE/AssessmentROUND2/bottlehtml/btm/server"
	"log"
	"net/http"
)



func main(){
	//mux := &server.MyMux{}
	http.HandleFunc("/",server.SayHelloName)
	http.HandleFunc("/login",server.Login)
	http.HandleFunc("/about",server.About)
	http.HandleFunc("/query",server.Query)
	http.HandleFunc("/delete",server.Delete)
	err:=http.ListenAndServe(":9090",nil)
	//database_set.Query_DB()
	if err!=nil{
		log.Fatal("ListenAndServe",err)
	}

}
