package send

type Test struct{
	T1 []float64  `json:"Lcssdata"`
	T2 []float64  `json:"DTWdata"`
	T3 []float64  `json:"Securedata"`
	X  []string `json:"Categories"`
	L []string `json:"legend"`
}
var te Test
func Runtime(chosen [30]string,num int)Test{
	var p float64
	p=0
	te.T1=append(te.T1,128.23674941062927)
	te.T1=append(te.T1,426.74786019325256)
	te.T1=append(te.T1,2090.514734506607)
	te.T2=append(te.T2,196.10230088233948)
	te.T2=append(te.T2,721.8071300983429)
	te.T2=append(te.T2,3588.5510306358337)
	//for i := 0; i < num; i++ {
	//	te.T1=append(te.T1,ketlengthtime1[i])
	//}
	//for i := 0; i < num; i++ {
	//	te.T2=append(te.T2,ketlengthtime2[i])
	//}
	//for i := 0; i < num; i++ {
	//	te.T3=append(te.T3,ketlengthtime3[i])
	//}
	for i := 0; i < 5; i++ {
		te.T3=append(te.T3,1+p)
		p=p+200
	}
	te.X=append(te.X,"256")
	te.X=append(te.X,"512")
	te.X=append(te.X,"1024")
	for i := 0; i < num; i++ {
		if(chosen[i]=="LCSS"){
			te.L=append(te.L,"SLCSS")
		}else if(chosen[i]=="DTW"){
			te.L=append(te.L,"SDTW")
		}else{
			te.L=append(te.L,"SBD")
		}
	}
	return te
}


var average Test
func Average(chosen [30]string,num int)Test {
	var j,k,p float64
	j=0
	k=0
	p=0
	for i := 0; i < 5; i++ {
		average.T1=append(average.T1,1+j)
		j=j+1
	}
	for i := 0; i < 5; i++ {
		average.T2=append(average.T2,1+k)
		k=k+2
	}
	for i := 0; i < 5; i++ {
		average.T3=append(average.T3,1+p)
		p=p+3
	}
	for i := 0; i < num; i++ {
		if(chosen[i]=="LCSS"){
			average.L=append(average.L,"SLCSS")
		}else if(chosen[i]=="DTW"){
			average.L=append(average.L,"SDTW")
		}else{
			average.L=append(average.L,"SBD")
		}
	}
	return average
}

var me Test
func Mermory(chosen [30]string,num int)Test{
	var j,k,p float64
	j=0
	k=0
	p=0
	for i := 0; i < 5; i++ {
		me.T1=append(me.T1,1+j)
		j=j+1
	}
	for i := 0; i < 5; i++ {
		me.T2=append(me.T2,1+k)
		k=k+2
	}
	for i := 0; i < 5; i++ {
		me.T3=append(me.T3,1+p)
		p=p+3
	}
	for i := 0; i < num; i++ {
		if(chosen[i]=="LCSS"){
			me.L=append(me.L,"SLCSS")
		}else if(chosen[i]=="DTW"){
			me.L=append(me.L,"SDTW")
		}else{
			me.L=append(me.L,"SBD")
		}
	}
	return me
}