// Code generated from Parser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// ParserListener is a complete listener for a parse tree produced by ParserParser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterInstsglobal is called when entering the instsglobal production.
	EnterInstsglobal(c *InstsglobalContext)

	// EnterInstglobal is called when entering the instglobal production.
	EnterInstglobal(c *InstglobalContext)

	// EnterCallfunc is called when entering the callfunc production.
	EnterCallfunc(c *CallfuncContext)

	// EnterListargs is called when entering the listargs production.
	EnterListargs(c *ListargsContext)

	// EnterDecvar is called when entering the decvar production.
	EnterDecvar(c *DecvarContext)

	// EnterDeccst is called when entering the deccst production.
	EnterDeccst(c *DeccstContext)

	// EnterDeclfunc is called when entering the declfunc production.
	EnterDeclfunc(c *DeclfuncContext)

	// EnterListparams is called when entering the listparams production.
	EnterListparams(c *ListparamsContext)

	// EnterIfstruct is called when entering the ifstruct production.
	EnterIfstruct(c *IfstructContext)

	// EnterSwitchstruct is called when entering the switchstruct production.
	EnterSwitchstruct(c *SwitchstructContext)

	// EnterEnvs is called when entering the envs production.
	EnterEnvs(c *EnvsContext)

	// EnterCasesdefault is called when entering the casesdefault production.
	EnterCasesdefault(c *CasesdefaultContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterLoopfor is called when entering the loopfor production.
	EnterLoopfor(c *LoopforContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterLoopwhile is called when entering the loopwhile production.
	EnterLoopwhile(c *LoopwhileContext)

	// EnterGuard is called when entering the guard production.
	EnterGuard(c *GuardContext)

	// EnterReasign is called when entering the reasign production.
	EnterReasign(c *ReasignContext)

	// EnterAddsub is called when entering the addsub production.
	EnterAddsub(c *AddsubContext)

	// EnterDecvector is called when entering the decvector production.
	EnterDecvector(c *DecvectorContext)

	// EnterDefvector is called when entering the defvector production.
	EnterDefvector(c *DefvectorContext)

	// EnterListexp is called when entering the listexp production.
	EnterListexp(c *ListexpContext)

	// EnterSimplevec is called when entering the simplevec production.
	EnterSimplevec(c *SimplevecContext)

	// EnterFuncvector is called when entering the funcvector production.
	EnterFuncvector(c *FuncvectorContext)

	// EnterReasignvector is called when entering the reasignvector production.
	EnterReasignvector(c *ReasignvectorContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterEnv is called when entering the env production.
	EnterEnv(c *EnvContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitInstsglobal is called when exiting the instsglobal production.
	ExitInstsglobal(c *InstsglobalContext)

	// ExitInstglobal is called when exiting the instglobal production.
	ExitInstglobal(c *InstglobalContext)

	// ExitCallfunc is called when exiting the callfunc production.
	ExitCallfunc(c *CallfuncContext)

	// ExitListargs is called when exiting the listargs production.
	ExitListargs(c *ListargsContext)

	// ExitDecvar is called when exiting the decvar production.
	ExitDecvar(c *DecvarContext)

	// ExitDeccst is called when exiting the deccst production.
	ExitDeccst(c *DeccstContext)

	// ExitDeclfunc is called when exiting the declfunc production.
	ExitDeclfunc(c *DeclfuncContext)

	// ExitListparams is called when exiting the listparams production.
	ExitListparams(c *ListparamsContext)

	// ExitIfstruct is called when exiting the ifstruct production.
	ExitIfstruct(c *IfstructContext)

	// ExitSwitchstruct is called when exiting the switchstruct production.
	ExitSwitchstruct(c *SwitchstructContext)

	// ExitEnvs is called when exiting the envs production.
	ExitEnvs(c *EnvsContext)

	// ExitCasesdefault is called when exiting the casesdefault production.
	ExitCasesdefault(c *CasesdefaultContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitLoopfor is called when exiting the loopfor production.
	ExitLoopfor(c *LoopforContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitLoopwhile is called when exiting the loopwhile production.
	ExitLoopwhile(c *LoopwhileContext)

	// ExitGuard is called when exiting the guard production.
	ExitGuard(c *GuardContext)

	// ExitReasign is called when exiting the reasign production.
	ExitReasign(c *ReasignContext)

	// ExitAddsub is called when exiting the addsub production.
	ExitAddsub(c *AddsubContext)

	// ExitDecvector is called when exiting the decvector production.
	ExitDecvector(c *DecvectorContext)

	// ExitDefvector is called when exiting the defvector production.
	ExitDefvector(c *DefvectorContext)

	// ExitListexp is called when exiting the listexp production.
	ExitListexp(c *ListexpContext)

	// ExitSimplevec is called when exiting the simplevec production.
	ExitSimplevec(c *SimplevecContext)

	// ExitFuncvector is called when exiting the funcvector production.
	ExitFuncvector(c *FuncvectorContext)

	// ExitReasignvector is called when exiting the reasignvector production.
	ExitReasignvector(c *ReasignvectorContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitEnv is called when exiting the env production.
	ExitEnv(c *EnvContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)
}
