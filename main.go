package main

import (
	"fmt"
	"github.com/qiniu/api.v7/kodo"
	"os"
	"path/filepath"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
	"strings"
)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func main() {
	conf.ACCESS_KEY = Config.Ak
	conf.SECRET_KEY = Config.Sk
	//创建一个Client
	client := kodo.New(0, nil)
	policy := &kodo.PutPolicy{
		Scope:      Config.Space,
		Expires:    3600,
		InsertOnly: 1,
	}

	token := client.MakeUptoken(policy)
	uploader := kodocli.NewUploader(0, nil)
	var ret PutRet
	filepath.Walk(Config.Path, func(filename string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			target := strings.Replace(filename, "\\", "/", -1)
			target = strings.Split(target, "/public/")[1]
			fmt.Println(target)
			fmt.Println(filename)
			res := uploader.PutFile(nil, &ret, token, target, filename, nil)
			if res != nil {
				fmt.Println("io.Put failed:", res)
			}
		}
		return nil
	})
}
