package parser

import (
  // "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/antlr4-go/antlr/v4"

  "tebot/model"
)

////////////////////////////////////////////////// antlr listener

type MyTebotListener struct {
  BaseTebotListener

  Model *model.TebotOperations
}

func NewMyTebotListener(model *model.TebotOperations) *MyTebotListener {
  ret := new(MyTebotListener)
  ret.Model = model
  return ret
}

func (s *MyTebotListener) EnterTebot_selector(ctx *Tebot_selectorContext) {
  index := len(s.Model.Operations) - 1
  s.Model.Operations[index].Selector = ctx.GetText()
}

func (s *MyTebotListener) EnterTebot_value(ctx *Tebot_valueContext) {
  index := len(s.Model.Operations) - 1
  s.Model.Operations[index].Value = ctx.GetText()
}

func (s *MyTebotListener) EnterTebot_action(ctx *Tebot_actionContext) {
  index := len(s.Model.Operations) - 1
  s.Model.Operations[index].Action = ctx.GetText()
}

func (s *MyTebotListener) EnterTebot_operation(ctx *Tebot_operationContext) {
  op := &model.TebotOperation{}
  op.Script = ctx.GetText()
  s.Model.Operations = append(s.Model.Operations, op)
}

func (s *MyTebotListener) EnterTebot_assert(ctx *Tebot_assertContext) {
  index := len(s.Model.Operations) - 1
  s.Model.Operations[index].Action = "assert"
  s.Model.Operations[index].Assert = &model.TebotAssert{
    Dsn: ctx.GetDsn().GetText(),
    Sql: ctx.GetSql().GetText(),
    Expected: ctx.GetExpected().GetText(),
  }
}

func Parse(expr string) *model.TebotOperations {
  ret := &model.TebotOperations{}
  input := antlr.NewInputStream(expr)

  lexer := NewTebotLexer(input)
  stream := antlr.NewCommonTokenStream(lexer, 0)
  p := NewTebotParser(stream)

  ctxOperations := p.Tebot_operations()

  antlr.ParseTreeWalkerDefault.Walk(NewMyTebotListener(ret), ctxOperations)

  return ret
}