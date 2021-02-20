package main

import (
	"flag"
	"fmt"
	"strings"
	"unsafe"
)

var f = flag.String("f", "xtf", `
指定方式
xtf 小驼峰命名
dtf 大驼峰
xh 下划线
dx  大写
xx  小写

`)
var s = flag.String("s", "", "要格式化的字符串")

type Handler func(str string) string

func main() {

	flag.Parse()

	option := *f
	str := *s

	funcsMap := make(map[string]Handler)
	funcsMap["xtf"] = xtf
	funcsMap["dtf"] = dtf
	funcsMap["xh"] = xh
	funcsMap["dx"] = dx
	funcsMap["xx"] = xx


	options := parseOption(option)
	res := str
	for _, s2 := range options {
		res = funcsMap[s2](res)
	}
	fmt.Println(res)
}

func parseOption(option string) []string {
	return strings.Split(option, "-")
}

func xtf(str string) string {
	res := tf(str)
	resArray := []byte(res)
	resArray[0] += 32
	return *(*string)(unsafe.Pointer(&resArray))
}

func dtf(str string) string  {
	return tf(str)
}

func tf(str string) string  {
	strArr := strings.Split(str, "_")
	for idx, s := range strArr {
		lowerS := []byte(strings.ToLower(s))
		lowerS[0] -= 32
		ss := *(*string)(unsafe.Pointer(&lowerS))
		strArr[idx] = ss
	}
	return strings.Join(strArr, "")
}

func xh(str string) string {
	arr := make([]string, 0)

	s := []byte(str)
	l := 0
	h := 0
	for idx := 0; idx < len(s); idx++ {
		if idx != 0 && s[idx] >= 'A' && s[idx] <= 'Z' {
			h = idx
			ns := s[l:h]
			arr = append(arr, *(*string)(unsafe.Pointer(&ns)))
			l = h
		}
	}
	ns := s[l:len(s)]
	arr = append(arr, *(*string)(unsafe.Pointer(&ns)))
	res := strings.Join(arr, "_")
	return res

}

func dx(str string) string  {
	return strings.ToUpper(str)
}

func xx(str string) string {
	return strings.ToLower(str)
}