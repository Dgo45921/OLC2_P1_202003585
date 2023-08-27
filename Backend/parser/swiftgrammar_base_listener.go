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

// EnterStructinstruction is called when production structinstruction is entered.
func (s *BaseSwiftGrammarListener) EnterStructinstruction(ctx *StructinstructionContext) {}

// ExitStructinstruction is called when production structinstruction is exited.
func (s *BaseSwiftGrammarListener) ExitStructinstruction(ctx *StructinstructionContext) {}

// EnterStructblock is called when production structblock is entered.
func (s *BaseSwiftGrammarListener) EnterStructblock(ctx *StructblockContext) {}

// ExitStructblock is called when production structblock is exited.
func (s *BaseSwiftGrammarListener) ExitStructblock(ctx *StructblockContext) {}

// EnterStructdef is called when production structdef is entered.
func (s *BaseSwiftGrammarListener) EnterStructdef(ctx *StructdefContext) {}

// ExitStructdef is called when production structdef is exited.
func (s *BaseSwiftGrammarListener) ExitStructdef(ctx *StructdefContext) {}

// EnterVectormodification is called when production vectormodification is entered.
func (s *BaseSwiftGrammarListener) EnterVectormodification(ctx *VectormodificationContext) {}

// ExitVectormodification is called when production vectormodification is exited.
func (s *BaseSwiftGrammarListener) ExitVectormodification(ctx *VectormodificationContext) {}

// EnterForloop is called when production forloop is entered.
func (s *BaseSwiftGrammarListener) EnterForloop(ctx *ForloopContext) {}

// ExitForloop is called when production forloop is exited.
func (s *BaseSwiftGrammarListener) ExitForloop(ctx *ForloopContext) {}

// EnterRange is called when production range is entered.
func (s *BaseSwiftGrammarListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseSwiftGrammarListener) ExitRange(ctx *RangeContext) {}

// EnterRemoveatvec is called when production removeatvec is entered.
func (s *BaseSwiftGrammarListener) EnterRemoveatvec(ctx *RemoveatvecContext) {}

// ExitRemoveatvec is called when production removeatvec is exited.
func (s *BaseSwiftGrammarListener) ExitRemoveatvec(ctx *RemoveatvecContext) {}

// EnterAppendvec is called when production appendvec is entered.
func (s *BaseSwiftGrammarListener) EnterAppendvec(ctx *AppendvecContext) {}

// ExitAppendvec is called when production appendvec is exited.
func (s *BaseSwiftGrammarListener) ExitAppendvec(ctx *AppendvecContext) {}

// EnterRemovelastvec is called when production removelastvec is entered.
func (s *BaseSwiftGrammarListener) EnterRemovelastvec(ctx *RemovelastvecContext) {}

// ExitRemovelastvec is called when production removelastvec is exited.
func (s *BaseSwiftGrammarListener) ExitRemovelastvec(ctx *RemovelastvecContext) {}

// EnterVecdec is called when production vecdec is entered.
func (s *BaseSwiftGrammarListener) EnterVecdec(ctx *VecdecContext) {}

// ExitVecdec is called when production vecdec is exited.
func (s *BaseSwiftGrammarListener) ExitVecdec(ctx *VecdecContext) {}

// EnterBreakstatement is called when production breakstatement is entered.
func (s *BaseSwiftGrammarListener) EnterBreakstatement(ctx *BreakstatementContext) {}

// ExitBreakstatement is called when production breakstatement is exited.
func (s *BaseSwiftGrammarListener) ExitBreakstatement(ctx *BreakstatementContext) {}

// EnterContinuestatement is called when production continuestatement is entered.
func (s *BaseSwiftGrammarListener) EnterContinuestatement(ctx *ContinuestatementContext) {}

// ExitContinuestatement is called when production continuestatement is exited.
func (s *BaseSwiftGrammarListener) ExitContinuestatement(ctx *ContinuestatementContext) {}

// EnterSwitchstatement is called when production switchstatement is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstatement(ctx *SwitchstatementContext) {}

// ExitSwitchstatement is called when production switchstatement is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstatement(ctx *SwitchstatementContext) {}

// EnterCaselist is called when production caselist is entered.
func (s *BaseSwiftGrammarListener) EnterCaselist(ctx *CaselistContext) {}

// ExitCaselist is called when production caselist is exited.
func (s *BaseSwiftGrammarListener) ExitCaselist(ctx *CaselistContext) {}

// EnterCase is called when production case is entered.
func (s *BaseSwiftGrammarListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseSwiftGrammarListener) ExitCase(ctx *CaseContext) {}

// EnterDefaultstatement is called when production defaultstatement is entered.
func (s *BaseSwiftGrammarListener) EnterDefaultstatement(ctx *DefaultstatementContext) {}

