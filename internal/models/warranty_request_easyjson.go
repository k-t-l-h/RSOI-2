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

func easyjsonDa965172DecodeLab2MicroservicesKTLHInternalModels(in *jlexer.Lexer, out *WarrantyRequest) {
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
		case "reason":
			out.Reason = string(in.String())
		case "available":
			out.Available = int(in.Int())
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
func easyjsonDa965172EncodeLab2MicroservicesKTLHInternalModels(out *jwriter.Writer, in WarrantyRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"reason\":"
		out.RawString(prefix[1:])
		out.String(string(in.Reason))
	}
	if in.Available != 0 {
		const prefix string = ",\"available\":"
		out.RawString(prefix)
		out.Int(int(in.Available))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WarrantyRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDa965172EncodeLab2MicroservicesKTLHInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WarrantyRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDa965172EncodeLab2MicroservicesKTLHInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WarrantyRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDa965172DecodeLab2MicroservicesKTLHInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WarrantyRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDa965172DecodeLab2MicroservicesKTLHInternalModels(l, v)
}
