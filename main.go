package main

import (
	"os"
	"github.com/zhuqiuzhi/go-crwaler/links"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"flag"
)

func main() {
	// breadthFirth call crwal for each item in the worklist
	// os.Args[1:] is the initial worklist
	var breadth uint
	flag.UintVar(&breadth, "b", 3,  "广度优先搜索的广度")

	breadthFirth(crwal, os.Args[1:], breadth)
}

// breadthFirth encapsulates the essence of a bradth-first traversal
func breadthFirth(f func(string) []string, worklist []string, breadth uint) {
	seen := make(map[string]bool)

	// invaraiant: item in worklist are to be visited
	for i:= uint(0); len(worklist) > 0 && i < breadth ; i++ {
		items := worklist
		worklist = nil
		for _ ,item := range items{
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}

}

func crwal(url string) []string{
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}