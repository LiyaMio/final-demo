package send

//import (
//	"bufio"
//	"fmt"
//	"io"
//	"os"
//	"strconv"
//	"strings"
//)
//type Datanew struct{
//	X float64 `json:"latitude"`
//	Y float64 `json:"longitude"`
//}
//var datass[] Datanew=make([]Datanew,0,10000)
//func Readnew(key int64) []Datanew{
//		if key>=1 && key<=100 {
//			c := fmt.Sprintf("%s%05d%s", "./resource/normal/normal_", key, ".csv")
//			f, err := os.Open(c)
//			if err != nil {
//				fmt.Println(err)
//			}
//			var a, b float64
//			defer f.Close()
//			buff := bufio.NewReader(f)
//			for {
//				line, err := buff.ReadString('\n')
//				if err == io.EOF {
//					break
//				}
//				tmp := strings.Split(line, ",")
//				a, _ = strconv.ParseFloat(tmp[0], 64)
//				//string([]byte(tmp[1])[:len(tmp[1]) - 2])
//				b, _ = strconv.ParseFloat(strings.Split(tmp[1], "\n")[0], 64)
//				var point Datanew
//				point.X=a
//				point.Y=b
//				datass=append(datass,point)
//			}
//		}else if key>=101&&key<=200{
//			c := fmt.Sprintf("%s%05d%s", "./resource/SHH-Taxi/SHH-Taxi_", key-100, ".csv")
//			f, err := os.Open(c)
//			if err != nil {
//				fmt.Println(err)
//			}
//			var a, b float64
//			defer f.Close()
//			buff := bufio.NewReader(f)
//			for {
//				line, err := buff.ReadString('\n')
//				if err == io.EOF {
//					break
//				}
//				tmp := strings.Split(line, ",")
//				a, _ = strconv.ParseFloat(tmp[0], 64)
//				//string([]byte(tmp[1])[:len(tmp[1]) - 2])
//				b, _ = strconv.ParseFloat(strings.Split(tmp[1], "\n")[0], 64)
//				var point Datanew
//				point.X=a
//				point.Y=b
//				datass=append(datass,point)
//			}
//	}else if key>=201&&key<=300 {
//			c := fmt.Sprintf("%s%05d%s", "./resource/T-drive/T-drive_", key-200, ".csv")
//			f, err := os.Open(c)
//			if err != nil {
//				fmt.Println(err)
//			}
//			var a, b float64
//			defer f.Close()
//			buff := bufio.NewReader(f)
//			for {
//				line, err := buff.ReadString('\n')
//				if err == io.EOF {
//					break
//				}
//				tmp := strings.Split(line, ",")
//				a, _ = strconv.ParseFloat(tmp[0], 64)
//				//string([]byte(tmp[1])[:len(tmp[1]) - 2])
//				b, _ = strconv.ParseFloat(strings.Split(tmp[1], "\n")[0], 64)
//				var point Datanew
//				point.X = a
//				point.Y = b
//				datass = append(datass, point)
//			}
//		}else if key>=301&&key<=400 {
//			c := fmt.Sprintf("%s%05d%s", "./resource/uniform/uniform_", key-300, ".csv")
//			f, err := os.Open(c)
//			if err != nil {
//				fmt.Println(err)
//			}
//			var a, b float64
//			defer f.Close()
//			buff := bufio.NewReader(f)
//			for {
//				line, err := buff.ReadString('\n')
//				if err == io.EOF {
//					break
//				}
//				tmp := strings.Split(line, ",")
//				a, _ = strconv.ParseFloat(tmp[0], 64)
//				//string([]byte(tmp[1])[:len(tmp[1]) - 2])
//				b, _ = strconv.ParseFloat(strings.Split(tmp[1], "\n")[0], 64)
//				var point Datanew
//				point.X = a
//				point.Y = b
//				datass = append(datass, point)
//			}
//
//		}
//		//fmt.Println()
//		//for _, v := range datass.X {
//		//	fmt.Println(v)
//		//}
//		//for _, v := range datass.Y {
//		//	fmt.Println(v)
//		//}
//	//fmt.Println(datass)
//	return datass
//}

//var datass[] Datanew=make([]Datanew,0,1000)
//func Readnew() []Datanew{
//	for i := 1; i <= 1; i++ {
//		c := fmt.Sprintf("%s%05d%s", "./normal./normal_", i, ".csv")
//		fmt.Println(c)
//		f, err := os.Open(c)
//		if err != nil {
//			fmt.Println(err)
//		}
//		var a, b float64
//		defer f.Close()
//		buff := bufio.NewReader(f)
//		for {
//			line, err := buff.ReadString('\n')
//			if err == io.EOF {
//				break
//			}
//			tmp := strings.Split(line, ",")
//			a, _ = strconv.ParseFloat(tmp[0], 64)
//			//string([]byte(tmp[1])[:len(tmp[1]) - 2])
//			b, _ = strconv.ParseFloat(strings.Split(tmp[1], "\r\n")[0], 64)
//			var point Datanew
//			point.X=a
//			point.Y=b
//			datass=append(datass,point)
//
//		}
//		//fmt.Println()
//		//for _, v := range datas.X1 {
//		//	fmt.Println(v)
//		//}
//		//for _, v := range datas.Y1 {
//		//	fmt.Println(v)
//		//}
//	}
//	//fmt.Println(datass)
//	return datass
//}