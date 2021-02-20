package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type indentWriter struct {
	dat    strings.Builder
	Indent int
}

func (i *indentWriter) Add(format string, v ...interface{}) {
	i.dat.WriteString(strings.Repeat("\t", i.Indent))
	i.dat.WriteString(fmt.Sprintf(format, v...))
	i.dat.WriteRune('\n')
}

func (i *indentWriter) Data() string {
	return i.dat.String()
}

var wellKnown = map[descriptorpb.FieldDescriptorProto_Type]string{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   "double",
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    "float",
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    "qint64",
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   "quint64",
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    "qint32",
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  "quint64",
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  "quint32",
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     "bool",
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   "QString",
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    "QByteArray",
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "quint32",
	descriptorpb.FieldDescriptorProto_TYPE_ENUM:     "",
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "qint32",
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "qint64",
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "qint32",
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "qint64",
}

func convertType(s *descriptorpb.FieldDescriptorProto) string {
	if s.TypeName != nil {
		return strings.Join(strings.Split(s.GetTypeName(), "."), "::")
	}
	if v, ok := wellKnown[s.GetType()]; ok {
		return v
	}
	panic("unhandled")
}

// Indent indents a block of text with an indent string
func Indent(text, indent string) string {
	first := false
	if text[len(text)-1:] == "\n" {
		result := ""
		for _, j := range strings.Split(text[:len(text)-1], "\n") {
			if !first {
				first = true
				result += j + "\n"
			} else {
				result += indent + j + "\n"
			}
		}
		return result
	}
	result := ""
	for _, j := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		if !first {
			first = true
			result += j + "\n"
		} else {
			result += indent + j + "\n"
		}
	}
	return result[:len(result)-1]
}

func readTag(iw *indentWriter, len, idx string) {
	iw.Add(`auto [%s, %s] = [&]() -> std::pair<quint64,quint64> {`, len, idx)
	iw.Indent++
	{
		iw.Add(`auto num =`)
		iw.Indent++
		iw.Add(Indent(readVarIntNumber, strings.Repeat("\t", iw.Indent)))
		iw.Indent--
		iw.Add(`;`)
		iw.Add(``)
		iw.Add(`auto wireType = num & 0b111;`)
		iw.Add(`auto index = num ^ 0b111;`)
		iw.Add(`return {len, wireType};`)
	}
	iw.Indent--
	iw.Add(`}();`)
}

const readVarIntNumber = `[&]() -> quint64 {
	currentBuffer.clear();
	currentBitTwiddling.clear();

readLoop:
	if (!b.getChar(&current)) {
		unexpectedEOF = true;
		return 0;
	}
	currentBuffer.append(current);
	if (current & (1 << 7)) {
		goto readLoop;
	}

	auto it = currentBuffer.constEnd();
	while (it != currentBuffer.constBegin()) {
		it--;
		currentBitTwiddling.push_back((*it >> 0) & 1);
		currentBitTwiddling.push_back((*it >> 1) & 1);
		currentBitTwiddling.push_back((*it >> 2) & 1);
		currentBitTwiddling.push_back((*it >> 3) & 1);
		currentBitTwiddling.push_back((*it >> 4) & 1);
		currentBitTwiddling.push_back((*it >> 5) & 1);
		currentBitTwiddling.push_back((*it >> 6) & 1);
	}

	quint64 ret = 0;
	auto i = 0;
	for (const auto& item : currentBitTwiddling) {
		if (item) {
			ret |= 1UL << i;
		} else {
			ret &= ~(1UL << i);
		}
		i++;
	}

	return ret;
}()`

func genMessages(iw *indentWriter, msgs []*descriptorpb.DescriptorProto) {
	for _, msg := range msgs {
		iw.Add(`struct %s {`, msg.GetName())
		iw.Indent++

		genMessages(iw, msg.GetNestedType())

		for _, mem := range msg.Field {
			if mem.OneofIndex != nil {
				continue
			}
			iw.Add(`%s %s; // %d`, convertType(mem), mem.GetName(), mem.GetNumber())
		}

		for i, mem := range msg.OneofDecl {
			iw.Add(`std::variant<`)
			iw.Indent++
			first := true
			for _, id := range msg.Field {
				if id.OneofIndex == nil || id.GetOneofIndex() != int32(i) {
					continue
				}
				if first {
					first = false
					iw.Add(`%s // %d`, convertType(id), *id.Number)
				} else {
					iw.Add(`, %s // %d`, convertType(id), *id.Number)
				}
			}
			iw.Indent--
			iw.Add(`> %s;`, mem.GetName())
		}

		iw.Add(`static %s fromBytes(const QByteArray& ba) {`, msg.GetName())
		iw.Indent++
		{
			iw.Add(`QBuffer b(&ba);`)

			iw.Add(`QList<char> currentBuffer;`)
			iw.Add(`std::vector<bool> currentBitTwiddling;`)
			iw.Add(`char current = 0;`)
			iw.Add(`bool unexpectedEOF = false;`)

			iw.Add(`while (true) {`)
			iw.Indent++
			{
				readTag(iw, "idx", "kind")
				iw.Add(`if (unexpectedEOF) return {QStringLiteral("unexpected EOF");`)
				iw.Add(`currentBuffer.clear();`)
				iw.Add(`currentBitTwiddling.clear();`)

				iw.Add(`switch (kind) {`)
				// varint
				iw.Add(`case 0: {`)
				iw.Indent++
				{

				}
				iw.Indent--
				iw.Add(`}`)
			}
			iw.Indent--
			iw.Add(`}`)

		}
		iw.Indent--
		iw.Add(`}`)

		iw.Add(`QByteArray intoBytes() {`)
		iw.Indent++
		{

		}
		iw.Indent--
		iw.Add(`}`)

		iw.Indent--
		iw.Add(`}; // %s`, msg.GetName())
	}
}

func GenerateQtCxxProto(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.MessageType) == 0 {
			continue
		}

		ns := strings.Split(f.GetPackage(), ".")

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = convertCxxProto(f.GetName(), "hrpc.proto", "hpp")

		iw := indentWriter{}

		defer func() {
			file.Content = new(string)
			*file.Content = iw.Data()
			r.File = append(r.File, file)
		}()

		iw.Add(`#include <QBuffer>`)
		iw.Add(`#include <QByteArray>`)
		iw.Add(`#include <QString>`)
		iw.Add(`#include <QtGlobal>`)
		iw.Add(``)
		iw.Add(`#include <variant>`)
		iw.Add(`#include <utility>`)
		iw.Add(``)

		for _, n := range ns {
			iw.Add(`namespace %s {`, n)
			iw.Indent++
		}

		genMessages(&iw, f.GetMessageType())

		for i := len(ns) - 1; i >= 0; i-- {
			iw.Indent--
			iw.Add(`} // %s`, ns[i])
		}
	}
	return
}
