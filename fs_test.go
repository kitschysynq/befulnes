package befulnes

import "testing"

func TestFile(t *testing.T) {
	w := &WordList{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
	}

	err := w.ToFile("thor.words")
	if err != nil {
		t.Errorf("failed writing to file: %q", err.Error())
	}

	x := &WordList{}
	err = x.FromFile("thor.words")
	if err != nil {
		t.Errorf("failed reading from file: %q", err.Error())
	}
	t.Logf("x.Length() == %d", x.Length())
	for i, v := range *x {
		t.Logf("  %d: %q", i, v)
	}

	if !equal(w, x) {
		t.Errorf("files do not match")
	}
}
