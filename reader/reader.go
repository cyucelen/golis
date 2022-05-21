package reader

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/cyucelen/golis/fn"
	"github.com/cyucelen/golis/types"
)

type Reader struct {
	tokens []string
	pos    int
}

func (r *Reader) Next() string {
	token := r.tokens[r.pos]
	r.pos++
	return token
}

func (r *Reader) Peek() string {
	return r.tokens[r.pos]
}

func (r *Reader) ReadForm() (types.Object, error) {
	token := r.Peek()
	switch token[0] {
	case '(':
		objects, err := r.ReadUntil(')')
		return types.NewList(objects...), err
	case '[':
		objects, err := r.ReadUntil(']')
		return types.NewVector(objects...), err
	case '{':
		objects, err := r.ReadUntil('}')
		if err != nil {
			return nil, err
		}
		return types.NewHashMap(objects...)
	default:
		return r.ReadAtom()
	}
}

func (r *Reader) ReadUntil(ending byte) ([]types.Object, error) {
	r.Next()

	objects := []types.Object{}

	for {
		if r.pos >= len(r.tokens) {
			return nil, errors.New("EOF")
		}

		if r.Peek()[0] == ending {
			r.Next()
			break
		}

		object, err := r.ReadForm()
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	return objects, nil
}

func (r *Reader) ReadAtom() (types.Object, error) {
	token := r.Next()

	if fn.IsStartsWith(token, '"') {
		if fn.IsSingleChar(token) || !fn.IsEndsWith(token, '"') {
			return nil, errors.New("EOF")
		}

		return types.NewString(strings.Trim(token, `"`)), nil
	}

	if fn.IsStartsWith(token, ':') {
		return types.NewKeyword(token[1:]), nil
	}

	number, err := strconv.Atoi(token)
	if err != nil {
		return types.NewSymbol(token), nil
	}

	return types.NewNumber(number), nil
}

func tokenize(s string) []string {
	pattern := `[\s,]*(~@|[\[\]{}()'` + "`" + `~^@]|"(?:\\.|[^\\"])*"?|;.*|[^\s\[\]{}('"` + "`" + `,;)]*)`
	r := regexp.MustCompile(pattern)
	return fn.TrimSpacesFromEach(fn.TrimCommasFromEach(r.FindAllString(s, -1)))
}

func ReadString(s string) (types.Object, error) {
	tokens := tokenize(s)
	// fmt.Println("TOKENS: <" + strings.Join(tokens, "|") + ">")
	r := Reader{tokens: tokens}
	return r.ReadForm()
}
