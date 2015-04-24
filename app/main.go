package main

import (
	"fmt"
	"log"

	"github.com/juragan360/hdfs"
)

func main() {
	h, err := hdfs.NewHdfs(hdfs.NewHdfsConfig("http://localhost:50070", "mo.omer"))
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	es := h.MakeDirs([]string{"/user/mo.omer/test1", "/user/mo.omer/test2", "/user/mo.omer/test3"}, "0")

	if es != nil {
		for k, v := range es {
			fmt.Printf("Error when create %v : %v \n", k, v)
		}
	}
}
