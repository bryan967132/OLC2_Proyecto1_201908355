// Code generated from Parser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// BaseParserListener is a complete listener for a parse tree produced by ParserParser.
type BaseParserListener struct{}

var _ ParserListener = &BaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseParserListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseParserListener) ExitInit(ctx *InitContext) {}

// EnterInstsglobal is called when production instsglobal is entered.
func (s *BaseParserListener) EnterInstsglobal(ctx *InstsglobalContext) {}

// ExitInstsglobal is called when production instsglobal is exited.
func (s *BaseParserListener) ExitInstsglobal(ctx *InstsglobalContext) {}

// EnterInstglobal is called when production instglobal is entered.
func (s *BaseParserListener) EnterInstglobal(ctx *InstglobalContext) {}

// ExitInstglobal is called when production instglobal is exited.
func (s *BaseParserListener) ExitInstglobal(ctx *InstglobalContext) {}

// EnterCallfunc is called when production callfunc is entered.
func (s *BaseParserListener) EnterCallfunc(ctx *CallfuncContext) {}

// ExitCallfunc is called when production callfunc is exited.
func (s *BaseParserListener) ExitCallfunc(ctx *CallfuncContext) {}

// EnterListargs is called when production listargs is entered.
func (s *BaseParserListener) EnterListargs(ctx *ListargsContext) {}

// ExitListargs is called when production listargs is exited.
func (s *BaseParserListener) ExitListargs(ctx *ListargsContext) {}

// EnterDecvar is called when production decvar is entered.
func (s *BaseParserListener) EnterDecvar(ctx *DecvarContext) {}

// ExitDecvar is called when production decvar is exited.
func (s *BaseParserListener) ExitDecvar(ctx *DecvarContext) {}

// EnterDeccst is called when production deccst is entered.
func (s *BaseParserListener) EnterDeccst(ctx *DeccstContext) {}

// ExitDeccst is called when production deccst is exited.
func (s *BaseParserListener) ExitDeccst(ctx *DeccstContext) {}

// EnterDeclfunc is called when production declfunc is entered.
func (s *BaseParserListener) EnterDeclfunc(ctx *DeclfuncContext) {}

// ExitDeclfunc is called when production declfunc is exited.
func (s *BaseParserListener) ExitDeclfunc(ctx *DeclfuncContext) {}

// EnterListparams is called when production listparams is entered.
func (s *BaseParserListener) EnterListparams(ctx *ListparamsContext) {}

// ExitListparams is called when production listparams is exited.
func (s *BaseParserListener) ExitListparams(ctx *ListparamsContext) {}

// EnterIfstruct is called when production ifstruct is entered.
func (s *BaseParserListener) EnterIfstruct(ctx *IfstructContext) {}

// ExitIfstruct is called when production ifstruct is exited.
func (s *BaseParserListener) ExitIfstruct(ctx *IfstructContext) {}

// EnterSwitchstruct is called when production switchstruct is entered.
func (s *BaseParserListener) EnterSwitchstruct(ctx *SwitchstructContext) {}

// ExitSwitchstruct is called when production switchstruct is exited.
func (s *BaseParserListener) ExitSwitchstruct(ctx *SwitchstructContext) {}

// EnterEnvs is called when production envs is entered.
func (s *BaseParserListener) EnterEnvs(ctx *EnvsContext) {}

// ExitEnvs is called when production envs is exited.
func (s *BaseParserListener) ExitEnvs(ctx *EnvsContext) {}

// EnterCasesdefault is called when production casesdefault is entered.
func (s *BaseParserListener) EnterCasesdefault(ctx *CasesdefaultContext) {}

// ExitCasesdefault is called when production casesdefault is exited.
func (s *BaseParserListener) ExitCasesdefault(ctx *CasesdefaultContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseParserListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseParserListener) ExitCases(ctx *CasesContext) {}

