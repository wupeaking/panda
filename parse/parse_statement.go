package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
)

func (p *Parser) parseVarStatement() *ast.VarStatement {
	varStmt := ast.VarStatement{}
	varToken := p.curToken
	p.forwardToken()
	if !p.curTokenIs(token.IDENTIFIER) {
		p.errs = append(p.errs,
			fmt.Errorf("变量声明后面应该是标识符, line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}

	if p.nextTokenIs(token.ASSIGN) {
		varStmt.Token = varToken
		varStmt.Name = &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.forwardToken()
		p.forwardToken()
		varStmt.Value = p.ParseExpression(LOWEST)
		p.forwardToken()
		if !p.curTokenIs(token.SEMI) {
			p.errs = append(p.errs,
				fmt.Errorf("期待一个;在line: %d, pos: %d",
					p.curToken.Line, p.curToken.Position,
				),
			)
			return nil
		}
		return &varStmt

	} else if p.nextTokenIs(token.SEMI) {
		varStmt.Token = varToken
		varStmt.Name = &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.forwardToken()
		return &varStmt
	} else {
		p.errs = append(p.errs,
			fmt.Errorf("变量声明后面应该是等号或者分号, line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
}

func (p *Parser) parseAssginStatement() *ast.AssginStatement {
	assginStmt := ast.AssginStatement{}
	idToken := p.curToken
	p.forwardToken()
	assginStmt.Name = &ast.IdentifierExpression{
		Token: idToken,
		Value: idToken.Literal,
	}
	p.forwardToken()
	assginStmt.Value = p.ParseExpression(LOWEST)
	p.forwardToken()
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个分号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
	}
	return &assginStmt
}

// parseIndexAssginStatement 解析索引赋值语句
// a[0] = xxx
// a["aaaa"] = xxx
func (p *Parser) parseIndexAssginStatement() *ast.AssginStatement {
	assginStmt := ast.AssginStatement{}
	idToken := p.curToken
	p.forwardToken() // [
	assginStmt.Name = &ast.IdentifierExpression{
		Token: idToken,
		Value: idToken.Literal,
	}
	p.forwardToken()
	assginStmt.Index = p.ParseExpression(LOWEST)
	p.forwardToken() // ]
	if !p.curTokenIs(token.RBRACKET) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个] 在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
	if !p.nextTokenIs(token.ASSIGN) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个等号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
	p.forwardToken() // =
	p.forwardToken()
	assginStmt.Value = p.ParseExpression(LOWEST)
	if !p.nextTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个分号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
	p.forwardToken() // ;
	return &assginStmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	retStmt := ast.ReturnStatement{}
	retStmt.Token = p.curToken
	p.forwardToken()
	retStmt.Value = p.ParseExpression(LOWEST)
	p.forwardToken()
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个分号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
	}
	return &retStmt
}

func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	/*
		function aaa() {

		}
	*/
	exp := ast.FunctionStatement{}
	exp.Token = p.curToken // function
	if !p.nextTokenIs(token.IDENTIFIER) {
		p.errs = append(p.errs,
			fmt.Errorf("line: %d, pos: %d 期待function name 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // name
	exp.Name = &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.nextTokenIs(token.LPAREN) {
		p.errs = append(p.errs,
			fmt.Errorf("line: %d, pos: %d 期待( 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // (
	p.forwardToken() // a, b, c 或者)

	args := make([]ast.Expression, 0)
	if p.curTokenIs(token.RPAREN) {
		exp.Args = args
		goto parseBody
	}

	for {
		argExp := p.ParseExpression(LOWEST)
		if argExp == nil {
			return nil
		}
		args = append(args, argExp)

		if p.nextTokenIs(token.RPAREN) {
			break
		}
		if p.nextTokenIs(token.COMMA) {
			p.forwardToken() // ,
			p.forwardToken() //
			continue
		} else {
			p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待,或者) 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
			return nil
		}
	}
	exp.Args = args
	p.forwardToken() // )

parseBody:
	if !p.nextTokenIs(token.LBRACE) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	exp.FuncBody = &ast.BlockStatement{Token: p.curToken}
	p.forwardToken() // {

	p.forwardToken()
	nodes := make([]ast.Node, 0)
	for {
		node := p.parserASTNode()
		if node == nil {
			return nil
		}
		nodes = append(nodes, node)
		if p.nextTokenIs(token.RBRACE) {
			break
		}
		p.forwardToken()
	}
	exp.FuncBody.Statements = nodes
	p.forwardToken() // }
	return &exp
}

func (p *Parser) parseIfStatement() *ast.IFStatement {
	statement := ast.IFStatement{}
	statement.Token = p.curToken // if
	if !p.nextTokenIs(token.LPAREN) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待( 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // (
	p.forwardToken() //
	statement.Condition = p.ParseExpression(LOWEST)
	if !p.nextTokenIs(token.RPAREN) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待) 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // )

	if !p.nextTokenIs(token.LBRACE) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	statement.Consequence = &ast.BlockStatement{Token: p.curToken}
	p.forwardToken() // {

	p.forwardToken()
	nodes := make([]ast.Node, 0)
	for {
		node := p.parserASTNode()
		if node == nil {
			return nil
		}
		nodes = append(nodes, node)
		if p.nextTokenIs(token.RBRACE) {
			break
		}
		p.forwardToken()
	}
	statement.Consequence.Statements = nodes
	p.forwardToken() // }

	// todo:: 暂时还不支持 elif
	if p.nextTokenIs(token.ELSE) {
		// 解析 else 结构
		p.forwardToken() // else
		if !p.nextTokenIs(token.LBRACE) {
			p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
			return nil
		}
		statement.Alternative = &ast.BlockStatement{Token: p.curToken}
		p.forwardToken() // {
		p.forwardToken()

		nodes := make([]ast.Node, 0)
		for {
			node := p.parserASTNode()
			if node == nil {
				return nil
			}
			nodes = append(nodes, node)
			if p.nextTokenIs(token.RBRACE) {
				break
			}
			p.forwardToken()
		}
		statement.Alternative.Statements = nodes
		p.forwardToken() // }
	}

	return &statement
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	forStatement := ast.ForStatement{}
	forStatement.Token = p.curToken // for
	if !p.nextTokenIs(token.LPAREN) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待(实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // (
	p.forwardToken()

	forStatement.FirstCondition = p.parserASTNode()
	_, ok := forStatement.FirstCondition.(ast.Statement)
	if !ok {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d for语句首要循环条件需要是语句声明 ",
			p.curToken.Line, p.curToken.Position))
		return nil
	}
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待; 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() //;
	//p.forwardToken()

	forStatement.MidCondition = p.ParseExpression(LOWEST)
	if !p.nextTokenIs(token.SEMI) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待; 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}

	p.forwardToken() //;
	p.forwardToken()

	// 这里会有点问题 parserASTNode 解析的时候会处理; 但是for最后一个条件没有分号
	// todo:: 目前认定最后一个条件必须是赋值语句
	forStatement.LastCondition = p.parseAssginStatementHelper()
	if !p.nextTokenIs(token.RPAREN) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待) 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // )
	if !p.nextTokenIs(token.LBRACE) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}

	p.forwardToken() // {
	p.forwardToken()
	forStatement.LoopBody = &ast.BlockStatement{Token: p.curToken}
	nodes := make([]ast.Node, 0)
	for {
		node := p.parserASTNode()
		if node == nil {
			return nil
		}
		nodes = append(nodes, node)
		if p.nextTokenIs(token.RBRACE) {
			break
		}
		p.forwardToken()
	}
	forStatement.LoopBody.Statements = nodes
	p.forwardToken() // }
	return &forStatement
}

func (p *Parser) parseAssginStatementHelper() *ast.AssginStatement {
	assginStmt := ast.AssginStatement{}
	idToken := p.curToken
	p.forwardToken()
	assginStmt.Name = &ast.IdentifierExpression{
		Token: idToken,
		Value: idToken.Literal,
	}
	p.forwardToken()
	assginStmt.Value = p.ParseExpression(LOWEST)
	return &assginStmt
}

func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	statement := ast.BreakStatement{}
	statement.Token = p.curToken
	if !p.nextTokenIs(token.SEMI) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.nextToken.Type)))
		return nil
	}
	p.forwardToken() //;
	return &statement
}
