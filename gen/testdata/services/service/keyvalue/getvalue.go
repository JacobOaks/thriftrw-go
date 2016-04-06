// Code generated by thriftrw

package keyvalue

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/exceptions"
	"github.com/thriftrw/thriftrw-go/gen/testdata/services"
	"github.com/thriftrw/thriftrw-go/gen/testdata/unions"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type GetValueArgs struct {
	Key *services.Key `json:"key,omitempty"`
}

func (v *GetValueArgs) ToWire() wire.Value {
	var fields [1]wire.Field
	i := 0
	if v.Key != nil {
		fields[i] = wire.Field{ID: 1, Value: v.Key.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func (v *GetValueArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x services.Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *GetValueArgs) String() string {
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("GetValueArgs{%v}", strings.Join(fields[:i], ", "))
}

type GetValueResult struct {
	Success      *unions.ArbitraryValue            `json:"success,omitempty"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

func (v *GetValueResult) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	if v.Success != nil {
		fields[i] = wire.Field{ID: 0, Value: v.Success.ToWire()}
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = wire.Field{ID: 1, Value: v.DoesNotExist.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func (v *GetValueResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ArbitraryValue_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *GetValueResult) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("GetValueResult{%v}", strings.Join(fields[:i], ", "))
}

var GetValueHelper = struct {
	IsException    func(error) bool
	Args           func(key *services.Key) *GetValueArgs
	WrapResponse   func(*unions.ArbitraryValue, error) (*GetValueResult, error)
	UnwrapResponse func(*GetValueResult) (*unions.ArbitraryValue, error)
}{}

func init() {
	GetValueHelper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}
	GetValueHelper.Args = func(key *services.Key) *GetValueArgs {
		return &GetValueArgs{Key: key}
	}
	GetValueHelper.WrapResponse = func(success *unions.ArbitraryValue, err error) (*GetValueResult, error) {
		if err == nil {
			return &GetValueResult{Success: success}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for GetValueResult.DoesNotExist")
			}
			return &GetValueResult{DoesNotExist: e}, nil
		}
		return nil, err
	}
	GetValueHelper.UnwrapResponse = func(result *GetValueResult) (success *unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
