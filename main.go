// main.go
package main

import "github.com/grunmax/GetGolangShow/getshow"

func init() {
	getshow.InitLog()
}

func main() {
	getshow.GetShowFiles(10, 20)
}
