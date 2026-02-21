// Code generated from Tebot.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Tebot

import "github.com/antlr4-go/antlr/v4"

// TebotListener is a complete listener for a parse tree produced by TebotParser.
type TebotListener interface {
	antlr.ParseTreeListener

	// EnterTebot_selector is called when entering the tebot_selector production.
	EnterTebot_selector(c *Tebot_selectorContext)

	// EnterTebot_value is called when entering the tebot_value production.
	EnterTebot_value(c *Tebot_valueContext)

	// EnterTebot_action is called when entering the tebot_action production.
	EnterTebot_action(c *Tebot_actionContext)

	// EnterTebot_assign is called when entering the tebot_assign production.
	EnterTebot_assign(c *Tebot_assignContext)

	// EnterTebot_operation is called when entering the tebot_operation production.
	EnterTebot_operation(c *Tebot_operationContext)

	// EnterTebot_operations is called when entering the tebot_operations production.
	EnterTebot_operations(c *Tebot_operationsContext)

	// ExitTebot_selector is called when exiting the tebot_selector production.
	ExitTebot_selector(c *Tebot_selectorContext)

	// ExitTebot_value is called when exiting the tebot_value production.
	ExitTebot_value(c *Tebot_valueContext)

	// ExitTebot_action is called when exiting the tebot_action production.
	ExitTebot_action(c *Tebot_actionContext)

	// ExitTebot_assign is called when exiting the tebot_assign production.
	ExitTebot_assign(c *Tebot_assignContext)

	// ExitTebot_operation is called when exiting the tebot_operation production.
	ExitTebot_operation(c *Tebot_operationContext)

	// ExitTebot_operations is called when exiting the tebot_operations production.
	ExitTebot_operations(c *Tebot_operationsContext)
}
