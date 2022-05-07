package fileOpers

import (
	"fmt"
	"os"
	"bufio"
)

func AddXP(a int) {
	homePath := os.Getenv("HOME")
	f, err := os.Open(homePath + "/.myKegelTimer/xp.dat")
	if err != nil {
		panic(err.Error())
	}

	rdr:=bufio.NewReader(f)

	line,_,_:=rdr.ReadLine()

	fmt.Println(a)
	fmt.Println(string(line))
}
