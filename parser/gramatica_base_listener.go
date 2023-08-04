// Code generated from gramatica.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr4-go/antlr/v4"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasegramaticaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasegramaticaListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasegramaticaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasegramaticaListener) ExitExpr(ctx *ExprContext) {}
