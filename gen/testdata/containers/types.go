// Code generated by thriftrw

package containers

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/enums"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type EnumContainers struct {
	ListOfEnums []enums.EnumDefault                     `json:"listOfEnums"`
	SetOfEnums  map[enums.EnumWithValues]struct{}       `json:"setOfEnums"`
	MapOfEnums  map[enums.EnumWithDuplicateValues]int32 `json:"mapOfEnums"`
}

type _List_EnumDefault_ValueList []enums.EnumDefault

func (v _List_EnumDefault_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_EnumDefault_ValueList) Close() {
}

type _Set_EnumWithValues_ValueList map[enums.EnumWithValues]struct{}

func (v _Set_EnumWithValues_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_EnumWithValues_ValueList) Close() {
}

type _Map_EnumWithDuplicateValues_I32_MapItemList map[enums.EnumWithDuplicateValues]int32

func (m _Map_EnumWithDuplicateValues_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: k.ToWire(), Value: wire.NewValueI32(v)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_EnumWithDuplicateValues_I32_MapItemList) Close() {
}

func (v *EnumContainers) ToWire() wire.Value {
	var fields [3]wire.Field
	i := 0
	if v.ListOfEnums != nil {
		fields[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TI32, Size: len(v.ListOfEnums), Items: _List_EnumDefault_ValueList(v.ListOfEnums)})}
		i++
	}
	if v.SetOfEnums != nil {
		fields[i] = wire.Field{ID: 2, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.SetOfEnums), Items: _Set_EnumWithValues_ValueList(v.SetOfEnums)})}
		i++
	}
	if v.MapOfEnums != nil {
		fields[i] = wire.Field{ID: 3, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI32, ValueType: wire.TI32, Size: len(v.MapOfEnums), Items: _Map_EnumWithDuplicateValues_I32_MapItemList(v.MapOfEnums)})}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func _EnumDefault_Read(w wire.Value) (enums.EnumDefault, error) {
	var v enums.EnumDefault
	err := v.FromWire(w)
	return v, err
}

