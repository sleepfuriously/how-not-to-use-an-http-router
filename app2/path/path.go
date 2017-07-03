package path

type Iterator struct {
	segments []string
}

func NewIterator(path string) *Iterator {
	const slash = '/'
	var (
		segments  []string
		inSegment = false
		start     = 0
	)
	for i := range path {
		switch {
		case inSegment && path[i] == slash:
			inSegment = false
			segments = append(segments, path[start:i])
		case !inSegment && path[i] != slash:
			inSegment = true
			start = i
		}
	}
	// Unterminated segment at end of path.
	if inSegment {
		segments = append(segments, path[start:])
	}
	return &Iterator{segments}
}

func (iter *Iterator) Next() (string, bool) {
	if len(iter.segments) == 0 {
		return "", false
	}
	head := iter.segments[0]
	iter.segments = iter.segments[1:]
	return head, true
}
