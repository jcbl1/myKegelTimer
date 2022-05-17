package fileOpers

import (
	"bufio"
	"fmt"
	"os"
)

func AddXP(a int) {
	homePath := os.Getenv("HOME")
	err:=os.Mkdir(homePath+"/.myKegelTimer",0755)
	if err!=nil&&!os.IsExist(err){
		panic(err.Error())
	}
	f, err := os.OpenFile(homePath+"/.myKegelTimer/xp.dat", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err.Error())
	}

	rdr := bufio.NewReader(f)

	line, _, _ := rdr.ReadLine()

	fmt.Println(a)
	fmt.Println(string(line))
}
