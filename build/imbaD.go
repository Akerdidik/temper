package main

import (
	"fmt"
	"os"
)
func pow10(n int)int{
	res:=1
	for i:=0;i<n;i++{
		res*=10
	}
	return res
}
func atoi(s string)int{
	res:=0
	runes:=[]rune(s)
	lano:=len(runes)-1
	for i:=0;i<len(runes);i++{
		temp:=int(runes[i]-48)*pow10(lano-i)
		res+=temp
	}
	return res
}
func QuadD(x, y int) {
	if x < 0 || y < 0 {
		fmt.Print("")
		return
	} else {
		for h := 0; h < y; h++ {
			for w := 0; w < x; w++ {
				if w == 0 && (h == 0 || h == y-1) { 
					fmt.Print("A")
				} else if w > 0 && w < x-1 && (h == 0 || h == y-1) || h < y-1 && h > 0 && (w == 0 || w == x-1) {
					fmt.Print("B")
				} else if (h == y-1 || h == 0) && (w == 0 || w == x-1) {
					fmt.Print("C")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
func main(){
	args:=os.Args[1:]
	first:=atoi(args[0])
	second:=atoi(args[1])
	QuadD(first,second)
}