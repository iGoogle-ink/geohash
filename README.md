# geohash

<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>

# 使用手册

## 安装
```bash
$ go get -u github.com/iGoogle-ink/geohash
```

## 计算geohash
```go
//计算geohash
//    lat：纬度
//    lng：纬度
//    precision：精度值
geohash := geohash.Encode(39.928167, 116.389550, 8)
fmt.Println(geohash)
```