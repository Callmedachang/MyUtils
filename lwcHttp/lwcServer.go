package lwcHttp

import (
"net/http"
"log"
)

type Server struct {
	Host string
	Port string
}
func (s *Server) StartServer(){
	err := http.ListenAndServe(s.Host+":"+s.Port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func (s *Server) SetHandler(url string,handler http.HandlerFunc){
	http.HandleFunc(url, handler)
}
