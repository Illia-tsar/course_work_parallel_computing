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

// GetFilenames is a wrapper around GetFilenamesFromDir. It extracts the
// filepaths from the specified directories with respect to calculated indexes
func GetFilenames() []string {
	dirs := [5]string{"data/test/neg", "data/test/pos", "data/train/neg", "data/train/pos", "data/train/unsup"}
	bStartIdx := 9000
	bEndIdx := 9250
	eStartIdx := 36000
	eEndIdx := 37000

	var outArr []string
	var temp []string
	for i, v := range dirs {
		if i != 4 {
			temp = GetFilenamesFromDir(bStartIdx, bEndIdx, v)
		} else {
			temp = GetFilenamesFromDir(eStartIdx, eEndIdx, v)
		}
		outArr = append(outArr, temp...)
	}
	return outArr
}
