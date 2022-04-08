package main

import (
	"sort"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			//if !f.Generate {
			//	continue
			//}
			if strings.HasSuffix(string(f.GoImportPath), "/control") {
				generateFile(gen, f)
			}
		}
		return nil
	})
}

// generateFile generates a *.pb.go file enforcing field-order serialization.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_neofs.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-neofs. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P(`import "github.com/nspcc-dev/neofs-api-go/v2/util/proto"`)

	//for _, e := range file.Enums {
	//	g.P("type " + e.GoIdent.GoName + " int32")
	//	g.P("const (")
	//	for _, ev := range e.Values {
	//		g.P(ev.GoIdent.GoName, " = ", ev.Desc.Number())
	//	}
	//	g.P(")")
	//}
	for _, msg := range file.Messages {
		emitMessage(g, msg)
	}
	return g
}

func emitMessage(g *protogen.GeneratedFile, msg *protogen.Message) {
	for _, inner := range msg.Messages {
		emitMessage(g, inner)
	}

	fs := sortFields(msg.Fields)

	// StableSize implementation.
	g.P("// StableSize returns the size of x in protobuf format.")
	g.P("//")
	g.P("// Structures with the same field values have the same binary size.")
	g.P("func (x *", msg.GoIdent.GoName, ") StableSize() (size int) {")
	if len(fs) != 0 {
		for _, f := range fs {
			emitFieldSize(g, f)
		}
	}
	g.P("return size")
	g.P("}\n")

	// StableMarshal implementation.
	g.P("// StableMarshal marshals x in protobuf binary format with stable field order.")
	g.P("//")
	g.P("// If buffer length is less than x.StableSize(), new buffer is allocated.")
	g.P("//")
	g.P("// Returns any error encountered which did not allow writing the data completely.")
	g.P("// Otherwise, returns the buffer in which the data is written.")
	g.P("//")
	g.P("// Structures with the same field values have the same binary format.")
	g.P("func (x *", msg.GoIdent.GoName, ") StableMarshal(buf []byte) ([]byte, error) {")
	if len(fs) != 0 {
		g.P("if x == nil { return []byte{}, nil }")
		g.P("if buf == nil { buf = make([]byte, x.StableSize()) }")
		g.P("var err error")
		g.P("var offset, n int")
		g.P("_, _, _ = err, offset, n")
		for _, f := range fs {
			emitFieldMarshal(g, f)
		}
	}
	g.P("return buf, nil")
	g.P("}\n")

	if strings.HasSuffix(msg.GoIdent.GoName, "Request") || strings.HasSuffix(msg.GoIdent.GoName, "Response") {
		// SignedDataSize implementation (only for requests and responses).
		g.P("// ReadSignedData fills buf with signed data of x.")
		g.P("// If buffer length is less than x.SignedDataSize(), new buffer is allocated.")
		g.P("//")
		g.P("// Returns any error encountered which did not allow writing the data completely.")
		g.P("// Otherwise, returns the buffer in which the data is written.")
		g.P("//")
		g.P("// Structures with the same field values have the same signed data.")
		g.P("func (x *", msg.GoIdent.GoName, ") SignedDataSize() int {")
		g.P("return x.GetBody().StableSize()")
		g.P("}\n")

		// ReadSignedData implementation (only for requests and responses).
		g.P("// SignedDataSize returns size of the request signed data in bytes.")
		g.P("//")
		g.P("// Structures with the same field values have the same signed data size.")
		g.P("func (x *", msg.GoIdent.GoName, ") ReadSignedData(buf []byte) ([]byte, error) {")
		g.P("return x.GetBody().StableMarshal(buf)")
		g.P("}\n")

		// Signature setters and getters.
		g.P("func (x *", msg.GoIdent.GoName, ") SetSignature(sig *Signature) {")
		g.P("x.Signature = sig")
		g.P("}\n")
	}
}

