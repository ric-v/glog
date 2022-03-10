// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package glog

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

func easyjsonC5e9d7b4EncodeGithubComRicVGlog(out *jwriter.Writer, in LoggerJson) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.String(string(in.Time))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix)
		out.String(string(in.File))
	}
	{
		const prefix string = ",\"line\":"
		out.RawString(prefix)
		out.String(string(in.Line))
	}
	{
		const prefix string = ",\"msg\":"
		out.RawString(prefix)
		if in.Msg == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Msg {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoggerJson) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5e9d7b4EncodeGithubComRicVGlog(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
