package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func sender(conn net.Conn , macaddr string) bool {
	conn.Write([]byte(macaddr))
	fmt.Println("send over",macaddr)
	return true
}



func wol(who string) bool {
	var macaddr string
	statue := false
	switch who {
	case "/449158ec31537019": macaddr = "0c9d92bd7d15"//"kun"
	case "/c9a2599317183c9c": macaddr = "0c9d92bdfa36"//"chengxiang"
	default: macaddr = ""	
	}
	if macaddr==""{
		return statue
	}
	server := "yumn.tk:777"
	//server := "127.0.0.1:8001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		statue = false
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		statue = false
	}

	fmt.Println("connect success")
	statue = sender(conn , macaddr)
	return statue
}

func webResponse(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	wolpara := r.URL.Path
	fmt.Println("path", wolpara)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	statue := wol(wolpara)
	if !statue{
		fmt.Fprintf(w,"失败")
		return	
	} else {
		fmt.Fprintf(w,"成功!")
		return
	}	
}

func submitResponse(w http.ResponseWriter, r *http.Request){
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	wolpara := r.URL.Path
	fmt.Println("path", wolpara)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	// statue := wol(wolpara)
	// if !statue{
	// 	fmt.Fprintf(w,"失败")
	// 	return	
	// }
	// else{
	// 	fmt.Fprintf(w,"成功!")
	// 	return
	// }
}

func main() {
	http.HandleFunc("/",webResponse)
	http.HandleFunc("/da1e3053f72d38d8",submitResponse)

    err := http.ListenAndServe(":9000",nil)
    if err != nil{
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    }
}

