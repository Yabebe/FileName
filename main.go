package main

import (
	"fmt"
	"os"

	"./readir"
)

var pathToDir string = "/Users/toshiakiyabe/GoogleDriveFileStream/マイドライブ/01_Allhands_Crezit"

func main() {
	//対象ディレクトリの存在確認
	fp, err := os.Open(pathToDir)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close()

	//対象ディレクトリ直下の一覧取得
	readir.GetDir(pathToDir)

	NameChange("xxx", "yyy")

}

func NameChange(beforeName, afterName string) {
	if err := os.Rename(beforeName, afterName); err != nil {
		fmt.Println(err)
	}
}
