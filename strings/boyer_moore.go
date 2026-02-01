package strings

const ALPHABET_SIZE = 256

type BoyerMoore struct {
	badChar [ALPHABET_SIZE]int
}

func NewBoyerMoore() *BoyerMoore {
	bm := &BoyerMoore{}
	for i := 0; i < ALPHABET_SIZE; i++ {
		bm.badChar[i] = -1
	}
	return bm
}

func (bm *BoyerMoore) badCharHeuristic(pattern string) {
	m := len(pattern)
	for i := 0; i < m-1; i++ {
		bm.badChar[int(pattern[i])] = m - 1 - i
	}
}

func (bm *BoyerMoore) Search(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	var result []int

	if m == 0 {
		return result
	}

	bm.badCharHeuristic(pattern)

	s := 0
	for s <= n-m {
		j := m - 1

		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			result = append(result, s)
			if s+m < n {
				shift := bm.badChar[int(text[s+m])]
				if shift < 1 {
					shift = 1
				}
				s += shift
			} else {
				s += 1
			}
		} else {
			shift := bm.badChar[int(text[s+j])]
			if shift < 1 {
				shift = 1
			}
			s += shift
		}
	}

	return result
}

func (bm *BoyerMoore) SearchFirst(text, pattern string) int {
	indices := bm.Search(text, pattern)
	if len(indices) == 0 {
		return -1
	}
	return indices[0]
}

func (bm *BoyerMoore) Count(text, pattern string) int {
	return len(bm.Search(text, pattern))
}

func (bm *BoyerMoore) Contains(text, pattern string) bool {
	return bm.SearchFirst(text, pattern) != -1
}
