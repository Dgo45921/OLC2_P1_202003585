// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseSwiftGrammarListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseSwiftGrammarListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseSwiftGrammarListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseSwiftGrammarListener) ExitArgument(ctx *ArgumentContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterBreakstatement is called when production breakstatement is entered.
func (s *BaseSwiftGrammarListener) EnterBreakstatement(ctx *BreakstatementContext) {}

// ExitBreakstatement is called when production breakstatement is exited.
func (s *BaseSwiftGrammarListener) ExitBreakstatement(ctx *BreakstatementContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElseifstatement is called when production elseifstatement is entered.
func (s *BaseSwiftGrammarListener) EnterElseifstatement(ctx *ElseifstatementContext) {}

// ExitElseifstatement is called when production elseifstatement is exited.
func (s *BaseSwiftGrammarListener) ExitElseifstatement(ctx *ElseifstatementContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseSwiftGrammarListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseSwiftGrammarListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterVardec is called when production vardec is entered.
func (s *BaseSwiftGrammarListener) EnterVardec(ctx *VardecContext) {}

// ExitVardec is called when production vardec is exited.
func (s *BaseSwiftGrammarListener) ExitVardec(ctx *VardecContext) {}

// EnterConstdec is called when production constdec is entered.
func (s *BaseSwiftGrammarListener) EnterConstdec(ctx *ConstdecContext) {}

// ExitConstdec is called when production constdec is exited.
func (s *BaseSwiftGrammarListener) ExitConstdec(ctx *ConstdecContext) {}

// EnterAsignation is called when production asignation is entered.
func (s *BaseSwiftGrammarListener) EnterAsignation(ctx *AsignationContext) {}

// ExitAsignation is called when production asignation is exited.
func (s *BaseSwiftGrammarListener) ExitAsignation(ctx *AsignationContext) {}

// EnterUnarysum is called when production unarysum is entered.
func (s *BaseSwiftGrammarListener) EnterUnarysum(ctx *UnarysumContext) {}

// ExitUnarysum is called when production unarysum is exited.
func (s *BaseSwiftGrammarListener) ExitUnarysum(ctx *UnarysumContext) {}

// EnterUnarysub is called when production unarysub is entered.
func (s *BaseSwiftGrammarListener) EnterUnarysub(ctx *UnarysubContext) {}

// ExitUnarysub is called when production unarysub is exited.
func (s *BaseSwiftGrammarListener) ExitUnarysub(ctx *UnarysubContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}
