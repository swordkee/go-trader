package protocol

import "testing"
import (
	"bytes"
	. "github.com/robaho/go-trader/pkg/common"
	"reflect"
)

func TestEncodeDecodeBook(t *testing.T) {

	instrument := NewInstrument(12345, "IBM")
	IMap.Put(instrument)

	book := Book{Instrument: instrument, Sequence: 123456789}
	book.Bids = []BookLevel{{NewDecimal("99.4567"), NewDecimal("100")}}
	book.Asks = []BookLevel{{NewDecimal("100.4567"), NewDecimal("120")}}

	buf := new(bytes.Buffer)
	encodeBook(buf, &book)

	book2 := decodeBook(buf, instrument)

	if !reflect.DeepEqual(book, *book2) {
		t.Error("books do not match", &book, book2)
	}
}
