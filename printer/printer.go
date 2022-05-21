package printer

import (
	"strconv"
	"strings"

	"github.com/cyucelen/golis/fn"
	"github.com/cyucelen/golis/types"
)

func PrintString(object types.Object) string {
	switch o := object.(type) {
	case types.Number:
		return strconv.Itoa(o.Value)
	case types.String:
		return `"` + o.Value + `"`
	case types.Keyword:
		return ":" + o.Value
	case types.Symbol:
		return o.Name
	case *types.List:
		ss := fn.MapObjectToString(o.Values(), PrintString)
		return "(" + strings.Join(ss, " ") + ")"
	case *types.Vector:
		ss := fn.MapObjectToString(o.Values(), PrintString)
		return "[" + strings.Join(ss, " ") + "]"
	case *types.HashMap:
		ss := fn.MapObjectToString(o.KVList(), PrintString)
		return "{" + strings.Join(ss, " ") + "}"
	default:
		return "Undefined Type"
	}
}
