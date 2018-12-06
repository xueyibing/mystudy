package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main()  {

	fmt.Println(os.Args)

	s :=base64.URLEncoding.EncodeToString([]byte("POST:/upload"))

	fmt.Println(s)
	//err := errors.New("dsd")

	//e := fmt.Errorf("%s",err.Error())

	fmt.Println(strings.TrimRight("1234:456/","/"))
	ss := "1234:456"
	fmt.Println( ss[0: strings.Index(ss,":")])

}