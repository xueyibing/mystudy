package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type three struct {
	T1 string	`json:"t1"`
	T2 string	`json:"t2"`
}

//结构体中的字段首字母大写，不然那marshal失败
type data struct {

	One string	`json:"one1"`
	Two string	`json:"two2"`
	Three three `json: "three"`

}


func main()  {


	unmarshal_file()
	return
	unmarshal()

	return
	da := data{
		"one",
		"two",
		three{
			"1",
			"2",
		},
	}
	ds , err := json.Marshal(da)
	fmt.Println(err)
	fmt.Println(string(ds))

	dm := make(map[string]interface{})
	json.Unmarshal(ds,&dm)

	fmt.Println(dm["Three"].(map[string]interface{})["t1"])


	//fmt.Println(data)


}

func unmarshal()  {

	type scope struct {
		Bucket string	`json:"scope"`
		Deadline uint32 `json:"deadline"`
	}

	var s scope
	js := `{"scope":"b2","deadline":1642444718}`

	err := json.Unmarshal([]byte(js),&s)

	fmt.Println(err)

	fmt.Println(s)


}

func unmarshal_file()  {

	data, err := ioutil.ReadFile("./json.conf")
	if err != nil {

		return
	}

	err = json.Unmarshal(data, nil)
	if err != nil {

	}


}