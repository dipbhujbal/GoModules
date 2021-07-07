package Father_p

import (
	"fmt"
)
type Father_struct struct{
name string
}

func (F Father_struct) Display_father(fname string) {
F.name = fname
fmt.Printf("Father name is %s Bhujbal\n", F.name)

}
