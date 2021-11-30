package scan

import (
	"go.uber.org/zap"
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

func logFatalOnError(logger *zap.Logger, err error) {
	if err != nil {
		logger.Fatal("fatal error", zap.Error(err))
	}
}

func scanDir(logger *zap.Logger, wg *sync.WaitGroup, dirs <-chan string, files chan<- fileInfoWithPath) {
	for dir := range dirs {
		logger.Info("reading dir", zap.String("dir", dir))
		// panic("test")
		entries, err := os.ReadDir(dir)
		logFatalOnError(logger, err)

		for _, entry := range entries {
			if !entry.IsDir() {
				logger.Info("getting info about file",
					zap.String("filename", entry.Name()),
					zap.String("filepath", filepath.Join(dir, entry.Name())))
				info, err := entry.Info()
				logFatalOnError(logger, err)
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

func ScanDir(logger *zap.Logger, root string) <-chan []string {
	wg := &sync.WaitGroup{}
	dirs := make(chan string, 2*numWorker)
	files := make(chan fileInfoWithPath, 5*numWorker)
	done := make(chan struct{})
	ds := newDuplicateStore()

	go saveFileInfo(ds, files, done)

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go scanDir(logger, wg, dirs, files)
	}

	dirQueue := make([]string, 1)
	dirQueue[0] = root
	dirs <- root

	for len(dirQueue) > 0 {
		dirpath := dirQueue[0]
		dirQueue = dirQueue[1:]

		entries, err := os.ReadDir(dirpath)
		logFatalOnError(logger, err)

		for _, entry := range entries {
			if entry.IsDir() {
				fi, err := entry.Info()
				logFatalOnError(logger, err)

				newDirpath := filepath.Join(dirpath, fi.Name())

				dirQueue = append(dirQueue, newDirpath)
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
