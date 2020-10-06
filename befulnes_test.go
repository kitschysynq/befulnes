package befulnes

import "testing"

func TestShuffle(t *testing.T) {
	w := &WordList{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
	}
	sorted := &WordList{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
	}

	total := 100.0
	match := 0
	for i := 0; i < 100; i++ {
		w.Shuffle()
		if w.Length() != sorted.Length() {
			t.Errorf("length changed after shuffle")
		}
		if equal(w, sorted) {
			match++
		}
	}

	if (float64(match) / total) > 0.2 {
		t.Errorf("more than 20%% of shuffled lists matched the original")
	}
}

func TestConsume(t *testing.T) {
	w := &WordList{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
	}

	l := w.Length()

	for i := 0; i < l; i++ {
		_ = w.Consume()
		if w.Length() >= l-i {
			t.Errorf("wordlist too long after consuming")
		}
	}
}

func equal(a, b *WordList) bool {
	if a.Length() != b.Length() {
		return false
	}
	for i := range *a {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}
