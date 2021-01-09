package readir

import (
	"fmt"
	"io/ioutil"
)

func GetDir(path string) {
	files, _ := ioutil.ReadDir(path) //ここにディレクトリの場所を記載
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
