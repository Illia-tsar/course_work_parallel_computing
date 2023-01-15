package course_work_parallel_computing

// InvertedIndexItem contains the term, its frequency
// and an array of documents it's present in
type InvertedIndexItem struct {
	Term         string
	Frequency    int
	DocumentList []int
}

// InvertedIndex contains a hash map to check the presence
// of a word and an array of InvertedIndexItem
type InvertedIndex struct {
	HashMap map[string]*InvertedIndexItem
	Items   []*InvertedIndexItem
}
