package compare

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gsync/TencentCos"
	"gsync/utils"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var (
	workerCount    = 1
	maxWorkerCount = 16
	searchRequest  = make(chan string)
	workerDone     = make(chan bool)
	FileMd5List    sync.Map
)

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go traversal(path, true)
		case <-workerDone:
			workerCount--
			if workerCount <= 0 {
				return
			}
		}
	}
}

func traversal(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		if master {
			workerDone <- true
		}
		return
	}
	for _, file := range files {
		name := file.Name()
		if !file.IsDir() {
			iv := 0
			h := md5.New()
			buf := make([]byte, 20480)
			r, _ := os.Open(path + name)
			for {
				n, err := r.Read(buf)
				iv += n
				if err == io.EOF {
					break
				}
				h.Write(buf[:n])
			}
			FileMd5List.Store(path+name, hex.EncodeToString(h.Sum(nil)))
			r.Close()
		}
		if file.IsDir() {
			if workerCount < maxWorkerCount {
				searchRequest <- path + name + "/"
			} else {
				traversal(path+name+"/", false)
			}
		}
	}
	if master {
		workerDone <- true
	}
}

func BuildFileMd5List(path string) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	go traversal(path, true)
	waitForWorkers()
}

func ShowFileMd5List(DstPath string) {
	BuildFileMd5List(TencentCos.PathConv(DstPath))
	FileMd5List.Range(func(key, value interface{}) bool {
		fmt.Printf("%s %s\n", utils.Blue(key), utils.Green(value))
		return true
	})
}
