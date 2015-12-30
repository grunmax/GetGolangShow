package getshow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

var errorlog *os.File
var Logger *log.Logger

func Err(userMessage string, e interface{}) {
	if e != nil {
		s := fmt.Sprintf("%s E:%v\n", userMessage, e)
		fmt.Printf(s)
		Logger.Panicf(s)
	}
}

func Log(userMessage string, v interface{}) {
	if v != nil {
		s := fmt.Sprintf("%s : %v\n", userMessage, v)
		fmt.Printf(s)
		Logger.Printf(s)
	}
}

func InitLog() {
	errorlog, e := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		fmt.Println(e)
	}
	Logger = log.New(errorlog, "app: ", log.Lshortfile|log.LstdFlags)
}

func IsFileExist(name string) bool {
	_, e := os.Stat(name)
	return e == nil
}

func GetUser() *user.User {
	user, e := user.Current()
	Err("Get user", e)
	return user
}

func WriteFile(fname string, fbody []byte) {
	e := ioutil.WriteFile(fname, fbody, 0644)
	Err("Write file", e)
}

func ReadFile(fname string) []byte {
	dat, e := ioutil.ReadFile(fname)
	Err("Read file", e)
	return dat
}

func DeleteFile(fname string) {
	e := os.Remove(fname)
	Err("Delete file", e)
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

func leftPadL(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
