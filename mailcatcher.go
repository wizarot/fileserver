package main

import (
	"flag"
	"fmt"
	"go-mailcatcher/utils"
	"log"
	"net"
	"net/http"
	"os"

	// "github.com/gobuffalo/packr/v2"
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

func main() {
	// 获取本机IP
	localIP := GetIP()
	// 获取监听端口,建议mac电脑大于1024,否则需要sudo
	var port string
	flag.StringVar(&port, "port", "3000", "指定要使用的端口,默认是3000,建议Mac电脑不要使用小于1024的端口,否则需要sudo") // -port=9000输入
	flag.Parse()

	// static目录下使用 fileServer
	// box := packr.New("static","./static")
	// http.Handle("/static/", http.FileServer(box))
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	http.Handle("/static/", staticFileServer)

	// 自己复制一个fileServer然后照着修改!
	fs1 := utils.FileServer(utils.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs1))

	// 其它路径使用serverTemplate处理
	// http.HandleFunc("/", serveTemplate)

	log.Println("http://"+localIP+":"+port, "Listening...")
	http.ListenAndServe(":"+port, nil)
}
