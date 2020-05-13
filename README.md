# FileServer 简易文件服务器

2020年武汉肺炎,在家里学习GO语言,顺便做一点东西出来
这个项目的作用是fileServer,但是有样式,又可以直接播放 mp3 mp4文件.

## 打包方式
安装 go.rice 工具把静态文件都放到可执行文件中,这样在任何地方执行访问到静态文件

```
# 生成fileserver可执行文件
go build
# 将静态文件附加到fileserver中
rice append --exec fileserver
# 编译windows版可执行文件exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build fileserver.go
# 追加文件
rice append --exec fileserver.exe
```
## 使用方式
1. 把编译好的可执行文件 fileserver 放到$PATH中.
2. 在需要分享的目录中,执行fileserver,然后就能访问了!

v1.0
参数 --port=3000 可以设置监听端口

新版本v1.1
增加参数 --username=admin --password=123 当password不为空,需要输入密码才能登陆

## 新特性feature:TODO
* [x] 加一个简单密码验证,可以设定访问密码
* [ ] 找专家调一调前端,适应手机端访问
* [ ] 通过增加dns向局域网广播一个hostname?? 似乎有bug..
* [x] 命令行二维码扫码下载或访问
* [ ] 上传文件功能
* [ ] 改用go module进行包管理,精简最终包大小
