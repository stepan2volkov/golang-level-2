package scan

import (
	"testing"
)

func TestFilePathStore(t *testing.T) {
	duplicatedFiles := []struct {
		fileinfo fileInfo
		filepath string
	}{
		{fileInfo{"1.txt", 10}, "/test/1.txt"},
		{fileInfo{"1.txt", 10}, "/test1/1.txt"},
		{fileInfo{"1.txt", 10}, "/test2/1.txt"},
		{fileInfo{"1.txt", 10}, "/test3/1.txt"},
	}

	notDuplicatedFiles := []struct {
		fileinfo fileInfo
		filepath string
	}{
		{fileInfo{"1.txt", 11}, "/test4/1.txt"},
		{fileInfo{"2.txt", 15}, "/test/2.txt"},
		{fileInfo{"3.txt", 10}, "/test/3.txt"},
		{fileInfo{"4.txt", 10}, "/test/4.txt"},
	}

	store := newDuplicateStore()
	for _, file := range duplicatedFiles {
		store.save(file.fileinfo, file.filepath)
	}

	for _, file := range notDuplicatedFiles {
		store.save(file.fileinfo, file.filepath)
	}

	duplicates := store.getDuplicates()
	for duplicate := range duplicates {
		if len(duplicate) != len(duplicatedFiles) {
			t.Fatalf("Expected %d duplicated, got %d", len(duplicatedFiles), len(duplicate))
		}

		for i, got := range duplicate {
			want := duplicatedFiles[i].filepath
			if want != got {
				t.Errorf("Expected '%s', got '%s'", want, got)
			}
		}
	}
}
