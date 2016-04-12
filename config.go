package main

//程序的配置项
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Qiniu struct {
	Ak    string `json:"ak"`
	Sk    string `json:"sk"`
	Space string `json:"space"`
	Path  string `json:"path"`
}

var Config Qiniu

func init() {
	txt, _ := ioutil.ReadFile("qiniu.config")
	err := json.Unmarshal(txt, &Config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
