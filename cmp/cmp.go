package cmp

func FindMoreItem(a, b []string) (more []string) {
	for i := 0; i < len(a); i++ {
		var succ bool = false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				succ = true
				break
			}
		}
		if !succ {
			more = append(more, a[i])
		}
	}
	return
}

func Compare(a, b []string) (more, less []string) {
	more = FindMoreItem(a, b)
	less = FindMoreItem(b, a)
	return
}

func DiffSliceString(a, b []string) bool {
	item := FindMoreItem(a, b)
	if len(item) != 0 {
		return true
	}
	item = FindMoreItem(b, a)
	if len(item) != 0 {
		return true
	}
	return false
}
