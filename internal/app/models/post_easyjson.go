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

func easyjson5a72dc82DecodeDBForumInternalAppModels(in *jlexer.Lexer, out *PostList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostList, 0, 0)
			} else {
				*out = PostList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Post
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeDBForumInternalAppModels(out *jwriter.Writer, in PostList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeDBForumInternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeDBForumInternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeDBForumInternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeDBForumInternalAppModels(l, v)
}
func easyjson5a72dc82DecodeDBForumInternalAppModels1(in *jlexer.Lexer, out *PostInfo) {
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
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(Post)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(User)
				}
				easyjson5a72dc82DecodeDBForumInternalAppModels2(in, out.Author)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(Thread)
				}
				easyjson5a72dc82DecodeDBForumInternalAppModels3(in, out.Thread)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(Forum)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
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
func easyjson5a72dc82EncodeDBForumInternalAppModels1(out *jwriter.Writer, in PostInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Post != nil {
		const prefix string = ",\"post\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Post).MarshalEasyJSON(out)
	}
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson5a72dc82EncodeDBForumInternalAppModels2(out, *in.Author)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson5a72dc82EncodeDBForumInternalAppModels3(out, *in.Thread)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeDBForumInternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeDBForumInternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeDBForumInternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeDBForumInternalAppModels1(l, v)
}
func easyjson5a72dc82DecodeDBForumInternalAppModels3(in *jlexer.Lexer, out *Thread) {
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
			out.ID = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "forum":
			out.Forum = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "votes":
			out.Votes = int(in.Int())
		case "slug":
			out.Slug = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
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
func easyjson5a72dc82EncodeDBForumInternalAppModels3(out *jwriter.Writer, in Thread) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"votes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Votes))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson5a72dc82DecodeDBForumInternalAppModels2(in *jlexer.Lexer, out *User) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson5a72dc82EncodeDBForumInternalAppModels2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Nickname != "" {
		const prefix string = ",\"nickname\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	if in.Fullname != "" {
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.About))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}
func easyjson5a72dc82DecodeDBForumInternalAppModels4(in *jlexer.Lexer, out *Post) {
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
			out.ID = uint64(in.Uint64())
		case "parent":
			out.Parent = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = uint64(in.Uint64())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
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
func easyjson5a72dc82EncodeDBForumInternalAppModels4(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"parent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Parent))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Thread))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeDBForumInternalAppModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeDBForumInternalAppModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeDBForumInternalAppModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeDBForumInternalAppModels4(l, v)
}
