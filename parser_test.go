package wt_concurrency

import "testing"

func TestAST(test *testing.T) {
	writer := Actor{0, "Writer"}
	reader := Actor{1, "Reader"}

	instance := Instance{
		TableName: "tableUri",
		Actors:    []Actor{writer, reader},
	}
	instance.NextOp(BeginTxn{Actor: writer})
	instance.NextOp(BeginTxn{Actor: reader})
	instance.NextOp(Write{Actor: writer, Key: 1, Value: 1})
	instance.NextOp(CommitTxn{Actor: writer})
	instance.NextOp(Read{Actor: reader, Key: 1})
	instance.NextOp(CommitTxn{Actor: reader})

	instance.Compile("artifacts/test_ast.cpp")
}
