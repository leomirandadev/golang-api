package filesController

import (
	"sync"
	"vcfConverter/models/files"
)

var wg sync.WaitGroup

func DeleteSome(idFiles []int64) bool {
	for _, idFile := range idFiles {
		wg.Add(1)
		go delete(idFile)
	}

	wg.Wait()

	return true
}

func delete(idFile int64) {
	defer wg.Done()
	files.Delete(idFile)
}
