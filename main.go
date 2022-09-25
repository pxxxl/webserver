package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var files []string

	f, err1 := os.Open("/page/index")
	if err1 != nil {
		fmt.Println("read file fail", err1)
		return
	}
	defer f.Close()
	fd, _ := ioutil.ReadAll(f)
	webpage := string(fd)

	root := "/poems/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		f, err := os.Open(root)
		if err != nil {
			fmt.Println("read file fail", err)
			return err
		}
		defer f.Close()
		fd, err := ioutil.ReadAll(f)
		content := string(fd)
		strings.Replace(webpage, "<!--replacement -->", content, 1)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Fprintf(w, webpage)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
