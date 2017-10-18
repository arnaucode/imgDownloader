package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//http://www.image-net.org/search?q=hotdog
	savelog()
	//create directory webOutput
	_ = os.Mkdir("imgs", os.ModePerm)
	readLinks("imagesLinks.txt")

}
func readLinks(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var i int
	for scanner.Scan() {
		log.Println(scanner.Text())
		saveLinkToImg(scanner.Text(), i)
		//log.Println(strconv.Itoa(i) + "/" + strconv.Itoa(nLines))
		log.Println(i)
		i++
	}
}

func countLines(scanner *bufio.Scanner) int {
	var count int
	for scanner.Scan() {
		count++
		fmt.Println(count)
	}
	return count
}

func saveLinkToImg(url string, i int) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("imgs/" + strconv.Itoa(i) + ".png")
	check(err)
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	check(err)
	file.Close()
	fmt.Println("Success!")
}
