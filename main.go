package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"time"
)

func main() {
	start := time.Now()
	//postMultiAsset()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed after 30 transactions", elapsed)
}

func postRequest() error {

	url := "http://172.25.62.160:8080/bs_platform/loan/assetUpload"

	//assetUID := "jyzb_jyzb_004_H201801018000" + strconv.Itoa(sequenceNo)
	//outTradeNo := "jyzb0220180117000" + strconv.Itoa(sequenceNo)
	//
	//post := "{\"assetUid\":" + "\"" + assetUID + "\"" +
	//	",\"timestamp\":\"123456\",\"bizContent\":{}," +
	//	"\"orgCode\":\"jyzb\",\"outTradeNo\":" + "\"" + outTradeNo + "\"" + "}"

	//post:= "{\"assetUid\":\"160815609421112015\",\"timestamp\":\"123456\",\"assetDetails\":\"160815609421112015,***n66_m,2016-08-15 20:41:30,2017-08-15 23:59:59,12,2399,199.92,0,24,0,3,3,1,1,HT201606300001\",\"orgCode\":\"jyzb\"}"

	post := "{\"assetUid\":\"160815609421112018\",\"timestamp\":\"123456\",\"bizContent\":{\"assetDetails\":\"160815609421112015,***n66_m,2016-08-15 20:41:30,2017-08-15 23:59:59,12,2399,199.92,0,24,0,3,3,1,1,HT201606300001\"},\"orgCode\":\"jyzb\",\"outTradeNo\":\"HT201606300001\"}"
	var jsonStr = []byte(post)
	//fmt.Println("jsonStr", jsonStr)
	//fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil

}

var bizContent = `{
        "assets":[
            {
                "assetUid":"160815609421112015",
                "assetDetails":"160815609421112015,***n66_m,2016-08-15 20:41:30,2017-08-15 23:59:59,12,2399,199.92,0,24,0,3,3,1,1,HT201606300001"
            }
        ]
    }`

var asset = `{        
                "assetUid":"160815609421112017",
                "assetDetails":"160815609421112017,***n66_m,2016-08-15 20:41:30,2017-08-15 23:59:59,12,2399,199.92,0,24,0,3,3,1,1,HT201606300001"
}
`

var (
	batchSize = 100
)

func makeBatch() (string, error) {
	var arr []interface{}

	js, _ := simplejson.NewJson([]byte(bizContent))
	nodes, _ := js.Map()
	p := nodes["assets"]
	d := p.([]interface{})
	for _, v := range d {
		for i := 0; i < batchSize; i++ {
			arr = append(arr, v)
		}
	}

	res := make(map[string]interface{})
	res["assets"] = arr
	c, _ := json.Marshal(res)
	//fmt.Println(string(c))
	res["bizContent"] = string(c)
	res["assetUid"] = "160815609421112015"
	res["timestamp"] = "20180123155136"
	res["outTradeNo"] = "160815609421112015"
	res["orgCode"] = "jyzb"
	delete(res, "assets")
	request, err := json.Marshal(res)
	return string(request), err
}

func postMultiAsset(request string ) error {

	url := "http://172.25.62.160:8080/bs_platform/loan/assetUpload"
	var jsonStr = []byte(request)
	//fmt.Println("jsonStr", jsonStr)
	//fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
