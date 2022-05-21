package main

import (
	"fmt"
	//"github.com/jcbl1/myKegelTimer/fileOpers"
	"math/rand"
	"time"
)

func main() {
	//degug
	//	fileOpers.AddXP(500)

	rand.Seed(time.Now().Unix())

	fiftyms, err := time.ParseDuration("50ms")
	if err != nil {
		panic(err.Error())
	}

	/*
		onesec, err := time.ParseDuration("1s")
		if err != nil {
			panic(err.Error())
		}
	*/

	fmt.Printf("预备~")
	time.Sleep(time.Second)
	fmt.Printf("3,")
	time.Sleep(time.Second)
	fmt.Printf("2,")
	time.Sleep(time.Second)
	fmt.Printf("1")
	time.Sleep(time.Second)
	fmt.Println(" 开始！")
	holdStat := true
	for i := 0; i < 20*175; i++ {
		if i%100 == 0 && i != 0 {
			/*
				fmt.Printf("\r")
				if i%600 == 0 {
					fmt.Printf("还剩%d秒！", (20*180-i)/20)
				}*/
			switch holdStat {
			case true:
				/*
					fmt.Print("放！")
					time.Sleep(fiftyms)
				*/
				holdStat = false
				//addOil()
				//goto label
			case false:
				//fmt.Print("收！")
				//time.Sleep(fiftyms)
				holdStat = true
				//addOil()
				//goto label
			}
		}
		switch holdStat {
		case true:
			fmt.Printf("\r[%d] 收！", (20*180-i)/20)
		case false:
			fmt.Printf("\r[%d] 放！", (20*180-i)/20)
		}
		if i > 20*100 {
			fmt.Print("加油，胜利就在前方！")
		}

		time.Sleep(fiftyms)
		//label:
	}
	time.Sleep(time.Second)
	fmt.Printf("\r恭喜你完成训练！\n")
	time.Sleep(time.Second)
	fmt.Printf("在本次训练中你一共获得了%v点经验！再接再厉吧！\n", rand.Int()%400000+2345234)
	time.Sleep(time.Second)

	//	time.Sleep(dur)
	//	fmt.Println("hh")

}

/*
func addOil() {
	if rand.Int()%10 == 5 {
		fmt.Print("加油！胜利就在前方！")
	}
}
*/
