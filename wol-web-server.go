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



func wol(magicword string) bool {
	var macaddr string
	statue := false
	switch magicword {
	case "word1": macaddr = "macaddr1"
	case "word2": macaddr = "macaddr2"
	default : return statue
		
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

// func webResponse(w http.ResponseWriter, r *http.Request)  {
// 	r.ParseForm()  //解析参数，默认是不会解析的
// 	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
// 	wolpara := r.URL.Path
// 	fmt.Println("path", wolpara)
// 	//fmt.Println("scheme", r.URL.Scheme)
// 	//fmt.Println(r.Form["url_long"])
    
	
// 	statue := wol(wolpara)
// 	if !statue{
// 		fmt.Fprintf(w,"失败")
// 		return	
// 	} else {
// 		fmt.Fprintf(w,"成功!")
// 		return
// 	}	
// }

func submitResponse(w http.ResponseWriter, r *http.Request){
	statue := false
	fmt.Fprintf(w,"Parsing the form data...\n")
	r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(w, r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Fprintf(w,"Sanding waking command...\n")
	wolpara := r.PostFormValue("magic word")
	//fmt.Println(wolpara)
	//fmt.Println("scheme", r.URL.Scheme)
	if wolpara != "Your Magic Word"{
		statue = wol(wolpara)
	}else{
		fmt.Fprintf(w,"Type in your magic word please\n")
	}
	if !statue{
		fmt.Fprintf(w,"Failed!\n")
		return	
	}else{
		fmt.Fprintf(w,"Succeed!\n")
		return
	}
}

func main() {
	//http.HandleFunc("/",webResponse)
	http.HandleFunc("/da1e3053f72d38d8",submitResponse)

    err := http.ListenAndServe(":9000",nil)
    if err != nil{
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    }
}