// EnterCase is called when production case is entered.
func (s *BaseParserListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseParserListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseParserListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseParserListener) ExitDefault(ctx *DefaultContext) {}

// EnterLoopfor is called when production loopfor is entered.
func (s *BaseParserListener) EnterLoopfor(ctx *LoopforContext) {}

// ExitLoopfor is called when production loopfor is exited.
func (s *BaseParserListener) ExitLoopfor(ctx *LoopforContext) {}

// EnterRange is called when production range is entered.
func (s *BaseParserListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseParserListener) ExitRange(ctx *RangeContext) {}

// EnterLoopwhile is called when production loopwhile is entered.
func (s *BaseParserListener) EnterLoopwhile(ctx *LoopwhileContext) {}

// ExitLoopwhile is called when production loopwhile is exited.
func (s *BaseParserListener) ExitLoopwhile(ctx *LoopwhileContext) {}

// EnterGuard is called when production guard is entered.
func (s *BaseParserListener) EnterGuard(ctx *GuardContext) {}

// ExitGuard is called when production guard is exited.
func (s *BaseParserListener) ExitGuard(ctx *GuardContext) {}

// EnterReasign is called when production reasign is entered.
func (s *BaseParserListener) EnterReasign(ctx *ReasignContext) {}

// ExitReasign is called when production reasign is exited.
func (s *BaseParserListener) ExitReasign(ctx *ReasignContext) {}

// EnterAddsub is called when production addsub is entered.
func (s *BaseParserListener) EnterAddsub(ctx *AddsubContext) {}

// ExitAddsub is called when production addsub is exited.
func (s *BaseParserListener) ExitAddsub(ctx *AddsubContext) {}

// EnterDecvector is called when production decvector is entered.
func (s *BaseParserListener) EnterDecvector(ctx *DecvectorContext) {}

// ExitDecvector is called when production decvector is exited.
func (s *BaseParserListener) ExitDecvector(ctx *DecvectorContext) {}

// EnterDefvector is called when production defvector is entered.
func (s *BaseParserListener) EnterDefvector(ctx *DefvectorContext) {}

// ExitDefvector is called when production defvector is exited.
func (s *BaseParserListener) ExitDefvector(ctx *DefvectorContext) {}

// EnterListexp is called when production listexp is entered.
func (s *BaseParserListener) EnterListexp(ctx *ListexpContext) {}

// ExitListexp is called when production listexp is exited.
func (s *BaseParserListener) ExitListexp(ctx *ListexpContext) {}

// EnterSimplevec is called when production simplevec is entered.
func (s *BaseParserListener) EnterSimplevec(ctx *SimplevecContext) {}

// ExitSimplevec is called when production simplevec is exited.
func (s *BaseParserListener) ExitSimplevec(ctx *SimplevecContext) {}

// EnterFuncvector is called when production funcvector is entered.
func (s *BaseParserListener) EnterFuncvector(ctx *FuncvectorContext) {}

// ExitFuncvector is called when production funcvector is exited.
func (s *BaseParserListener) ExitFuncvector(ctx *FuncvectorContext) {}

// EnterReasignvector is called when production reasignvector is entered.
func (s *BaseParserListener) EnterReasignvector(ctx *ReasignvectorContext) {}

// ExitReasignvector is called when production reasignvector is exited.
func (s *BaseParserListener) ExitReasignvector(ctx *ReasignvectorContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseParserListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseParserListener) ExitPrint(ctx *PrintContext) {}

// EnterEnv is called when production env is entered.
func (s *BaseParserListener) EnterEnv(ctx *EnvContext) {}

// ExitEnv is called when production env is exited.
func (s *BaseParserListener) ExitEnv(ctx *EnvContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseParserListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseParserListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseParserListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseParserListener) ExitInstruction(ctx *InstructionContext) {}

// EnterType is called when production type is entered.
func (s *BaseParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseParserListener) ExitType(ctx *TypeContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseParserListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseParserListener) ExitExp(ctx *ExpContext) {}
