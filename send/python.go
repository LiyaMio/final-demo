package send

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func LcssPrint(Keylength int,datasetnum int,thresholddataset int)(float64,float64){
	//cmd:=exec.Command("python","main.py","256", "1", "normal", "10")
	KeyLength:=strconv.Itoa(Keylength)

	thresholdDataset:=strconv.Itoa(thresholddataset)
	var chosen2 string
	if datasetnum>=1&&datasetnum<=100{
		chosen2 ="normal"

	}else if datasetnum>=101&&datasetnum<=200{
		chosen2 ="SHH-Taxi"
		datasetnum=datasetnum-100
	}else if datasetnum>=201&&datasetnum<=300{
		chosen2 ="T-drive"
		datasetnum=datasetnum-200
	}else if datasetnum>=301&&datasetnum<=400{
		chosen2 ="uniform"
		datasetnum=datasetnum-300
	}
	datasetNum:=strconv.Itoa(datasetnum)
	cmd:=exec.Command("python","main.py",KeyLength, datasetNum, chosen2, thresholdDataset)
	cmd.Dir="D:\\code\\mapweb\\algorithm\\lcss"
	output,err:=cmd.Output()
	if err!=nil{
		log.Fatalln(err)
	}
	tmp := strings.Split(string(output), "\r\n")
	a,_:=strconv.ParseFloat(tmp[0],64)
	b,_:=strconv.ParseFloat(tmp[1],64)
	fmt.Println(a)
	fmt.Println(b)
	return a,b
}

func DTWPrint(Keylength int,datasetnum int)(float64,float64){
	var chosen2 string
	if datasetnum>=1&&datasetnum<=100{
		chosen2 ="normal"

	}else if datasetnum>=101&&datasetnum<=200{
		chosen2 ="SHH-Taxi"
		datasetnum=datasetnum-100
	}else if datasetnum>=201&&datasetnum<=300{
		chosen2 ="T-drive"
		datasetnum=datasetnum-200
	}else if datasetnum>=301&&datasetnum<=400{
		chosen2 ="uniform"
		datasetnum=datasetnum-300
	}
	KeyLength:=strconv.Itoa(Keylength)
	datasetNum:=strconv.Itoa(datasetnum)
	cmd:=exec.Command("python","main.py",KeyLength, datasetNum, chosen2)
	cmd.Dir="D:\\code\\mapweb\\algorithm\\DTW"
	output,err:=cmd.Output()
	if err!=nil{
		log.Fatalln(err)
	}
	tmp := strings.Split(string(output), "\r\n")
	a,_:=strconv.ParseFloat(tmp[0],64)
	b,_:=strconv.ParseFloat(tmp[1],64)
	fmt.Println(a)
	fmt.Println(b)
	return a,b
}

func SecurePrint(Keylength int,datasetnum int,thresholddataset int)(float64,float64){
	//cmd:=exec.Command("python","main.py","256", "1", "normal", "10")
	KeyLength:=strconv.Itoa(Keylength)
	datasetNum:=strconv.Itoa(datasetnum)
	thresholdDataset:=strconv.Itoa(thresholddataset)
	var chosen2 string
	if datasetnum>=1&&datasetnum<=100{
		chosen2 ="normal"
	}else if datasetnum>=101&&datasetnum<=200{
		chosen2 ="SHH-Taxi"
		datasetnum=datasetnum-100
	}else if datasetnum>=201&&datasetnum<=300{
		chosen2 ="T-drive"
		datasetnum=datasetnum-200
	}else if datasetnum>=301&&datasetnum<=400{
		chosen2 ="uniform"
		datasetnum=datasetnum-300
	}
	cmd:=exec.Command("python","main.py",KeyLength,thresholdDataset,chosen2,  datasetNum)
	cmd.Dir="D:\\code\\mapweb\\algorithm\\Secure-Trajectory-Similarity-Computation"
	output,err:=cmd.Output()
	if err!=nil{
		log.Fatalln(err)
	}
	tmp := strings.Split(string(output), "\r\n")
	a,_:=strconv.ParseFloat(tmp[0],64)
	b,_:=strconv.ParseFloat(tmp[1],64)
	fmt.Println(a)
	fmt.Println(b)
	return a,b
}
//func MyData(cnt [1000]int,leng int)(lcssans [100]float64,dtw [100]float64) {
//	//for i := 0; i < leng;i++{
//	//	_ ,lcssans[i]=LcssPrint(256,cnt[i],200)
//	//	_ ,dtw[i]=DTWPrint(256,cnt[i])
//	//}
//	_ ,dtw[0]=DTWPrint(256,24)
//	_ ,dtw[1]=DTWPrint(256,25)
//	_ ,dtw[2]=DTWPrint(256,21)
//	_ ,dtw[3]=DTWPrint(256,6)
//	_ ,dtw[4]=DTWPrint(256,8)
//	return lcssans,dtw
//}
