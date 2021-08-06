package scan

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var numWorker = runtime.NumCPU()

type fileInfoWithPath struct {
	fileInfo
	filepath string
}

func logFatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func scanDir(wg *sync.WaitGroup, dirs <-chan string, files chan<- fileInfoWithPath) {
	for dir := range dirs {
		entries, err := os.ReadDir(dir)
		logFatalOnError(err)

		for _, entry := range entries {
			if !entry.IsDir() {
				info, err := entry.Info()
				logFatalOnError(err)
				files <- fileInfoWithPath{
					fileInfo{Filename: entry.Name(), Size: info.Size()},
					filepath.Join(dir, entry.Name()),
				}
			}
		}
	}
	wg.Done()
}

func saveFileInfo(ds *duplicateStore, files <-chan fileInfoWithPath, done chan<- struct{}) {
	for file := range files {
		ds.save(file.fileInfo, file.filepath)
	}
	done <- struct{}{}
}

func ScanDir(root string) <-chan []string {
	wg := &sync.WaitGroup{}
	dirs := make(chan string, 2*numWorker)
	files := make(chan fileInfoWithPath, 5*numWorker)
	done := make(chan struct{})
	ds := newDuplicateStore()

	go saveFileInfo(ds, files, done)

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go scanDir(wg, dirs, files)
	}

	queue := make([]string, 1)
	queue[0] = root
	dirs <- root

	for len(queue) > 0 {
		dirpath := queue[0]
		queue = queue[1:]

		entries, err := os.ReadDir(dirpath)
		logFatalOnError(err)

		for _, entry := range entries {
			if entry.IsDir() {
				fi, err := entry.Info()
				logFatalOnError(err)

				newDirpath := filepath.Join(dirpath, fi.Name())

				queue = append(queue, newDirpath)
				dirs <- newDirpath
			}
		}
	}
	close(dirs)
	wg.Wait()
	close(files)
	<-done

	return ds.getDuplicates()
}
