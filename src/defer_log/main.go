package main

import (
	"fmt"
	"github.com/qiniu/log"
	"qbox.us/errors"
)

type UrlBindAction struct {
Url       string        `json:"url" bson:"url"`               // url格式   ex:POST:/upload
Action    string        `json:"action" bson:"action"`         // 动作名	ex:CreateObject
Desc      string        `json:"desc" bson:"desc"`             // 描述
}

func main()  {

	u := &UrlBindAction{
		"u",
		"a",
		"d",
	}

testlog("e",u)

}

func testlog(s1 string, u *UrlBindAction ) (err error)  {

	defer func () {
		if err != nil {
			log.Errorf("testlog err,s1:%s, u:%v, err:%v",s1, u,err)
		}
	}()

	s,err := "",errors.New("this is a err")
	fmt.Println(s)

	ss := []string{"1","2","3"}

	fmt.Printf("%s",ss)

	if err != nil {

		return
	}
	return
}
