package luhn

import "testing"

func TestValidWithEmptyString(t *testing.T) {
  const in, out = "", false
  if x := Valid(in); x != out {
    t.Errorf("Valid(%v) = %v, want %v", in, x, out)
  }
}

func TestValidWithCorrectNumber(t *testing.T) {
  const in, out = "4111111111111111", true
  if x := Valid(in); x != out {
    t.Errorf("Valid(%v) = %v, want %v", in, x, out)
  }
}

func TestValidWithIncorrectNumber(t *testing.T) {
  const in, out = "4111111111111110", false
  if x := Valid(in); x != out {
    t.Errorf("Valid(%v) = %v, want %v", in, x, out)
  }
}

func TestdigitalRoot(t *testing.T) {
  const in, out = 25, 7
  if x := digitalRoot(in); x != out {
    t.Errorf("digitalRoot(%v) = %v, want %v", in, x, out)
  }
}