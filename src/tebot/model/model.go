package model

import (
  "strings"
)

type TebotOperation struct {
  Action   string
  Value    string
  Selector string
  Script   string
}

func (op *TebotOperation) GetValue() string {
  text := op.Value
  if text == "" {
    return ""
  }
  if strings.Index(text, "\"") == 0 {
    runes := []rune(text)
    // ... Convert back into a string from rune slice.
    return string(runes[1:len(runes)-1])
  }
  return text
}

func (op *TebotOperation) GetSelector() string {
  text := op.Selector
  if text == "" {
    return ""
  }

  if strings.Index(text, "\"") == 0 {
    runes := []rune(text)
    // ... Convert back into a string from rune slice.
    return string(runes[1:len(runes)-1])
  }
  return text
}

func (op *TebotOperation) GetSelectorOrValue() string {
  text := op.GetValue()
  if text == "" {
    text = op.GetSelector()
  }
  return text
}

type TebotOperations struct {
  Operations []*TebotOperation
}
