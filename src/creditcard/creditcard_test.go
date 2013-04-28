package creditcard

import "testing"

func TestSetNumber(t *testing.T) {
  const out = "4111222233334444"
  c := &CreditCard{""}
  c.SetNumber(out)
  if x := c.Number; x != out {
    t.Errorf("SetNumber(%v) = %v, want %v", out, x, out)
  }
}

func TestSetNumberStripsSpaces(t *testing.T) {
  const in, out = "4111 2222 3333 4444", "4111222233334444"
  c := &CreditCard{""}
  if x, _ := c.SetNumber(in); x != out {
    t.Errorf("SetNumber(%v) = %v, want %v", in, x, out)
  }
}

func TestSetNumberStripsHyphens(t *testing.T) {
  const in, out = "4111-2222-3333-4444", "4111222233334444"
  c := &CreditCard{""}
  if x, _ := c.SetNumber(in); x != out {
    t.Errorf("SetNumber(%v) = %v, want %v", in, x, out)
  }
}

func TestTypeReturnsAmex1(t *testing.T) {
  const out = "AMEX"
  c := &CreditCard{"378282246310005"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsAmex2(t *testing.T) {
  const out = "AMEX"
  c := &CreditCard{"348282246310005"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsVisa(t *testing.T) {
  const out = "Visa"
  c := &CreditCard{"4111111111111111"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsVisaForShortNumber(t *testing.T) {
  const out = "Visa"
  c := &CreditCard{"4111111111111"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsMastercard1(t *testing.T) {
  const out = "Mastercard"
  c := &CreditCard{"5122222222222222"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsMastercard2(t *testing.T) {
  const out = "Mastercard"
  c := &CreditCard{"5222222222222222"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsMastercard3(t *testing.T) {
  const out = "Mastercard"
  c := &CreditCard{"5322222222222222"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsMastercard4(t *testing.T) {
  const out = "Mastercard"
  c := &CreditCard{"5422222222222222"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsMastercard5(t *testing.T) {
  const out = "Mastercard"
  c := &CreditCard{"5522222222222222"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsDiscover(t *testing.T) {
  const out = "Discover"
  c := &CreditCard{"6011111111111117"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}

func TestTypeReturnsUnknown(t *testing.T) {
  const out = "Unknown"
  c := &CreditCard{"0000000000000000"}
  if x := c.Type(); x != out {
    t.Errorf("Type() = %v, want %v", x, out)
  }
}
