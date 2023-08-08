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

	// EnterDecmatrix is called when entering the decmatrix production.
	EnterDecmatrix(c *DecmatrixContext)

	// EnterTypematrix is called when entering the typematrix production.
	EnterTypematrix(c *TypematrixContext)

	// EnterDefmatrix is called when entering the defmatrix production.
	EnterDefmatrix(c *DefmatrixContext)

	// EnterListvector is called when entering the listvector production.
	EnterListvector(c *ListvectorContext)

	// EnterListvector2 is called when entering the listvector2 production.
	EnterListvector2(c *Listvector2Context)

	// EnterSimplematrix is called when entering the simplematrix production.
	EnterSimplematrix(c *SimplematrixContext)

	// EnterDefstruct is called when entering the defstruct production.
	EnterDefstruct(c *DefstructContext)

	// EnterListattribs is called when entering the listattribs production.
	EnterListattribs(c *ListattribsContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterDecstruct is called when entering the decstruct production.
	EnterDecstruct(c *DecstructContext)

	// EnterListdupla is called when entering the listdupla production.
	EnterListdupla(c *ListduplaContext)

	// EnterUseattribs is called when entering the useattribs production.
	EnterUseattribs(c *UseattribsContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterUseattribs1 is called when entering the useattribs1 production.
	EnterUseattribs1(c *Useattribs1Context)

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

	// ExitDecmatrix is called when exiting the decmatrix production.
	ExitDecmatrix(c *DecmatrixContext)

	// ExitTypematrix is called when exiting the typematrix production.
	ExitTypematrix(c *TypematrixContext)

	// ExitDefmatrix is called when exiting the defmatrix production.
	ExitDefmatrix(c *DefmatrixContext)

	// ExitListvector is called when exiting the listvector production.
	ExitListvector(c *ListvectorContext)

	// ExitListvector2 is called when exiting the listvector2 production.
	ExitListvector2(c *Listvector2Context)

	// ExitSimplematrix is called when exiting the simplematrix production.
	ExitSimplematrix(c *SimplematrixContext)

	// ExitDefstruct is called when exiting the defstruct production.
	ExitDefstruct(c *DefstructContext)

	// ExitListattribs is called when exiting the listattribs production.
	ExitListattribs(c *ListattribsContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitDecstruct is called when exiting the decstruct production.
	ExitDecstruct(c *DecstructContext)

	// ExitListdupla is called when exiting the listdupla production.
	ExitListdupla(c *ListduplaContext)

	// ExitUseattribs is called when exiting the useattribs production.
	ExitUseattribs(c *UseattribsContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitUseattribs1 is called when exiting the useattribs1 production.
	ExitUseattribs1(c *Useattribs1Context)

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
