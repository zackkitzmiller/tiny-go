package tiny

import "testing"

func getTiny() *tiny {
	return NewTiny("5SX0TEjkR1mLOw8Gvq2VyJxIFhgCAYidrclDWaM3so9bfzZpuUenKtP74QNH6B")
}

func TestToTiny(t *testing.T) {
	tiny := getTiny()
	converted := tiny.To(5)
	if converted != "E" {
		t.Errorf("expected %q, got %q", "E", converted)
	}
}

func TestFromTiny(t *testing.T) {
	tiny := getTiny()
	reversed := tiny.From("E")
	if reversed != 5 {
		t.Errorf("expected %q, got %q", 5, reversed)
	}
}

func TestMant(t *testing.T) {
	tiny := getTiny()
	for i := 0; i <= 500; i++ {
		tinied := tiny.From(tiny.To(i))
		if tinied != i {
			t.Errorf("expected %q, got %q", i, tinied)
		}
	}
}
