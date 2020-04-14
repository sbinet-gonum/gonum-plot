// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser // import "gonum.org/v1/plot/internal/latex/parser"

import (
	"fmt"
	"log"
	"strings"

	"gonum.org/v1/plot/internal/latex/ast"
	"gonum.org/v1/plot/internal/latex/token"
)

func ParseExpr(x string) (ast.Node, error) {
	p := newParser(x)
	return p.parse()
}

type state int

const (
	normalState state = iota
	mathState
)

type parser struct {
	s     *texScanner
	state state

	macros map[string]macroParser
}

func newParser(x string) *parser {
	p := &parser{
		s: newScanner(strings.NewReader(x)),
	}
	p.addBuiltinMacros()
	return p
}

func (p *parser) parse() (ast.Node, error) {
	var nodes []ast.Node
	for p.s.Next() {
		tok := p.s.Token()
		log.Printf("--- token: %#v (tok:%v)", tok, tok.Kind)
		node := p.parseNode(tok)
		if node == nil {
			continue
		}
		nodes = append(nodes, node)
	}

	return ast.List(nodes), nil
}

func (p *parser) next() token.Token {
	if !p.s.Next() {
		return token.Token{Kind: token.EOF}
	}
	return p.s.tok
}

func (p *parser) expect(v rune) {
	p.next()
	if p.s.tok.Text != string(v) {
		panic(fmt.Errorf("expected %q, got %q", v, p.s.tok.Text))
	}
}

func (p *parser) parseNode(tok token.Token) ast.Node {
	switch tok.Kind {
	case token.Comment:
		return nil
	case token.Macro:
		return p.parseMacro(tok)
	case token.Word:
		return p.parseWord(tok)
	case token.Number:
		return p.parseNumber(tok)
	case token.Dollar:
		return p.parseMathExpr(tok)
	case token.Hat:
		return p.parseHat(tok)
	case token.Sub, token.Add, token.Mul, token.Div, token.Equal:
		return p.parseOp(tok)
	case token.Lbrace:
		switch p.state {
		case mathState:
			return p.parseMathLbrace(tok)
		default:
			panic("not implemented")
		}
	case token.Other:
		switch tok.Text {
		default:
			panic("not implemented: " + tok.String())
		}
	default:
		panic("impossible: " + tok.String())
	}
}

func (p *parser) parseMathExpr(tok token.Token) ast.Node {
	state := p.state
	p.state = mathState
	defer func() {
		p.state = state
	}()

	math := &ast.MathExpr{
		Left: tok.Pos,
	}
loop:
	for p.s.Next() {
		switch p.s.tok.Kind {
		case token.Dollar:
			math.Right = p.s.tok.Pos
			break loop
		default:
			node := p.parseNode(p.s.tok)
			if node == nil {
				continue
			}
			math.List = append(math.List, node)
		}
	}

	return math
}

func (p *parser) parseMacro(tok token.Token) ast.Node {
	name := tok.Text
	macro, ok := p.macros[name]
	if !ok {
		panic("unknown macro " + name)
		//return nil
	}
	return macro.parseMacro(p)
}

func (p *parser) parseWord(tok token.Token) ast.Node {
	return &ast.Var{
		VarPos: tok.Pos,
		Name:   tok.Text,
	}
}

func (p *parser) parseNumber(tok token.Token) ast.Node {
	return &ast.Literal{
		LitPos: tok.Pos,
		Text:   tok.Text,
	}
}

func (p *parser) parseMacroArg(macro *ast.Macro) {
	var arg ast.Arg
	p.expect('{')
	arg.Lbrace = p.s.tok.Pos

loop:
	for p.s.Next() {
		switch p.s.tok.Kind {
		case token.Rbrace:
			arg.Rbrace = p.s.tok.Pos
			break loop
		default:
			node := p.parseNode(p.s.tok)
			if node == nil {
				continue
			}
			arg.List = append(arg.List, node)
		}
	}
	macro.Args = append(macro.Args, &arg)
}

func (p *parser) parseOptMacroArg(macro *ast.Macro) {
	nxt := p.s.sc.Peek()
	if nxt != '[' {
		return
	}

	var opt ast.OptArg

	p.expect('[')
	opt.Lbrack = p.s.tok.Pos

loop:
	for p.s.Next() {
		switch p.s.tok.Kind {
		case token.Rbrack:
			opt.Rbrack = p.s.tok.Pos
			break loop
		default:
			node := p.parseNode(p.s.tok)
			if node == nil {
				continue
			}
			opt.List = append(opt.List, node)
		}
	}
	macro.Args = append(macro.Args, &opt)
}

func (p *parser) parseVerbatimMacroArg(macro *ast.Macro) {
}

func (p *parser) parseHat(tok token.Token) ast.Node {
	hat := &ast.Super{
		HatPos: tok.Pos,
	}

	switch next := p.s.sc.Peek(); next {
	case '{':
		p.expect('{')
		var list ast.List
	loop:
		for p.s.Next() {
			switch p.s.tok.Kind {
			case token.Rbrace:
				break loop
			default:
				node := p.parseNode(p.s.tok)
				if node == nil {
					continue
				}
				list = append(list, node)
			}
		}
		hat.Node = list
	default:
		hat.Node = p.parseNode(p.next())
	}

	return hat
}

func (p *parser) parseOp(tok token.Token) ast.Node {
	return &ast.Op{
		OpPos: tok.Pos,
		Text:  tok.Text,
	}
}

func (p *parser) parseMathLbrace(tok token.Token) ast.Node {
	var (
		lst    ast.List
		ldelim = tok.Kind
		rdelim = map[token.Kind]token.Kind{
			token.Lbrace: token.Rbrace,
			token.Lparen: token.Rparen,
		}[ldelim]
	)

	if rdelim == token.Invalid {
		panic("impossible: no matching right-delim for: " + tok.String())
	}

loop:
	for p.s.Next() {
		switch p.s.tok.Kind {
		case rdelim:
			break loop
		default:
			node := p.parseNode(p.s.tok)
			if node == nil {
				continue
			}
			lst = append(lst, node)
		}
	}
	return lst
}
