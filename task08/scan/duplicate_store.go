package scan

type fileInfo struct {
	Filename string
	Size     int64
}

type duplicateStore struct {
	filepaths map[fileInfo][]string
}

func newDuplicateStore() *duplicateStore {
	return &duplicateStore{filepaths: make(map[fileInfo][]string)}
}
func (fps *duplicateStore) save(fileinfo fileInfo, filepath string) {
	fps.filepaths[fileinfo] = append(fps.filepaths[fileinfo], filepath)
}

func (fps *duplicateStore) getDuplicates() <-chan []string {
	ret := make(chan []string)
	go func() {
		for _, paths := range fps.filepaths {
			if len(paths) > 1 {
				ret <- paths
			}
		}
		close(ret)
	}()

	return ret
}