// ExitDefaultstatement is called when production defaultstatement is exited.
func (s *BaseSwiftGrammarListener) ExitDefaultstatement(ctx *DefaultstatementContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterEliflist is called when production eliflist is entered.
func (s *BaseSwiftGrammarListener) EnterEliflist(ctx *EliflistContext) {}

// ExitEliflist is called when production eliflist is exited.
func (s *BaseSwiftGrammarListener) ExitEliflist(ctx *EliflistContext) {}

// EnterElif is called when production elif is entered.
func (s *BaseSwiftGrammarListener) EnterElif(ctx *ElifContext) {}

// ExitElif is called when production elif is exited.
func (s *BaseSwiftGrammarListener) ExitElif(ctx *ElifContext) {}

// EnterElsestament is called when production elsestament is entered.
func (s *BaseSwiftGrammarListener) EnterElsestament(ctx *ElsestamentContext) {}

// ExitElsestament is called when production elsestament is exited.
func (s *BaseSwiftGrammarListener) ExitElsestament(ctx *ElsestamentContext) {}

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

// EnterIsemptyvec is called when production isemptyvec is entered.
func (s *BaseSwiftGrammarListener) EnterIsemptyvec(ctx *IsemptyvecContext) {}

// ExitIsemptyvec is called when production isemptyvec is exited.
func (s *BaseSwiftGrammarListener) ExitIsemptyvec(ctx *IsemptyvecContext) {}

// EnterCountvec is called when production countvec is entered.
func (s *BaseSwiftGrammarListener) EnterCountvec(ctx *CountvecContext) {}

// ExitCountvec is called when production countvec is exited.
func (s *BaseSwiftGrammarListener) ExitCountvec(ctx *CountvecContext) {}

// EnterVectoraccess is called when production vectoraccess is entered.
func (s *BaseSwiftGrammarListener) EnterVectoraccess(ctx *VectoraccessContext) {}

// ExitVectoraccess is called when production vectoraccess is exited.
func (s *BaseSwiftGrammarListener) ExitVectoraccess(ctx *VectoraccessContext) {}

// EnterIndexesList is called when production indexesList is entered.
func (s *BaseSwiftGrammarListener) EnterIndexesList(ctx *IndexesListContext) {}

// ExitIndexesList is called when production indexesList is exited.
func (s *BaseSwiftGrammarListener) ExitIndexesList(ctx *IndexesListContext) {}

// EnterVecac is called when production vecac is entered.
func (s *BaseSwiftGrammarListener) EnterVecac(ctx *VecacContext) {}

// ExitVecac is called when production vecac is exited.
func (s *BaseSwiftGrammarListener) ExitVecac(ctx *VecacContext) {}

// EnterMatrix_type is called when production matrix_type is entered.
func (s *BaseSwiftGrammarListener) EnterMatrix_type(ctx *Matrix_typeContext) {}

// ExitMatrix_type is called when production matrix_type is exited.
func (s *BaseSwiftGrammarListener) ExitMatrix_type(ctx *Matrix_typeContext) {}

// EnterRepeatingvector is called when production repeatingvector is entered.
func (s *BaseSwiftGrammarListener) EnterRepeatingvector(ctx *RepeatingvectorContext) {}

// ExitRepeatingvector is called when production repeatingvector is exited.
func (s *BaseSwiftGrammarListener) ExitRepeatingvector(ctx *RepeatingvectorContext) {}

// EnterManualdef is called when production manualdef is entered.
func (s *BaseSwiftGrammarListener) EnterManualdef(ctx *ManualdefContext) {}

// ExitManualdef is called when production manualdef is exited.
func (s *BaseSwiftGrammarListener) ExitManualdef(ctx *ManualdefContext) {}

// EnterManualmatrixdef is called when production manualmatrixdef is entered.
func (s *BaseSwiftGrammarListener) EnterManualmatrixdef(ctx *ManualmatrixdefContext) {}

// ExitManualmatrixdef is called when production manualmatrixdef is exited.
func (s *BaseSwiftGrammarListener) ExitManualmatrixdef(ctx *ManualmatrixdefContext) {}

// EnterValues2 is called when production values2 is entered.
func (s *BaseSwiftGrammarListener) EnterValues2(ctx *Values2Context) {}

// ExitValues2 is called when production values2 is exited.
func (s *BaseSwiftGrammarListener) ExitValues2(ctx *Values2Context) {}

// EnterDecmatrix is called when production decmatrix is entered.
func (s *BaseSwiftGrammarListener) EnterDecmatrix(ctx *DecmatrixContext) {}

// ExitDecmatrix is called when production decmatrix is exited.
func (s *BaseSwiftGrammarListener) ExitDecmatrix(ctx *DecmatrixContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}
