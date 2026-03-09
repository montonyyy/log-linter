package logcheck

import (
	"go/ast"
	"go/types"
	"log/slog"
	"os"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "check for format errors in logs",
	Run:  run,
}

var sensitiveWords string

func getConf() {
	data, err := os.ReadFile("sens.conf")
	if err != nil {
		slog.Info("configuration file for sensitive words not exists. will be used default config")
		sensitiveWords = "password, apiKey, token"
		return
	}
	sensitiveWords = string(data)
}

func run(pass *analysis.Pass) (any, error) {
	getConf()
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			left, ok := sel.X.(*ast.Ident)
			if !ok {
				return true
			}
			obj := pass.TypesInfo.Uses[left]

			if pkg, ok := obj.(*types.PkgName); ok {
				path := pkg.Imported().Path()
				if path != "log" && path != "log/slog" && path != "go.uber.org/zap" {
					return true
				}
			} else {
				selection := pass.TypesInfo.Selections[sel]
				recv := selection.Recv()
				if ptr, ok := recv.(*types.Pointer); ok {
					recv = ptr.Elem()
				}
				named, ok := recv.(*types.Named)
				if !ok {
					return true
				}
				namedObj := named.Obj()
				if namedObj == nil || namedObj.Pkg() == nil {
					return true
				}
				path := namedObj.Pkg().Path()
				if path != "log" && path != "log/slog" && path != "go.uber.org/zap" {
					return true
				}
			}

			for _, arg := range call.Args {
				checkArg(pass, arg)
			}
			return true
		})
	}
	return nil, nil
}

func checkArg(pass *analysis.Pass, arg ast.Expr) {
	if bin, ok := arg.(*ast.BinaryExpr); ok {
		checkArg(pass, bin.X)
		checkArg(pass, bin.Y)
	} else {
		checkBasicLit(pass, arg)
		checkIdent(pass, arg)
	}
}

func checkBasicLit(pass *analysis.Pass, arg ast.Expr) {
	if lit, ok := arg.(*ast.BasicLit); ok {
		var str string

		if len(lit.Value) > 2 {
			str = lit.Value[1 : len(lit.Value)-1]
		}

		if unicode.IsUpper(rune(str[0])) {
			pass.Reportf(lit.Pos(), "contains capital letter")
		}
		isLatin := true
		isSymbol := false

		for _, v := range str {
			if unicode.IsLetter(v) && !unicode.Is(unicode.Latin, v) && isLatin == true {
				isLatin = false
				pass.Reportf(lit.Pos(), "contains not an english letter")
			}
			if !unicode.IsLetter(v) && !unicode.IsDigit(v) && isSymbol == false {
				isSymbol = true
				pass.Reportf(lit.Pos(), "contains symbol letter")
			}
			if !isLatin && isSymbol {
				break
			}
		}
	}
}

func checkIdent(pass *analysis.Pass, arg ast.Expr) {
	if ident, ok := arg.(*ast.Ident); ok {
		val := ident.Name
		if strings.Contains(sensitiveWords, val) {
			pass.Reportf(ident.Pos(), "contains sensitive data")
		}
	}
}
