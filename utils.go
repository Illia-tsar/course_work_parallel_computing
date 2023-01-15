package course_work_parallel_computing

import (
	"log"
	"os"
)

// GetFilenamesFromDir gets all filenames from the specified directory
// and concatenates the directory path with the filename
func GetFilenamesFromDir(startIdx int, endIdx int, dataDir string) []string {
	f, err := os.Open(dataDir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	files = files[startIdx:endIdx]
	output := make([]string, len(files))
	for i, v := range files {
		output[i] = dataDir + "/" + v.Name()
	}

	return output
}
