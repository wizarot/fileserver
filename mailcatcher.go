package main

import (
	"fmt"
	"go-mailcatcher/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// func main() {
// 	mux := http.NewServeMux()

// 	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
// 	mux.Handle("/static", http.NotFoundHandler())
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	err := http.ListenAndServe(":3000", mux)
// 	log.Fatal(err)
// }

// // 使用自定义的文件系统,并将其传递给 http.FileServer
// type neuteredFileSystem struct {
// 	fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
// 	f, err := nfs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if s.IsDir() {
// 		// 如果是目录,那么给出一个tmp?
// 		index := strings.TrimSuffix(path, "/") + "/static/example.html"
// 		if _, err := nfs.fs.Open(index); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return f, nil
// }

func main() {
	// 自己复制一个fileServer然后照着修改!
	fs1 := utils.FileServer(utils.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs1))
	// static目录下使用 fileServer
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 其它路径使用serverTemplate处理
	// http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	// lp := filepath.Join("views", "layout.html")
	// // 解析r.URL.Path到对应tpl
	// fp := filepath.Join("views", filepath.Clean(r.URL.Path))

	// // Return a 404 if the template doesn't exist
	// info, err := os.Stat(fp)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		http.NotFound(w, r)
	// 		return
	// 	}
	// }

	path := fmt.Sprintf(".%s", filepath.Clean(r.URL.Path))
	fmt.Println("path:", path)
	fi, err := os.Stat(path)
	if err != nil {
		println("not found")
		http.NotFound(w, r)
		return
	}

	// 文件 / 目录 TODO:参考 FileService
	if fi.IsDir() {
		// 目录:列出来
		dir, _ := os.Open(path)
		files, _ := dir.Readdir(-1)
		dirFiles := "" //暂时先输出成一个字符串
		for _, v := range files {
			// 跳过隐藏文件
			if !strings.HasPrefix(v.Name(), ".") {
				// fmt.Println(v.IsDir())
				dirFiles += v.Name() + "</br>"
			}

		}
		fmt.Fprintf(w, "Dir, %q", dirFiles)
	} else {
		// 文件:传回去
		// 读取文件
		fmt.Fprintf(w, "File, %q", path)
	}

	return

	// fi, _ := dir.Readdir(-1)
	// for _, v := range fi {
	// 	fmt.Println(v.Name())
	// }

	// // Return a 404 if the request is for a directory
	// // TODO: 目录应当显示对应dir
	// if info.IsDir() {
	// 	// 读取当前目录dir内文件

	// 	println(path)
	// 	dir, _ := os.Open(path)
	// 	fi, _ := dir.Readdir(-1)
	// 	for _, v := range fi {
	// 		fmt.Println(v.Name())
	// 	}
	// 	http.NotFound(w, r)
	// 	return
	// }

	// tmpl, err := template.ParseFiles(lp, fp)
	// if err != nil {
	// 	// Log the detailed error
	// 	log.Println(err.Error())
	// 	// Return a generic "Internal Server Error" message
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return
	// }

	// if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, http.StatusText(500), 500)
	// }
}
