package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 	请开始你的表演
	fmt.Println("开始运行")
	a := make(map[string]map[string]string) //二层map嵌套
	a["2"] = make(map[string]string)        //初始化第二层map
	a["2"]["value"] = "100"
	a["2"]["expire"] = "1111"

	a["1"] = make(map[string]string)
	a["1"]["value"] = "200"
	a["1"]["expire"] = "2222"
	c := cache{
		store: a,
	}
	//1.获取key对应的value
	http.HandleFunc("/api/cache/1", c.get)
	http.ListenAndServe(":8000", nil)

	//2.设置key
	http.HandleFunc("/api/cache/key", c.set)
	http.ListenAndServe(":8000", nil)

	//3.删除key
	http.HandleFunc("/api/cache/1", c.del)
	http.ListenAndServe(":8000", nil)

	//4.设置key过期时间
	http.HandleFunc("/api/cache/21/112220", c.setExpire)
	http.ListenAndServe(":8000", nil)
}

type Resp struct {
	Code int                          `json:"code"`
	Msg  string                       `json:"msg"`
	Data map[string]map[string]string `json:"data"`
}

type cache struct {
	store map[string]map[string]string
}

func (c *cache) get(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		key := req.URL.Path
		k := strings.Split(key, "/")[3]
		_, ok := c.store[k]
		nowTime := time.Now().Unix()
		var code = 0
		var msg = ""
		var data = map[string]map[string]string{}
		if ok == false {
			code = 404
			msg = "key不存在"
			resp := Resp{
				Code: code,
				Msg:  msg,
				Data: data,
			}
			bts, _ := json.Marshal(&resp)
			w.Write(bts)
		} else {
			keyTime := c.store[k]
			expire, _ := strconv.Atoi((keyTime["expire"]))
			if int(nowTime) > expire {
				code = 404
				msg = "key已过期"
			}
			code = 0
			msg = "获取成功"
			value := c.store[k]
			data[k] = value
			resp := Resp{
				Code: code,
				Msg:  msg,
				Data: data,
			}
			bts, _ := json.Marshal(&resp)
			w.Write(bts)
		}
	}

}

//设置缓存的key
func (c *cache) set(w http.ResponseWriter, req *http.Request) {
	respByte, _ := ioutil.ReadAll(req.Body)
	var m = map[string]string{}
	json.Unmarshal(respByte, &m)
	var va = map[string]string{
		"value":  m["value"],
		"expire": m["expire"],
	}
	c.store[m["key"]] = va
	r := Resp{
		Code: 0,
		Msg:  "设置成功",
		Data: c.store,
	}
	bts, _ := json.Marshal(&r)
	w.Write(bts)
}

//删除key
func (c *cache) del(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path
	k := strings.Split(key, "/")[3]
	delete(c.store, k)
	resp := Resp{
		Code: 0,
		Msg:  "删除成功",
	}
	bts, _ := json.Marshal(&resp)
	w.Write(bts)
}

//设置过期时间
func (c *cache) setExpire(w http.ResponseWriter, req *http.Request) {

	key := req.URL.Path
	k_url := strings.Split(key, "/")[3]
	respByte, _ := ioutil.ReadAll(req.Body)
	var request = map[string]string{}

	json.Unmarshal(respByte, &request)

	var code = 0
	var msg = ""
	_, ok := c.store[k_url]
	if ok == false {
		code = 1
		msg = "key不存在"
	} else {
		c.store[k_url]["expire"] = request["expire"]
		code = 0
		msg = "设置成功"
	}
	resp := Resp{
		Code: code,
		Msg:  msg,
		Data: c.store,
	}
	bts, _ := json.Marshal(&resp)
	w.Write(bts)
}
