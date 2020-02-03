package main

import (
	"fileserver/utils"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
)

// 获取本机ip
func GetIP() string {
	localIP := ""
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = ipnet.IP.String()
			}
		}
	}
	return localIP
}

// 设置访问的用户名,密码
func authentication(next http.Handler, username string, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if password != "" { //密码为空,那么就不验证了
			if !ok || user != username || pass != password {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"Demo Test\"")
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 获取本机IP
	localIP := GetIP()
	// 获取监听端口,建议mac电脑大于1024,否则需要sudo
	var port string
	flag.StringVar(&port, "port", "3000", "指定要使用的端口,默认是3000,建议Mac电脑不要使用小于1024的端口,否则需要sudo --port=3000") // -port=9000输入
	var username string
	flag.StringVar(&username, "username", "admin", "指定要使用的用戶名,默认 --username=admin")
	var password string
	flag.StringVar(&password, "password", "", "指定要使用的密码,默认为空则不需要输入 --password=123")

	flag.Parse()

	box1 := rice.MustFindBox("views")
	viewsFileServer := http.StripPrefix("/views/", http.FileServer(box1.HTTPBox()))
	http.Handle("/layout/", viewsFileServer)

	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	http.Handle("/static/", staticFileServer)

	// 自己复制一个fileServer然后照着修改!
	fs1 := utils.FileServer(utils.Dir("./"))
	http.Handle("/", authentication(http.StripPrefix("/", fs1), username, password))

	log.Println("http://"+localIP+":"+port, "Listening...")
	if password != "" {
		log.Println("Auth username:" + username + " password:" + password)
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
