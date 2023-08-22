// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterVecdec is called when entering the vecdec production.
	EnterVecdec(c *VecdecContext)

	// EnterBreakstatement is called when entering the breakstatement production.
	EnterBreakstatement(c *BreakstatementContext)

	// EnterContinuestatement is called when entering the continuestatement production.
	EnterContinuestatement(c *ContinuestatementContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElseifstatement is called when entering the elseifstatement production.
	EnterElseifstatement(c *ElseifstatementContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterVardec is called when entering the vardec production.
	EnterVardec(c *VardecContext)

	// EnterConstdec is called when entering the constdec production.
	EnterConstdec(c *ConstdecContext)

	// EnterAsignation is called when entering the asignation production.
	EnterAsignation(c *AsignationContext)

	// EnterUnarysum is called when entering the unarysum production.
	EnterUnarysum(c *UnarysumContext)

	// EnterUnarysub is called when entering the unarysub production.
	EnterUnarysub(c *UnarysubContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitVecdec is called when exiting the vecdec production.
	ExitVecdec(c *VecdecContext)

	// ExitBreakstatement is called when exiting the breakstatement production.
	ExitBreakstatement(c *BreakstatementContext)

	// ExitContinuestatement is called when exiting the continuestatement production.
	ExitContinuestatement(c *ContinuestatementContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElseifstatement is called when exiting the elseifstatement production.
	ExitElseifstatement(c *ElseifstatementContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitVardec is called when exiting the vardec production.
	ExitVardec(c *VardecContext)

	// ExitConstdec is called when exiting the constdec production.
	ExitConstdec(c *ConstdecContext)

	// ExitAsignation is called when exiting the asignation production.
	ExitAsignation(c *AsignationContext)

	// ExitUnarysum is called when exiting the unarysum production.
	ExitUnarysum(c *UnarysumContext)

	// ExitUnarysub is called when exiting the unarysub production.
	ExitUnarysub(c *UnarysubContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
