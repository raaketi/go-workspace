package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var files []string

const (
	fileToIgnore = ".DS_Store"
)

func main() {
	start := time.Now()
	root := "/Path/to/the/dir"
	log_pat := `([\d]{4}\-[\d]{2}\-[\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2})\s(\w+).(\w+)\s(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s(\w+)\s([^\s]+)\s([^\s]+)\s([^\s]+)\s-\s(([^\s]+))\s[^\s]+\s[^\s]+\s[^\s]+\s([^\s]+)\s([^\s]+)\s([^\s]+)`
	var reg_pattern = regexp.MustCompile(log_pat)
	wg := new(sync.WaitGroup)
	file_name := make(chan string)
	results := make(chan bool)
	go get_files(root, file_name, wg)
	for i := range file_name {
		wg.Add(1)
		go read_file(i, results, wg, reg_pattern)
  }
	t_file := <-results
	go func() {
		wg.Wait()
		if t_file {
			close(results)
		}
	}()
	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())

}

func get_files(root string, file_name chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), fileToIgnore) {
			file_name <- path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	close(file_name)
}

func read_file(target_file string, results chan<- bool, wg *sync.WaitGroup, log_pat *regexp.Regexp) {
	defer wg.Done()
	file, err := os.Open(target_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if log_pat.MatchString(scanner.Text()) {
			if !strings.Contains(scanner.Text(), "/ping/ping.html") && !strings.Contains(scanner.Text(), "/ping/monitor.html") {
				res := log_pat.FindAllStringSubmatch(scanner.Text(), -1)[0]
				fmt.Println(res[1] + "," + res[2] + "," + res[3] + "," + res[4] + "," + res[5] + "," + res[6] + "," + res[7] + "," + res[9] + "," + res[10] + "," + res[11] + "," + res[12] + "," + res[13] + "," + target_file)
			}
		}
	}
	results <- true
}
