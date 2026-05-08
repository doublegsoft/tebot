// Code generated from Tebot.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Tebot

import "github.com/antlr4-go/antlr/v4"

// BaseTebotListener is a complete listener for a parse tree produced by TebotParser.
type BaseTebotListener struct{}

var _ TebotListener = &BaseTebotListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTebotListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTebotListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTebotListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTebotListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTebot_selector is called when production tebot_selector is entered.
func (s *BaseTebotListener) EnterTebot_selector(ctx *Tebot_selectorContext) {}

// ExitTebot_selector is called when production tebot_selector is exited.
func (s *BaseTebotListener) ExitTebot_selector(ctx *Tebot_selectorContext) {}

// EnterTebot_value is called when production tebot_value is entered.
func (s *BaseTebotListener) EnterTebot_value(ctx *Tebot_valueContext) {}

// ExitTebot_value is called when production tebot_value is exited.
func (s *BaseTebotListener) ExitTebot_value(ctx *Tebot_valueContext) {}

// EnterTebot_action is called when production tebot_action is entered.
func (s *BaseTebotListener) EnterTebot_action(ctx *Tebot_actionContext) {}

// ExitTebot_action is called when production tebot_action is exited.
func (s *BaseTebotListener) ExitTebot_action(ctx *Tebot_actionContext) {}

// EnterTebot_assign is called when production tebot_assign is entered.
func (s *BaseTebotListener) EnterTebot_assign(ctx *Tebot_assignContext) {}

// ExitTebot_assign is called when production tebot_assign is exited.
func (s *BaseTebotListener) ExitTebot_assign(ctx *Tebot_assignContext) {}

// EnterTebot_operation is called when production tebot_operation is entered.
func (s *BaseTebotListener) EnterTebot_operation(ctx *Tebot_operationContext) {}

// ExitTebot_operation is called when production tebot_operation is exited.
func (s *BaseTebotListener) ExitTebot_operation(ctx *Tebot_operationContext) {}

// EnterTebot_assert is called when production tebot_assert is entered.
func (s *BaseTebotListener) EnterTebot_assert(ctx *Tebot_assertContext) {}

// ExitTebot_assert is called when production tebot_assert is exited.
func (s *BaseTebotListener) ExitTebot_assert(ctx *Tebot_assertContext) {}

// EnterTebot_operations is called when production tebot_operations is entered.
func (s *BaseTebotListener) EnterTebot_operations(ctx *Tebot_operationsContext) {}

// ExitTebot_operations is called when production tebot_operations is exited.
func (s *BaseTebotListener) ExitTebot_operations(ctx *Tebot_operationsContext) {}
