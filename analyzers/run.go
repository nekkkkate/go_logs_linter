package analyzers

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if call, ok := node.(*ast.CallExpr); ok {
				if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
					if isLogger(pass, selector) {
						if msg, pos, ok := extractArgs(call); ok {
							lowerCaseFirst(pass, msg, pos)
							englishOnly(pass, msg, pos)
							noSpecialChars(pass, msg, pos)
							noSensitiveData(pass, msg, pos)
						}
					}

				}
			}
			return true
		})
	}
	return nil, nil
}

func extractArgs(call *ast.CallExpr) (msg string, pos token.Pos, ok bool) {
	if len(call.Args) == 0 {
		return "", token.NoPos, false
	}

	arg := call.Args[0]

	switch argType := arg.(type) {
	case *ast.BasicLit:
		if argType.Kind == token.STRING {
			if msg, err := strconv.Unquote(argType.Value); err == nil {
				return msg, argType.Pos(), true
			}
		}
		return "", token.NoPos, false
	case *ast.BinaryExpr:
		if argType.Op == token.ADD {
			if left, ok := argType.X.(*ast.BasicLit); ok && left.Kind == token.STRING {
				if msg, err := strconv.Unquote(left.Value); err == nil {
					return msg, left.Pos(), true
				}
				return "", token.NoPos, false
			}
		}
	default:
		return "", token.NoPos, false
	}

	return "", token.NoPos, false

}
