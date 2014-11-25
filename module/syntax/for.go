package main

import (
	"fmt"
)

/*
for index,char := range string {}
for index,value := range array {}
for index,value := range slice {}
for key,value := range map {}
*/
func main() {
	for i := 0 ; i < 5 ; i++ {
		fmt.Println(i)
	}

	for pos,val := range "测试一下ABC" { // rune
		fmt.Printf("character '%c' type is %T value is %s adn start at byte postion %d \r\n",val,val,val,pos)
		fmt.Printf("string(%v)=> %s\r\n",val,string(val))
	}
}