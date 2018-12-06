package main

import "fmt"

//////////////
func ret_adder() func(int,int)int {

	seed := 10
	return func(i int, i2 int) int {

		return i + i2 + seed
	}

}
func exec_1()  {
	i := 10
	i2 := 11
	adder := ret_adder()
	fmt.Println(adder(i,i2))
}






func main()  {

	exec_1()


}
