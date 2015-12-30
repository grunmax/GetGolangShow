package getshow

import (
	"fmt"
	"strconv"
	"sync"
)

const URL_ = "https://golangshow.com/cdn/episodes/%s.mp3"
const MAXSHOW = 35
const FILEMP3_ = "%s/Music/GolangShow/%s.mp3"

func GetShowFiles(fromNumber, toNumber int) {
	if (fromNumber > toNumber) || (toNumber > MAXSHOW) || (fromNumber < 1) {
		Err("wrong numbers %v", []int{fromNumber, toNumber})
	}

	messagesCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(toNumber - fromNumber + 1)

	for i := fromNumber; i <= toNumber; i++ {
		go func(number int) {
			defer wg.Done()
			Log("Start", number)
			//time.Sleep(time.Second * 1) //debug
			getShowFile(number)
			messagesCh <- number
		}(i)
	}

	go func() {
		for i := range messagesCh {
			Log("Ready", i)
		}
	}()

	wg.Wait()
}

func getShowFile(showNumber int) {

	dataFunc := func(n int) (string, string) {
		ns := strconv.Itoa(n)
		nsf := leftPadL(ns, "0", 3)
		url := fmt.Sprintf(URL_, nsf)
		user := GetUser()
		homeDir := user.HomeDir
		fname := fmt.Sprintf(FILEMP3_, homeDir, nsf)
		return url, fname
	}

	//go
	url, fname := dataFunc(showNumber)
	if IsFileExist(fname) {
		Log("Skipped", fname)
		return
	}
	content := GetHttpBytes(url)
	WriteFile(fname, content)
}
