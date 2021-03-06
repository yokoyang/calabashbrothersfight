package execution

import (
	"testing"
	"fmt"
)
type Level struct {
	Label          string
	Title          string
	Description    string
	ThreadContexts []*ThreadContext
	GlobalContext  *GlobalContext
}
func TestNewThread(t *testing.T) {
	gc := NewGlobalContext(Pair{
		"a",
		GlobalStateType{Value: 0, Name: "a"},
	})
	//gc.Values["a"] = GlobalStateType{Value: 0, Name: "a"}
	/**
	a = 0;
	if (a == 3) {
		a = 10;
	}
	print(a);
	 */
	assign := NewAssignmentInstruction("a", NewLiteralExpression(3))
	assignIn := NewAssignmentInstruction("a", NewLiteralExpression(10))
	ifUp := NewStartIfStatement(
		NewEqualityExpression(NewVariableExpression("a"), NewLiteralExpression(3)),
		"if1")
	ifEnd := NewEndIfStatement("if1")

	context := NewThreadContext(0, 0, 0, []Instruction{assign, ifUp, assignIn, ifEnd})
	for i := 0; i < 3; {
		ins := context.Instructions[i]
		ins.Execute(gc, context)
		i = context.ProgramCounter
	}
	fmt.Println(gc.Values["a"])
}

func TestLevel(t *testing.T) {
	level := &Level{

	}
	gc := level.GlobalContext
	ins := level.ThreadContexts[0].Instructions
	le := len(*ins)
	for i := 0; i < le; {
		ii := (*ins)[i]
		ii.Execute(gc, level.ThreadContexts[0])
		i = level.ThreadContexts[0].ProgramCounter
	}
}
