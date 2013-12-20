package main

import _ "fmt"
import "os"
import _ "io/ioutil"
import "bufio"
import "image"

func main() {
	fi, _ := os.Open("xwing.jpg")
	//defer fi.Close()
	r := bufio.NewReader(fi)
	i := image.Decode(r)
}
