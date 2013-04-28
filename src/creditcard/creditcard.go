package creditcard

import (
  "luhn"
  "regexp"
)

var types = map[string] string {
  "AMEX": "^3[47]\\d{13}$",
  "Visa": "^4(\\d{12}|\\d{15})$",
  "Mastercard": "^5[1-5]\\d{14}$",
  "Discover": "^6011\\d{12}",
}

type CreditCard struct {
  Number string
}

func NewCreditCard(num string) (c *CreditCard) {
  c = new(CreditCard)
  c.SetNumber(num)
  return
}

func (c *CreditCard) String() string {
  return c.Number
}

func (c *CreditCard) SetNumber(num string) (string, error) {
  repl, err := regexp.Compile("\\D")

  if err == nil {
    cleanNum := repl.ReplaceAllString(num, "")
    c.Number = cleanNum
  }

  return c.Number, err
}

func (c *CreditCard) Valid() bool {
  match, _ := regexp.MatchString("^\\d+$", c.Number)
  return match && luhn.Valid(c.Number)
}

func (c *CreditCard) Type() string {
  for name, pattern := range types {
    match, _ := regexp.MatchString(pattern, c.Number)
    if match {
      return name
    }
  }

  return "Unknown"
}