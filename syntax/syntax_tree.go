package syntax

import "github.com/techerfan/expression/contracts"

var _parser *parser

// A tree structure for storing the syntax of the given expression
type SyntaxTree struct {
	Root           contracts.ExpressionSyntax
	Diagnostics    []*contracts.Diagnostic
	endOfFileToken *contracts.SyntaxToken
	tokens         []*contracts.SyntaxToken
}

// Initializes a new instance of SyntaxTree
func NewSyntaxTree(
	diagnostics []*contracts.Diagnostic,
	root contracts.ExpressionSyntax,
	endOfFileToken *contracts.SyntaxToken,
	tokens []*contracts.SyntaxToken,
) *SyntaxTree {
	return &SyntaxTree{
		Diagnostics:    diagnostics,
		Root:           root,
		endOfFileToken: endOfFileToken,
		tokens:         tokens,
	}
}

// Parses the given expression
func Parse(text string) *SyntaxTree {
	_parser = newParser(text)
	return _parser.Parse()
}

// Returns all the tokens used in the given expression
func (s *SyntaxTree) Tokens() []*contracts.SyntaxToken {
	if _parser == nil {
		return []*contracts.SyntaxToken{}
	}
	return _parser.Tokens()
}

// Returns all the variables used in the given expression
func (s *SyntaxTree) Variables() []string {
	var variables []string
	for _, token := range s.tokens {
		if token.Kind() == contracts.IdentifierToken {
			variables = append(variables, token.Text)
		}
	}

	return variables
}
