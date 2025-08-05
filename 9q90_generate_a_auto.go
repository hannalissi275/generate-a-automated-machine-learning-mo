package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/proebsting/go-pretty/v6"
)

type AutoMLParser struct{}

func (p *AutoMLParser) Parse(model string) (map[string]interface{}, error) {
	is := antlr.NewInputStream(model)
	lexer := NewAutoMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenStreamDefaultTokenChannel)
	parsed, parseErrors := NewAutoMLParser(stream).model()
	if len(parseErrors) > 0 {
		return nil, parseErrors[0]
	}
	return parsed, nil
}

func main() {
	parser := &AutoMLParser{}
	model := `
	model: regression {
		features: [age, country, purchase_history]
		target: purchase_amount
	}
`

	parsed, err := parser.Parse(model)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pretty.Dumper(parsed))
}

type AutoMLLexer struct {
	*antlr.BaseLexer
}

func NewAutoMLLexer(input antlr.CharStream) *AutoMLLexer {
	lexer := &AutoMLLexer{}
	lexer.BaseLexer = antlr.NewBaseLexer(input)
	return lexer
}

func (l *AutoMLLexer) NextToken() antlr.Token {
	for {
		t := l.BaseLexer.NextToken()
		if t.GetTokenType() != antlr.TokenEOF {
			return t
		}
	}
}

type AutoMLParser struct {
	*antlr.BaseParser
}

func NewAutoMLParser(input antlr.TokenStream) *AutoMLParser {
	parser := &AutoMLParser{}
	parser.BaseParser = antlr.NewBaseParser(input)
	return parser
}

func (p *AutoMLParser) model() (map[string]interface{}, antlr.ParseError) {
	listener := &AutoMLListener{}
	p.GetInterpreter().Set-buildListener(listener)
	p.model_()
	return listener.parsed, p.GetInterpreter().syntaxErrors.listener SyntaxError
}

type AutoMLListener struct {
	parsed map[string]interface{}
}

func (l *AutoMLListener) ExitModel(ctx *ModelContext) {
	l.parsed = map[string]interface{}{
		"type": "regression",
		"features": []string{},
		"target": "",
	}
	for _, feature := range ctx.Allfeature() {
		l.parsed["features"] = append(l.parsed["features"].([]string), feature.GetText())
	}
	l.parsed["target"] = ctx.target().GetText()
}

func (l *AutoMLListener) EnterModel_(ctx *Model_Context) {
	l.parsed = map[string]interface{}{}
}

type ModelContext struct {
	*antlr.ParserRuleContextImpl
}

func (m *ModelContext) feature() []*FeatureContext {
	return m.GetTokens(m.Feature)
}

func (m *ModelContext) target() *TargetContext {
	return m.GetRuleContext.(*TargetContext)
}

type FeatureContext struct {
	*antlr.ParserRuleContextImpl
}

func (f *FeatureContext) GetText() string {
	return f.GetText()
}

type TargetContext struct {
	*antlr.ParserRuleContextImpl
}

func (t *TargetContext) GetText() string {
	return t.GetText()
}