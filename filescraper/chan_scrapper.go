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

// type Scrapelog struct{

// }

func main() {
	start := time.Now()
	root := "Logfolder"
	log_pat := `([\d]{4}\-[\d]{2}\-[\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2})\s(\w+).(\w+)\s(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s(\w+)\s([^\s]+)\s([^\s]+)\s([^\s]+)\s-\s(([^\s]+))\s[^\s]+\s[^\s]+\s[^\s]+\s([^\s]+)\s([^\s]+)\s([^\s]+)`
	var reg_pattern = regexp.MustCompile(log_pat)
	wg := new(sync.WaitGroup)
	file_name := make(chan []string)
	wg.Add(1)
	go get_file_names(root, file_name, wg)
	// for _, i := range file_name {
	// 	wg.Add(1)
	// 	// go read_file(i, wg, reg_pattern)
	// 	go parse_file(i, wg, reg_pattern)
	// }

	// wg.Add(1)
	go parse_file(file_name, wg, reg_pattern)
	wg.Wait()
	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())

}

func get_file_names(root string, file_name chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()
	var file_ex_names []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), fileToIgnore) {
			file_ex_names = append(file_ex_names, path)
		}
		return nil
	})
	log.Printf("%d", len(file_ex_names))
	file_name <- file_ex_names
	if err != nil {
		panic(err)
	}
}

func parse_file(target_file chan []string, wg *sync.WaitGroup, log_pat *regexp.Regexp) {
	wg.Add(1)
	defer wg.Done()
	for _, v := range <-target_file {
		log.Printf("started file %s ", v)
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text(), v)
		}
	}

}

/* Commenting below code for testing in local
func read_file(target_file string, wg *sync.WaitGroup, log_pat *regexp.Regexp) {
	defer wg.Done()
	file, err := os.Open(target_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if log_pat.MatchString(scanner.Text()) {
			if !strings.Contains(scanner.Text(), "unwanted string in line") && !strings.Contains(scanner.Text(), "unwanted string in line") {
				res := log_pat.FindAllStringSubmatch(scanner.Text(), -1)[0]
				fmt.Println(res[1] + "," + res[2] + "," + res[3] + "," + res[4] + "," + res[5] + "," + res[6] + "," + res[7] + "," + res[9] + "," + res[10] + "," + res[11] + "," + res[12] + "," + res[13] + "," + target_file)
			}
		}
	}
	//results <- true
}*/
