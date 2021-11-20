package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mapweb/send"
	"net/http"
	"strconv"
)

var key int64
var chooseData [1000]int    //哪些数据集
var chosen  [30]string      //选择哪种算法
var ci int              //算法计数
var i int               //计数
func main() {
	i = 0
	ci =0
	r := gin.Default()
	r.LoadHTMLGlob("HTML/*")
	r.StaticFS("/views", http.Dir("./views/"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new map.html", nil)
	})
	r.GET("/line.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "line.html", nil)
	})

	r.POST("/showdata", func(c *gin.Context) {
		//value1:=c.PostForm("Amenu")
		value2 := c.PostForm("Bmenu")
		//fmt.Println("val:",value1,value2)
		key = send.Select(value2)
		fmt.Println(key)
	})
	r.GET("/intidata", func(c *gin.Context) {
		datab:=send.Yread(key)
		//fmt.Println(datab)
		c.JSON(http.StatusOK,gin.H{
			"data":datab,
		})

	})

	r.GET("/data", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "newmap.html", nil)
		datas := send.Read(key)
		//datas:=send.Try()
		c.JSON(http.StatusOK, gin.H{
			"data": datas,
		})
	}) //传数据集(在地图上显示数据)

	r.POST("/chooseData", func(c *gin.Context) {
		//value1:=c.PostForm("Amenu")
		value := c.PostForm("Bmenu")
		//chooseData[i],_= strconv.ParseInt(value, 10, 32)
		chooseData[i], _ = strconv.Atoi(value)
		fmt.Println("val:", chooseData[i])
		i++
	}) //获得选择的点



	r.POST("/alg", func(c *gin.Context) {
		chosen[ci]=c.PostForm("id")
		ci++
		fmt.Println("val:", chosen)
	}) //接受DTW

	r.GET("/sumN", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": ci,
		})
	})
	//if chosen=="TEST"{
	//
	//} else if chosen=="DTW"{
	//
	//}else if chosen=="LCSS"{
	//
	//}
	//判断算法

	//3个图
	r.GET("/dataTime.json", func(dataTime *gin.Context) {
		//var keylengthtimelcss [100]float64
		//var keylengthtimedtw [100]float64
		////var keylengthtime3 [100]float64
		//var ke1 float64
		//ke1 = 0
		////var ke2 float64
		////ke2 = 0
		////var ke3 float64
		////ke3=0
		//var Keylength = [3]int{256, 512, 1024}
		//for k := 0; k < 3; k++ {
		//	for j := 0; j < i; j++ {
		//		//_, ke2 = send.DTWPrint(Keylength[k], chooseData[j])
		//		//keylengthtimedtw[k] = keylengthtimedtw[k] + ke2
		//		_, ke1 = send.LcssPrint(Keylength[k], chooseData[j], 10)
		//		keylengthtimelcss[k] = keylengthtimelcss[k] + ke1
		//		//_,ke3 = send.SecurePrint(Keylength[k],chooseData[j],10)
		//		//keylengthtime3[k]=keylengthtime3[k]+ke3
		//	}
		//}
		//调用算法
		te := send.Runtime(chosen,ci)
		fmt.Println(te)
		dataTime.JSON(http.StatusOK, gin.H{"data": te})
	})
	r.GET("/dataAverage.json", func(c *gin.Context) {
		//var keytoplcss [100]float64
		//var keytopdtw [100]float64
		//var keythreshold3 [100]float64

		average := send.Average(chosen,ci)
		//fmt.Println(average)
		c.JSON(http.StatusOK, gin.H{"data": average})
	})
	r.GET("/dataMemory.json", func(c *gin.Context) {
		//var keythresholdlcss [100]float64
		//var keythresholddtw [100]float64
		////var keythreshold3 [100]float64
		//var dtwkeythreshold float64
		//var ke1 float64
		//ke1 = 0
		//var ke2 float64
		//ke2 = 0
		//var ke3 float64
		//ke3=0
		//for j := 0; j < 4; j++ {
		//	_, ke2 = send.DTWPrint(256, chooseData[j])
		//	dtwkeythreshold = dtwkeythreshold + ke2
		//}
		//var Threshold = [4]int{10, 50, 100, 200}
		//for k := 0; k < 4; k++ {
		//	for j := 0; j < i; j++ {
		//		_, ke1 = send.LcssPrint(256, chooseData[j], Threshold[k])
		//		keythresholdlcss[k] = keythresholdlcss[k] + ke1
		//		//_,ke3 = send.SecurePrint(256,chooseData[j],Threshold[k])
		//		//keythreshold3[k]=keythreshold3[k]+ke3
		//	}
		//	keythresholddtw[k] = dtwkeythreshold
		//}
		me := send.Mermory(chosen,ci)
		//fmt.Println(me)
		c.JSON(http.StatusOK, gin.H{"data": me})
	})


	//echarts


	//图表
	r.GET("/alg1", func(c *gin.Context) {
		a:=send.Select1()
		c.JSON(http.StatusOK, gin.H{
			"data": a,
		})
	})
	r.GET("/alg2", func(c *gin.Context) {
		a:=send.Select2()
		c.JSON(http.StatusOK, gin.H{
			"data": a,
		})
	})
	r.GET("/alg3" , func(c *gin.Context) {
		a:=send.Select3()
		c.JSON(http.StatusOK, gin.H{
			"data": a,
		})
	})





	r.Run(":8090")
}


