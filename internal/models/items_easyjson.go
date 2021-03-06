// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD8ded6e2DecodeLab2MicroservicesKTLHInternalModels(in *jlexer.Lexer, out *Items) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "available":
			out.Available = int(in.Int())
		case "model":
			out.Model = string(in.String())
		case "size":
			out.Size = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD8ded6e2EncodeLab2MicroservicesKTLHInternalModels(out *jwriter.Writer, in Items) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"available\":"
		out.RawString(prefix)
		out.Int(int(in.Available))
	}
	{
		const prefix string = ",\"model\":"
		out.RawString(prefix)
		out.String(string(in.Model))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.String(string(in.Size))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Items) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD8ded6e2EncodeLab2MicroservicesKTLHInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Items) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD8ded6e2EncodeLab2MicroservicesKTLHInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Items) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD8ded6e2DecodeLab2MicroservicesKTLHInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Items) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD8ded6e2DecodeLab2MicroservicesKTLHInternalModels(l, v)
}
