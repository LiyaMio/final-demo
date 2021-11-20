package send

import (
	"strconv"
)
func Select(c string) int64{
	c = string([]byte(c)[4:])
	num,_:= strconv.ParseInt(c, 10, 32)
	return num
}






type Simliarty struct{
	X string `json:"Algorithm"`
	Y float64 `json:"Simliarty"`
}
var s Simliarty
func Select1()Simliarty{
	s.Y=1
	s.X="SLCSS"
	return s
}
var ss Simliarty
func Select2()Simliarty{
	ss.Y=1
	ss.X="SDTW"
	return ss
}
var sss Simliarty
func Select3()Simliarty{
	sss.Y=1
	sss.X="SBD Similarity"
	return sss
}