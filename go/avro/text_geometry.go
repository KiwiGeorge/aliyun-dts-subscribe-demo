// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TextGeometry struct {
	Type string `json:"type"`

	Value string `json:"value"`
}

const TextGeometryAvroCRC64Fingerprint = "\xa0e\x94\x95=\xfb\xb2r"

func NewTextGeometry() TextGeometry {
	r := TextGeometry{}
	return r
}

func DeserializeTextGeometry(r io.Reader) (TextGeometry, error) {
	t := NewTextGeometry()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTextGeometryFromSchema(r io.Reader, schema string) (TextGeometry, error) {
	t := NewTextGeometry()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTextGeometry(r TextGeometry, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r TextGeometry) Serialize(w io.Writer) error {
	return writeTextGeometry(r, w)
}

func (r TextGeometry) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"com.alibaba.dts.formats.avro.TextGeometry\",\"type\":\"record\"}"
}

func (r TextGeometry) SchemaName() string {
	return "com.alibaba.dts.formats.avro.TextGeometry"
}

func (_ TextGeometry) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TextGeometry) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TextGeometry) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TextGeometry) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TextGeometry) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TextGeometry) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TextGeometry) SetString(v string)   { panic("Unsupported operation") }
func (_ TextGeometry) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TextGeometry) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Type}

		return w

	case 1:
		w := types.String{Target: &r.Value}

		return w

	}
	panic("Unknown field index")
}

func (r *TextGeometry) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *TextGeometry) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ TextGeometry) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TextGeometry) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TextGeometry) HintSize(int)                     { panic("Unsupported operation") }
func (_ TextGeometry) Finalize()                        {}

func (_ TextGeometry) AvroCRC64Fingerprint() []byte {
	return []byte(TextGeometryAvroCRC64Fingerprint)
}

func (r TextGeometry) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TextGeometry) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for value")
	}
	return nil
}
