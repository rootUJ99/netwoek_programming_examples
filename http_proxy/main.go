package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ServerConfig struct {
	port string
	path string
	message string
}

type Handlers interface {
	handleListen()
	handleFunc()
}

func (s ServerConfig) handleListen() {
	http.ListenAndServe(s.port, nil)
}

func (s ServerConfig) handleFunc() {
	http.HandleFunc(s.path ,func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, fmt.Sprintf("%s from %s\n", s.message, s.port));	
		w.Write([]byte(fmt.Sprintf("%s from %s", s.message, s.port)));
	})
}

func (s ServerConfig) handleFuncProxy(){
	http.HandleFunc(s.path ,func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("%s from %s\n", s.message, s.port));	

		targetUrl, err :=url.Parse("http://127.0.0.1:8080/target")
		if err != nil {
			fmt.Println("invalid origin url")
		}
		fmt.Println(targetUrl)

		r.Host = targetUrl.Host
		r.URL.Host = targetUrl.Host
		r.URL.Scheme = targetUrl.Scheme
		r.RequestURI = ""

		targetServerRes, err := http.DefaultClient.Do(r)

		if err != nil {
			fmt.Println("reaching in err", err)
			w.WriteHeader(http.StatusInternalServerError)	
			_,_ = fmt.Fprint(w, err)
			return
		}
		// fmt.Println("reaching here")
		w.WriteHeader(http.StatusOK)	
		io.Copy(w, targetServerRes.Body)
	})

}

func main(){
	target:= ServerConfig{ 
		port: ":8080",
		path: "/target",
		message: "hellow from target",
	}
	target.handleFunc()
	proxy:= ServerConfig{ 
		port: ":8090",
		path: "/proxy",
		message: "hellow from proxy",
	}
	proxy.handleFuncProxy()

	go target.handleListen()
	proxy.handleListen()

	select {}
}
