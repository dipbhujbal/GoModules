package Sonp

import (
	"fmt"
)
type Son_struct struct {
name string

}


func (s Son_struct ) Display_son(sname string) {
s.name=sname
fmt.Printf("Sons name is %s Bhujbal", s.name)
}






