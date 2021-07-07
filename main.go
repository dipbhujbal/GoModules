package main

import (
      
      parent  "github.com/dipbhujbal/GoModules/Father"
      child   "github.com/dipbhujbal/GoModules/Father/Son"

)

func main(){

	p1:= new (parent.Father_struct)
	c1:=new (child.Son_struct)
	p1.Display_father("Sunil")
	c1.Display_son("Akshay")






}


