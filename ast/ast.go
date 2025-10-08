package ast

type Node interface {
	// TokenLiteral returns the literal value of the token it’s associated with.
	//will be used for testing and debugging only
	TokenLiteral() string
}

type Statement interface {
	Node // we are basically saying that Statement Extends Node
	statementNode()
}

type Expression interface {
	Node // we are basically saying that Expression Extends Node
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // returns the root node's literal
	} else {
		return ""
	}
}
