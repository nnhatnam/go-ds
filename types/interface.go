package types

// A type that satisfies types.Comparator can be compared with other types.Comparator using the Compare method
type Comparator interface {
	//call Cmp if we want to compare two Comparators a and b. Possible results
	// -1, if a < b
	// 0, if a == b
	// 1, if a > b
	Cmp(Comparator) int
}



