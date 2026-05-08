package model

import (
  "strings"
)

type TebotAssert struct {
  Dsn      string
  Sql      string
  Expected string
}

func (a *TebotAssert) GetDsn() string {
  return a.Dsn
}

func (a *TebotAssert) GetSql() string {
  return a.Sql
}

func (a *TebotAssert) GetExpected() string {
  return a.Expected
}

type TebotOperation struct {
  Action   string
  Value    string
  Selector string
  Script   string
  Assert   *TebotAssert
}

func (op *TebotOperation) GetValue() string {
  text := op.Value
  if text == "" {
    return ""
  }
  if strings.Index(text, "\"") == 0 {
    runes := []rune(text)
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
