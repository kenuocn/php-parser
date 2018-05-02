package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestContinueEmpty(t *testing.T) {
	t.Helper()
	src := `<? while (1) { continue; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Continue{Expr: nil},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestContinueLight(t *testing.T) {
	t.Helper()
	src := `<? while (1) { continue 2; }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Continue{
								Expr: &scalar.Lnumber{Value: "2"},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestContinue(t *testing.T) {
	t.Helper()
	src := `<? while (1) { continue(3); }`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.While{
				Cond: &scalar.Lnumber{Value: "1"},
				Stmt: &stmt.InnerStmtList{
					Stmts: &stmt.StmtList{
						Stmts: []node.Node{
							&stmt.Continue{
								Expr: &scalar.Lnumber{Value: "3"},
							},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