func _List_EnumDefault_Read(l wire.List) ([]enums.EnumDefault, error) {
	if l.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make([]enums.EnumDefault, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _EnumDefault_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func _EnumWithValues_Read(w wire.Value) (enums.EnumWithValues, error) {
	var v enums.EnumWithValues
	err := v.FromWire(w)
	return v, err
}

func _Set_EnumWithValues_Read(s wire.Set) (map[enums.EnumWithValues]struct{}, error) {
	if s.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[enums.EnumWithValues]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := _EnumWithValues_Read(x)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

func _EnumWithDuplicateValues_Read(w wire.Value) (enums.EnumWithDuplicateValues, error) {
	var v enums.EnumWithDuplicateValues
	err := v.FromWire(w)
	return v, err
}

func _Map_EnumWithDuplicateValues_I32_Read(m wire.Map) (map[enums.EnumWithDuplicateValues]int32, error) {
	if m.KeyType != wire.TI32 {
		return nil, nil
	}
	if m.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[enums.EnumWithDuplicateValues]int32, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := _EnumWithDuplicateValues_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := x.Value.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}

func (v *EnumContainers) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.ListOfEnums, err = _List_EnumDefault_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TSet {
				v.SetOfEnums, err = _Set_EnumWithValues_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TMap {
				v.MapOfEnums, err = _Map_EnumWithDuplicateValues_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *EnumContainers) String() string {
	var fields [3]string
	i := 0
	if v.ListOfEnums != nil {
		fields[i] = fmt.Sprintf("ListOfEnums: %v", v.ListOfEnums)
		i++
	}
	if v.SetOfEnums != nil {
		fields[i] = fmt.Sprintf("SetOfEnums: %v", v.SetOfEnums)
		i++
	}
	if v.MapOfEnums != nil {
		fields[i] = fmt.Sprintf("MapOfEnums: %v", v.MapOfEnums)
		i++
	}
	return fmt.Sprintf("EnumContainers{%v}", strings.Join(fields[:i], ", "))
}

type PrimitiveContainers struct {
	ListOfBinary      [][]byte            `json:"listOfBinary"`
	ListOfInts        []int64             `json:"listOfInts"`
	SetOfStrings      map[string]struct{} `json:"setOfStrings"`
	SetOfBytes        map[int8]struct{}   `json:"setOfBytes"`
	MapOfIntToString  map[int32]string    `json:"mapOfIntToString"`
	MapOfStringToBool map[string]bool     `json:"mapOfStringToBool"`
}

type _List_Binary_ValueList [][]byte

func (v _List_Binary_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueBinary(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Binary_ValueList) Close() {
}

type _List_I64_ValueList []int64

func (v _List_I64_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueI64(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_I64_ValueList) Close() {
}

type _Set_String_ValueList map[string]struct{}

func (v _Set_String_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueString(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_String_ValueList) Close() {
}

type _Set_Byte_ValueList map[int8]struct{}

func (v _Set_Byte_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueI8(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_Byte_ValueList) Close() {
}

type _Map_I32_String_MapItemList map[int32]string

func (m _Map_I32_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueI32(k), Value: wire.NewValueString(v)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_I32_String_MapItemList) Close() {
}

type _Map_String_Bool_MapItemList map[string]bool

func (m _Map_String_Bool_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueString(k), Value: wire.NewValueBool(v)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_Bool_MapItemList) Close() {
}

func (v *PrimitiveContainers) ToWire() wire.Value {
	var fields [6]wire.Field
	i := 0
	if v.ListOfBinary != nil {
		fields[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TBinary, Size: len(v.ListOfBinary), Items: _List_Binary_ValueList(v.ListOfBinary)})}
		i++
	}
	if v.ListOfInts != nil {
		fields[i] = wire.Field{ID: 2, Value: wire.NewValueList(wire.List{ValueType: wire.TI64, Size: len(v.ListOfInts), Items: _List_I64_ValueList(v.ListOfInts)})}
		i++
	}
	if v.SetOfStrings != nil {
		fields[i] = wire.Field{ID: 3, Value: wire.NewValueSet(wire.Set{ValueType: wire.TBinary, Size: len(v.SetOfStrings), Items: _Set_String_ValueList(v.SetOfStrings)})}
		i++
	}
	if v.SetOfBytes != nil {
		fields[i] = wire.Field{ID: 4, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI8, Size: len(v.SetOfBytes), Items: _Set_Byte_ValueList(v.SetOfBytes)})}
		i++
	}
	if v.MapOfIntToString != nil {
		fields[i] = wire.Field{ID: 5, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI32, ValueType: wire.TBinary, Size: len(v.MapOfIntToString), Items: _Map_I32_String_MapItemList(v.MapOfIntToString)})}
		i++
	}
	if v.MapOfStringToBool != nil {
		fields[i] = wire.Field{ID: 6, Value: wire.NewValueMap(wire.Map{KeyType: wire.TBinary, ValueType: wire.TBool, Size: len(v.MapOfStringToBool), Items: _Map_String_Bool_MapItemList(v.MapOfStringToBool)})}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func _List_Binary_Read(l wire.List) ([][]byte, error) {
	if l.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make([][]byte, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetBinary(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func _List_I64_Read(l wire.List) ([]int64, error) {
	if l.ValueType != wire.TI64 {
		return nil, nil
	}
	o := make([]int64, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI64(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func _Set_String_Read(s wire.Set) (map[string]struct{}, error) {
	if s.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make(map[string]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

func _Set_Byte_Read(s wire.Set) (map[int8]struct{}, error) {
	if s.ValueType != wire.TI8 {
		return nil, nil
	}
	o := make(map[int8]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI8(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

func _Map_I32_String_Read(m wire.Map) (map[int32]string, error) {
	if m.KeyType != wire.TI32 {
		return nil, nil
	}
	if m.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make(map[int32]string, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI32(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}

func _Map_String_Bool_Read(m wire.Map) (map[string]bool, error) {
	if m.KeyType != wire.TBinary {
		return nil, nil
	}
	if m.ValueType != wire.TBool {
		return nil, nil
	}
	o := make(map[string]bool, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetBool(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}

func (v *PrimitiveContainers) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.ListOfBinary, err = _List_Binary_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TList {
				v.ListOfInts, err = _List_I64_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TSet {
				v.SetOfStrings, err = _Set_String_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TSet {
				v.SetOfBytes, err = _Set_Byte_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TMap {
				v.MapOfIntToString, err = _Map_I32_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 6:
			if field.Value.Type() == wire.TMap {
				v.MapOfStringToBool, err = _Map_String_Bool_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *PrimitiveContainers) String() string {
	var fields [6]string
	i := 0
	if v.ListOfBinary != nil {
		fields[i] = fmt.Sprintf("ListOfBinary: %v", v.ListOfBinary)
		i++
	}
	if v.ListOfInts != nil {
		fields[i] = fmt.Sprintf("ListOfInts: %v", v.ListOfInts)
		i++
	}
	if v.SetOfStrings != nil {
		fields[i] = fmt.Sprintf("SetOfStrings: %v", v.SetOfStrings)
		i++
	}
	if v.SetOfBytes != nil {
		fields[i] = fmt.Sprintf("SetOfBytes: %v", v.SetOfBytes)
		i++
	}
	if v.MapOfIntToString != nil {
		fields[i] = fmt.Sprintf("MapOfIntToString: %v", v.MapOfIntToString)
		i++
	}
	if v.MapOfStringToBool != nil {
		fields[i] = fmt.Sprintf("MapOfStringToBool: %v", v.MapOfStringToBool)
		i++
	}
	return fmt.Sprintf("PrimitiveContainers{%v}", strings.Join(fields[:i], ", "))
}

type PrimitiveContainersRequired struct {
	ListOfStrings      []string           `json:"listOfStrings"`
	SetOfInts          map[int32]struct{} `json:"setOfInts"`
	MapOfIntsToDoubles map[int64]float64  `json:"mapOfIntsToDoubles"`
}

type _List_String_ValueList []string

func (v _List_String_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueString(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_String_ValueList) Close() {
}

type _Set_I32_ValueList map[int32]struct{}

func (v _Set_I32_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueI32(x))
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_I32_ValueList) Close() {
}

type _Map_I64_Double_MapItemList map[int64]float64

func (m _Map_I64_Double_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueI64(k), Value: wire.NewValueDouble(v)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_I64_Double_MapItemList) Close() {
}

func (v *PrimitiveContainersRequired) ToWire() wire.Value {
	var fields [3]wire.Field
	i := 0
	fields[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TBinary, Size: len(v.ListOfStrings), Items: _List_String_ValueList(v.ListOfStrings)})}
	i++
	fields[i] = wire.Field{ID: 2, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.SetOfInts), Items: _Set_I32_ValueList(v.SetOfInts)})}
	i++
	fields[i] = wire.Field{ID: 3, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI64, ValueType: wire.TDouble, Size: len(v.MapOfIntsToDoubles), Items: _Map_I64_Double_MapItemList(v.MapOfIntsToDoubles)})}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}

func _List_String_Read(l wire.List) ([]string, error) {
	if l.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make([]string, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func _Set_I32_Read(s wire.Set) (map[int32]struct{}, error) {
	if s.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[int32]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

func _Map_I64_Double_Read(m wire.Map) (map[int64]float64, error) {
	if m.KeyType != wire.TI64 {
		return nil, nil
	}
	if m.ValueType != wire.TDouble {
		return nil, nil
	}
	o := make(map[int64]float64, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI64(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetDouble(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}

func (v *PrimitiveContainersRequired) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.ListOfStrings, err = _List_String_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TSet {
				v.SetOfInts, err = _Set_I32_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TMap {
				v.MapOfIntsToDoubles, err = _Map_I64_Double_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *PrimitiveContainersRequired) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("ListOfStrings: %v", v.ListOfStrings)
	i++
	fields[i] = fmt.Sprintf("SetOfInts: %v", v.SetOfInts)
	i++
	fields[i] = fmt.Sprintf("MapOfIntsToDoubles: %v", v.MapOfIntsToDoubles)
	i++
	return fmt.Sprintf("PrimitiveContainersRequired{%v}", strings.Join(fields[:i], ", "))
}
