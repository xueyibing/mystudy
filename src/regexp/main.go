package main

import (
	"fmt"
	"regexp"
)

func main()  {


	sub()
	match()
}

func sub()  {

	s := "/mkbucketv2/([^/\\\\]+)/"
	reg,_ :=regexp.Compile(s)

	subs := reg.FindStringSubmatch(("/mkbucketv2/1234/"))

	for _,s:= range subs{

		fmt.Println(s)

	}
}

func match()  {

	reg :=regexp.MustCompile("^/[^/\\\\]+/[^/\\\\]+")

	fmt.Println(reg.Match([]byte("/f2.fds==/434")))

}