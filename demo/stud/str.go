package stud

import (
	"fmt"
	"unicode/utf8"
)

/*
----------------- string ------------
 * len() 注意：中文以utf-8保存，每个占3个字节，实际长度可以通过 utf8.RuneCountInString
 * 遍历字符串
*/
func StrTest() {
	//长度
	str := "Happy New Year!"
	str2 := "新年快乐"
	fmt.Println(len(str))                     //15
	fmt.Println(len(str2))                    //12,ascll字符串长度，每个中文3个字节
	fmt.Println(utf8.RuneCountInString(str2)) //4,unicode长度
	//遍历

}
