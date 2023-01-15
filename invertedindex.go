package course_work_parallel_computing

import "sync"

var itemsMutex sync.Mutex
var mapMutex sync.Mutex

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

// FindItem returns the index of a given
// Item in an array of InvertedIndexItem
func (invertedIndex *InvertedIndex) FindItem(Term string) int {
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index
		}
	}
	panic("Not Found")
}

// AddItem first checks the presence of a word in the hash map. If the
// answer is positive it increases the frequency and adds the current
// document to DocumentList. In other case it adds the word to the hash map and items list
func (invertedIndex *InvertedIndex) AddItem(Term string, Document int) {
	mapMutex.Lock()
	temp := invertedIndex.HashMap[Term]
	mapMutex.Unlock()
	if temp != nil {
		itemsMutex.Lock()
		FoundItemIdx := invertedIndex.FindItem(Term)
		itemsMutex.Unlock()

		itemsMutex.Lock()
		invertedIndex.Items[FoundItemIdx].Frequency++
		invertedIndex.Items[FoundItemIdx].DocumentList = append(invertedIndex.Items[FoundItemIdx].DocumentList, Document)
		itemsMutex.Unlock()
	} else {

		InvertedIndexItem := &InvertedIndexItem{
			Term:         Term,
			Frequency:    1,
			DocumentList: []int{Document},
		}

		mapMutex.Lock()
		itemsMutex.Lock()
		invertedIndex.HashMap[Term] = InvertedIndexItem
		invertedIndex.Items = append(invertedIndex.Items, InvertedIndexItem)
		mapMutex.Unlock()
		itemsMutex.Unlock()
	}
}
