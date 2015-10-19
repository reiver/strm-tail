package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// TAIL is a (string) constant that this Ender driver
	// is registered under.
	TAIL = "TAIL"

	defaultLimit = 5
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(TAIL, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Parse args.
	if 1 < len(args) {
//@TODO: Throw something better.
		panic("Too many parameters.")
	}

	var limit int = defaultLimit
	if 1 == len(args) {
		arg0 := args[0]
		switch n := arg0.(type) {
		case int:
			limit = n
		case int8:
			limit = int(n)
		case int16:
			limit = int(n)
		case int32:
			limit = int(n)
		case int64:
			limit = int(n)
		case uint:
			limit = int(n)
		case uint8:
			limit = int(n)
		case uint16:
			limit = int(n)
		case uint32:
			limit = int(n)
		case uint64:
			limit = int(n)
		default:
//@TODO: Throw something better.
			panic("Wrong type for limit.")
		}
	}

	// Execute.
	buffer := make([][]interface{}, limit)

	i :=0
	for datum := range src {
		buffer[i % limit] = datum

		i++
	}
	newLimit := limit+i
	if i < limit {
		newLimit -= (limit - i)
	}
	for ii:=i; ii<newLimit; ii++  {
		dst <- buffer[ii % limit]
	}
	close(dst)
}
