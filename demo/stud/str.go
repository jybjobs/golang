package stud

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
----------------- string ------------
 * len() 注意：中文以utf-8保存，每个占3个字节，实际长度可以通过 utf8.RuneCountInString
 * 遍历字符串
 * 修改字符串：不能直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现。
 * 连接字符串： + 或 stringBuilder
*/
func StrTest() {
	//1. 长度
	str := "Happy New Year!"
	str2 := "新年快乐"
	fmt.Println(len(str))                     //15
	fmt.Println(len(str2))                    //12,ascll字符串长度，每个中文3个字节
	fmt.Println(utf8.RuneCountInString(str2)) //4,unicode长度
	//2. 遍历
	for i := 0; i < len(str); i++ { //ascii遍历
		fmt.Printf("ascii:%c %d\n", str[i], str[i])
	}

	for _, s := range str2 { //unicode遍历
		fmt.Printf("unicode: %c %d \n", s, s)
	}
	//3. 字符串查询/截取
	str3 := "hello,world"
	n := strings.Index(str3, ",")
	w := strings.Index(str3[n:], "w")
	fmt.Println(str3, n, w)
	//4. 字符串连接
	var stringBuilder bytes.Buffer //声明字节缓冲区
	stringBuilder.WriteString(str) //把字符串写入缓冲区
	stringBuilder.WriteString(str2)
	fmt.Println(stringBuilder.String())
	//5. base64
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(str)) //编码
	fmt.Println(encodedMessage)
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
