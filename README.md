# 2020 开始一个小项目

2020年武汉肺炎,在家里学习GO语言,顺便做一点东西出来
这个项目的作用是fileServer,但是有样式,又可以直接播放 mp3 mp4文件.

## 打包方式
安装 go.rice 工具把静态文件都放到可执行文件中,这样在任何地方执行访问到静态文件

```
# 生成fileserver可执行文件
go build
# 将静态文件附加到fileserver中
rice append --exec fileserver
```
