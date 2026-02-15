package analyzers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var logMethods = map[string]bool{
	"Debug": true,
	"Info":  true,
	"Warn":  true,
	"Error": true,
	"Fatal": true,
}

func isLogger(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	if pass.TypesInfo == nil || !logMethods[selector.Sel.Name] {
		return false
	}

	return isSlog(pass, selector) || isZap(pass, selector)
}

func isSlog(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	if ident, ok := selector.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkgName, ok := obj.(*types.PkgName); ok {
				return pkgName.Imported().Path() == "log/slog"
			}
		}
	}

	if typ := pass.TypesInfo.TypeOf(selector.X); typ != nil {
		if ptr, ok := typ.(*types.Pointer); ok {
			typ = ptr.Elem()
		}

		if named, ok := typ.(*types.Named); ok {
			if pkgPath := named.Obj().Pkg(); pkgPath != nil {
				return pkgPath.Path() == "log/slog" && named.Obj().Name() == "Logger"
			}
		}
	}

	return false
}

func isZap(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	if typ := pass.TypesInfo.TypeOf(selector.X); typ != nil {
		if ptr, ok := typ.(*types.Pointer); ok {
			if named, ok := ptr.Elem().(*types.Named); ok {
				if pkgPath := named.Obj().Pkg(); pkgPath != nil {
					return pkgPath.Path() == "go.uber.org/zap" && named.Obj().Name() == "Logger"
				}
			}
		}
	}
	return false
}
