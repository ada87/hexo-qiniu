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
	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	//设置上传文件的路径
	filepath.Walk(Config.Path, func(filename string, fi os.FileInfo, err error) error {
		//		fmt.Println(filename)
		//		fmt.Println(fi.Name())
		//		fmt.Println(fi.IsDir())
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

	//	filepath := "G:\\install-unbuntu\\apr-1.5.1.tar.gz"
	//	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名

	//	res := uploader.PutFile(nil, &ret, token, "apr-1.5.1.tar.gz", filepath, nil)
	//	//	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)

	//	//打印返回的信息
	//	fmt.Println(ret)
	//打印出错信息
	//	if res != nil {
	//		fmt.Println("io.Put failed:", res)
	//		return
	//	}

}
