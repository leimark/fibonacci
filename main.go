package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {

	Port string `json: "port"`
	Logfile string `json: "logfile"`
	Loglevel string `json: "loglevel"'`
}


func getFibonacci (c *gin.Context){

	n, err := strconv.ParseInt(c.Param("n"), 10, 64)

	//If n cannot be convert to a number, then prompt the user to input correct parameter
	if err != nil {
		fmt.Print(err)
		//consider defer?
		c.JSON(http.StatusBadRequest, gin.H{"result": "", "msg": "The input parameter n must be an integer!"})
		//panic? log error.
	}

	//Check if n is negative, then prompt the user to input correct parameter
	if n < 0 {

		c.JSON(http.StatusBadRequest, gin.H{"result": "", "msg": "The input parameter n must NOT be negative! "})

	}

	fmt.Println(n)

	var fiboarray = calculateFibonacci(n)

	var buffer bytes.Buffer


	for _, value := range fiboarray {

		buffer.WriteString(strconv.FormatUint(value, 10))
		buffer.WriteString(",")
	}

	str := strings.TrimRight(buffer.String(),",")

	c.JSON(http.StatusOK, gin.H{"result": str, "msg": "ok"})
}


func calculateFibonacci (n int64) [] uint64 {

	resultarray := make([]uint64, n)

	if n == 1 {

		resultarray[0] = 0

	} else {

		var i int64

		resultarray[0], resultarray[1] = 0, 1

		for i = 2; i < n; i++ {

			resultarray[i] =  resultarray[i-1] + resultarray[i-2]
		}
	}

	return resultarray

}


func LoadConfig() *Config {

	//Get the full path of the configuration file
	path := filepath.Dir(os.Args[0])

	configPath := path + string(os.PathSeparator) + "config.json"

	confFile, err := os.Open(configPath)

	if err != nil {
		panic(err)
	}

	defer confFile.Close()

	fContent, err := ioutil.ReadAll(confFile)

	confContent := string(fContent)

	confObj := &Config{}

	err = json.Unmarshal([]byte(confContent), &confObj)

	if err != nil {
		panic(err)
	}

	return confObj
}


func main() {

	fmt.Println("Starting Fabonacci process ....")

	confObj := LoadConfig()

	//Create log file
	fLog, _ := os.Create(confObj.Logfile)

	defer fLog.Close()

	gin.DefaultWriter = io.MultiWriter(fLog)

	r := gin.Default()

	r.GET("/api/v1/Fibonacci/:n",getFibonacci ) //Register the get Fibonacci API

	r.Run(":"+confObj.Port)


}
