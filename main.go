package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/qiniu/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
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
	uploader := kodocli.NewUploader(0, nil)
	fast := true //true：快速模式：仅上传HTML，否则全部上传
	fmt.Println(os.Args)
	if len(os.Args) >= 2 {
		fast = false
	}
	fmt.Println("使用快速模式：", fast)
	var ret PutRet
	filepath.Walk(Config.Path, func(filename string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			if fast && !strings.HasSuffix(fi.Name(), ".html") {
				return nil
			}
			target := strings.Replace(filename, "\\", "/", -1)
			target = strings.Split(target, "/public/")[1]
			policy := &kodo.PutPolicy{
				Scope: Config.Space + ":" + target,
			}
			token := client.MakeUptoken(policy)
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
