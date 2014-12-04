package main

import(
	"fmt"
	"strings"
)

// echo implode(",",array(1,2,3,4)); // 1,2,3,4
// echo join(",",array(1,2,3,4)); // 1,2,3,4
func main() {
	arr := []string{"1","2","3","4"}
	fmt.Println(strings.Join(arr,","))
}