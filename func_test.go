package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL  = "http://192.168.1.2:8080/api/v1/Fibonacci/"

type ResponseBody struct {
	Msg string
	Result string
}

func testGetNormal(){

	testUrl := URL + "3"

	resp, err :=   http.Get(testUrl)
	if err != nil {
		fmt.Println(testUrl + " failed!")
		return
	}

	if resp.StatusCode != 200 {

		fmt.Println(testUrl + " failed! HTTP response is not 200 OK!")
		return

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(testUrl + " failed!")
		return
	}

	result := &ResponseBody{}

	json.Unmarshal(body, result)

	if result.Msg != "ok" || result.Result != "0,1,1" {

		fmt.Println(testUrl + " failed! Result is wrong!")
		return

	}

	fmt.Println(testUrl + " successed! ")
}

func testGetNonNumber(){

	testUrl := URL + "mess"

	resp, err :=   http.Get(testUrl)
	if err != nil {
		fmt.Println(testUrl + " failed!")
		return
	}

	if resp.StatusCode != 400 {

		fmt.Println(testUrl + " failed! HTTP response is not 400 Bad Request!")
		return

	} else {
		fmt.Println(testUrl + " successed! ")
	}

	defer resp.Body.Close()
}

func testGetNegativeNumber(){

	testUrl := URL + "-1"

	resp, err :=   http.Get(testUrl)
	if err != nil {
		fmt.Println(testUrl + " failed!")
		return
	}

	if resp.StatusCode != 400 {

		fmt.Println(testUrl + " failed! HTTP response is not 400 Bad Request!")
		return

	}


	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(testUrl + " failed!")
		return
	}

	result := &ResponseBody{}

	json.Unmarshal(body, result)


	if strings.Compare(result.Msg, "The input parameter n must NOT be negative! ") != 0 {

		fmt.Println(testUrl + " failed! Result is wrong!")
		return

	}

	fmt.Println(testUrl + " successed! ")
}

func main(){

	testGetNormal()
	testGetNonNumber()
	testGetNegativeNumber()

}