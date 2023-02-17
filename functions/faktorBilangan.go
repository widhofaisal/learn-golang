package functions

import ("fmt")

func FaktorBilangan(N int){
	for i:=1; i<=N; i++{
		if N%i==0{
			fmt.Println(i)
		}
	}
}