//
//func main(){
//	key=1
//	send.Yread(key)
//}
//func main(){
//	//send.DTWPrint(256,1)
//	var a,b,c,d,e,v float64
//	//_,a=send.LcssPrint(256,1,10)
//	//_,b=send.LcssPrint(256,2,10)
//	//_,c=send.LcssPrint(256,3,10)
//	//_,d=send.LcssPrint(256,4,10)
//	//_,e=send.LcssPrint(256,5,10)
//	//v=a+b+c+d+e
//	//fmt.Println("256:",v)
//	//
//	//_,a=send.LcssPrint(512,1,10)
//	//_,b=send.LcssPrint(512,2,10)
//	//_,c=send.LcssPrint(512,3,10)
//	//_,d=send.LcssPrint(512,4,10)
//	//_,e=send.LcssPrint(512,5,10)
//	//v=a+b+c+d+e
//	//fmt.Println("512:",v)
//	//
//	//_,a=send.LcssPrint(1024,1,10)
//	//_,b=send.LcssPrint(1024,2,10)
//	//_,c=send.LcssPrint(1024,3,10)
//	//_,d=send.LcssPrint(1024,4,10)
//	//_,e=send.LcssPrint(1024,5,10)
//	//v=a+b+c+d+e
//	//fmt.Println("1024:",v)
//	_,a=send.DTWPrint(256,1)
//	_,b=send.DTWPrint(256,2)
//	_,c=send.DTWPrint(256,3)
//	_,d=send.DTWPrint(256,4)
//	_,e=send.DTWPrint(256,5)
//	v=a+b+c+d+e
//	fmt.Println("256:",v)
//
//	_,a=send.DTWPrint(512,1)
//	_,b=send.DTWPrint(512,2)
//	_,c=send.DTWPrint(512,3)
//	_,d=send.DTWPrint(512,4)
//	_,e=send.DTWPrint(512,5)
//	v=a+b+c+d+e
//	fmt.Println("512:",v)
//
//	_,a=send.DTWPrint(1024,1)
//	_,b=send.DTWPrint(1024,2)
//	_,c=send.DTWPrint(1024,3)
//	_,d=send.DTWPrint(1024,4)
//	_,e=send.DTWPrint(1024,5)
//	v=a+b+c+d+e
//	fmt.Println("1024:",v)
//
//
//}