package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// var files []string

const (
	fileToIgnore = ".DS_Store"
)

// type Scrapelog struct{

// }

func init() {
	runtime.GOMAXPROCS(8)
}

func main() {
	start := time.Now()
	// log_pat := `([\d]{4}\-[\d]{2}\-[\d]{2}\s[\d]{2}:[\d]{2}:[\d]{2})\s(\w+).(\w+)\s(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s(\w+)\s([^\s]+)\s([^\s]+)\s([^\s]+)\s-\s(([^\s]+))\s[^\s]+\s[^\s]+\s[^\s]+\s([^\s]+)\s([^\s]+)\s([^\s]+)`
	// var reg_pattern = regexp.MustCompile(log_pat)
	root := "Logfolder"
	wg := new(sync.WaitGroup)
	file_name := make(chan string)
	workers := 24
	// wg.Add(workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go scrape_log_file(i, file_name, wg)
	}
	wg.Add(1)
	go func() {
		get_file_names_root(root, file_name, wg)
		defer wg.Done()
		defer close(file_name)
	}()
	wg.Wait()
	elapsedTime := time.Since(start)
	log.Printf("Total Time For Execution: %s" + elapsedTime.String())

}

func get_file_names_root(root string, file_name chan string, wg *sync.WaitGroup) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), fileToIgnore) {
			file_name <- path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func scrape_log_file(worker int, target_file chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range target_file {
		log.Printf("[Worker: %d] Started work.", worker)
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text(), v)
		}
		log.Printf("[Worker: %d] Finished file scan: %s", worker, v)
	}

}
