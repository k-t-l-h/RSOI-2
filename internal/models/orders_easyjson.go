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

func easyjson9d6b4be7DecodeLab2MicroservicesKTLHInternalModels(in *jlexer.Lexer, out *Orders) {
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
		case "ID":
			out.ID = int(in.Int())
		case "itemUid":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.ItemUuid).UnmarshalText(data))
			}
		case "orderDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OrderDate).UnmarshalJSON(data))
			}
		case "orderUid":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.OrderUuid).UnmarshalText(data))
			}
		case "status":
			out.Status = string(in.String())
		case "UserUuid":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.UserUuid).UnmarshalText(data))
			}
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
func easyjson9d6b4be7EncodeLab2MicroservicesKTLHInternalModels(out *jwriter.Writer, in Orders) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"itemUid\":"
		out.RawString(prefix)
		out.RawText((in.ItemUuid).MarshalText())
	}
	{
		const prefix string = ",\"orderDate\":"
		out.RawString(prefix)
		out.Raw((in.OrderDate).MarshalJSON())
	}
	{
		const prefix string = ",\"orderUid\":"
		out.RawString(prefix)
		out.RawText((in.OrderUuid).MarshalText())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"UserUuid\":"
		out.RawString(prefix)
		out.RawText((in.UserUuid).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Orders) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9d6b4be7EncodeLab2MicroservicesKTLHInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Orders) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9d6b4be7EncodeLab2MicroservicesKTLHInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Orders) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9d6b4be7DecodeLab2MicroservicesKTLHInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Orders) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9d6b4be7DecodeLab2MicroservicesKTLHInternalModels(l, v)
}
