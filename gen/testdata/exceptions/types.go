// Code generated by thriftrw

package exceptions

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type DoesNotExistException struct{ Key string }

func (v *DoesNotExistException) ToWire() wire.Value {
	var fields [1]wire.Field
	i := 0
	fields[i] = wire.Field{ID: 1, Value: wire.NewValueString(v.Key)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func (v *DoesNotExistException) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *DoesNotExistException) String() string {
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	return fmt.Sprintf("DoesNotExistException{%v}", strings.Join(fields[:i], ", "))
}
func (v *DoesNotExistException) Error() string {
	return v.String()
}

type EmptyException struct{}

func (v *EmptyException) ToWire() wire.Value {
	var fields [0]wire.Field
	i := 0
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func (v *EmptyException) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}
func (v *EmptyException) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("EmptyException{%v}", strings.Join(fields[:i], ", "))
}
func (v *EmptyException) Error() string {
	return v.String()
}
