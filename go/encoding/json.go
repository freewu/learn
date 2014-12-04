package main

import (
	"fmt"
	"encoding/json"
	//"os"
)

// echo json_encode(array("a" => 1,"b" => "x","c" => array(1,2,3))); // {"a":1,"b":"x","c":[1,2,3]}
// var_dump(json_decode('{"a":1,"b":"x","c":[1,2,3]}',true));
/*
array(3) {
  ["a"]=>
  int(1)
  ["b"]=>
  string(1) "x"
  ["c"]=>
  array(3) {
    [0]=>
    int(1)
    [1]=>
    int(2)
    [2]=>
    int(3)
  }
}
*/

type TestJson struct {
	A int   `json:"a"`  // 使用json定义看看
	B string
	C []int
}

func main() {
	// encode struct
	// stuct 不能小写的。。。
	j := TestJson {
		A : 1,
		B : "x",
		C : []int{1,2,3},
	}
	b,err := json.Marshal(j);
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	//os.Stdout.Write(b)
	fmt.Println(string(b))  // {"A":1,"B":"x","C":[1,2,3]}

	// decode stuct
	// 字符要处理成byte
	//b1 := []byte(`{"A":1,"B":"x","C":[1,2,3]}`)
	b1 := []byte(`{"a":1,"B":"x","C":[1,2,3]}`)
	var j1 TestJson
	err = json.Unmarshal(b1,&j1)
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	fmt.Printf("%+v\r\n",j1)
	//fmt.Printf("%x\r\n",j1)

	// encode map
	m := map[string]int{"apple":5,"pear":6}
	mb,err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	fmt.Println(string(mb)) // {"apple":5,"pear":6}
	// decode map
	//var md map[string]int;
	md := make(map[string]int)
	mb1 := []byte(`{"apple":5,"pear":6}`)
	err = json.Unmarshal(mb1,&md)
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	fmt.Printf("%+v\r\n",md)

	// map[string]interface{}
	m3 := make(map[string]interface{})
	m3["a"] = 1
	m3["b"] = "x"
	m3["c"] = []int{1,2,3}
	fmt.Printf("%+v\r\n",m3)

	// encode
	mb3,err := json.Marshal(m3);
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	fmt.Println(string(mb3))

	// decode
	md3 := make(map[string]interface{})
	err = json.Unmarshal(mb3,&md3)
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	fmt.Printf("%+v\r\n",md3)
}