func emitFieldSize(g *protogen.GeneratedFile, f *protogen.Field) {
	m := marshalers[f.Desc.Kind()]
	if m.Prefix == "" {
		g.P("// FIXME missing field marshaler: ", f.GoName, " of type ", f.Desc.Kind().String())
		g.P(`panic("unimplemented")`)
		return
	}

	name := castFieldName(f)
	if f.Oneof != nil {
		name = "x." + f.Oneof.GoName
		g.P("if inner, ok := ", name, ".(*", f.GoIdent.GoName, "); ok {")
		defer g.P("}")
		name = "inner." + f.GoName
	}

	switch {
	case f.Desc.IsList() && f.Desc.Kind() == protoreflect.MessageKind:
		g.P("for i := range ", name, "{")
		g.P("size += proto.NestedStructureSize(", f.Desc.Number(), ", ", name, "[i])")
		g.P("}")
	case f.Desc.IsList():
		g.P("size += proto.Repeated", m.Prefix, "Size(", f.Desc.Number(), ", ", name, ")")
	default:
		g.P("size += proto.", m.Prefix, "Size(", f.Desc.Number(), ", ", name, ")")
	}
}

func emitFieldMarshal(g *protogen.GeneratedFile, f *protogen.Field) {
	m := marshalers[f.Desc.Kind()]
	if m.Prefix == "" {
		g.P("// FIXME missing field marshaler: ", f.GoName, " of type ", f.Desc.Kind().String())
		g.P(`panic("unimplemented")`)
		return
	}

	name := castFieldName(f)
	if f.Oneof != nil {
		name = "x." + f.Oneof.GoName
		g.P("if inner, ok := ", name, ".(*", f.GoIdent.GoName, "); ok {")
		defer g.P("}")
		name = "inner." + f.GoName
	}

	prefix := m.Prefix
	if f.Desc.IsList() {
		prefix = "Repeated" + m.Prefix
	}
	switch {
	case f.Desc.IsList() && f.Desc.Kind() == protoreflect.MessageKind:
		g.P("for i := range ", name, "{")
		g.P("n, err = proto.NestedStructureMarshal(", f.Desc.Number(), ", buf[offset:], ", name, "[i])")
		g.P("if err != nil { return nil, err }")
		g.P("offset += n")
		g.P("}")
	case f.Desc.IsList():
		g.P("offset += proto.Repeated", m.Prefix, "Marshal(", f.Desc.Number(), ", buf[offset:], ", name, ")")
	default:
		if m.CanFail {
			g.P("n, err = proto.", prefix, "Marshal(", f.Desc.Number(), ", buf[offset:], ", name, ")")
			g.P("if err != nil { return nil, err }")
			g.P("offset += n")
		} else {
			g.P("offset += proto.", prefix, "Marshal(", f.Desc.Number(), ", buf[offset:], ", name, ")")
		}
	}
}

func castFieldName(f *protogen.Field) string {
	name := "x." + f.GoName
	if f.Desc.Kind() != protoreflect.EnumKind {
		return name
	}
	return "int32(" + name + ")"
}

type marshalerDesc struct {
	Prefix  string
	CanFail bool
}

// Unused kinds are commented.
var marshalers = map[protoreflect.Kind]marshalerDesc{
	protoreflect.BoolKind: {Prefix: "Bool"},
	protoreflect.EnumKind: {Prefix: "Enum"},
	//protoreflect.Int32Kind:    "",
	//protoreflect.Sint32Kind:   "",
	protoreflect.Uint32Kind: {Prefix: "UInt32"},
	protoreflect.Int64Kind:  {Prefix: "Int64"},
	//protoreflect.Sint64Kind:   "",
	protoreflect.Uint64Kind: {Prefix: "UInt64"},
	//protoreflect.Sfixed32Kind: "",
	protoreflect.Fixed32Kind: {Prefix: "Fixed32"},
	//protoreflect.FloatKind:    "",
	//protoreflect.Sfixed64Kind: "",
	protoreflect.Fixed64Kind: {Prefix: "Fixed64"},
	protoreflect.DoubleKind:  {Prefix: "Float64"},
	protoreflect.StringKind:  {Prefix: "String"},
	protoreflect.BytesKind:   {Prefix: "Bytes"},
	protoreflect.MessageKind: {Prefix: "NestedStructure", CanFail: true},
	//protoreflect.GroupKind:    "",
}

func sortFields(fs []*protogen.Field) []*protogen.Field {
	res := make([]*protogen.Field, len(fs))
	copy(res, fs)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Desc.Number() < res[j].Desc.Number()
	})
	return res
}
