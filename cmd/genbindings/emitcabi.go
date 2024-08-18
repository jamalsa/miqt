package main

import (
	"fmt"
	"sort"
	"strings"
)

func (p CppParameter) RenderTypeCpp() string {
	ret := p.ParameterType
	switch p.ParameterType {
	case "uchar":
		ret = "unsigned char"
	case "uint":
		ret = "unsigned int"
	case "ulong":
		ret = "unsigned long"
	case "qint8":
		ret = "int8_t"
	case "quint8":
		ret = "uint8_t"
	case "qint16", "short":
		ret = "int16_t"
	case "quint16", "ushort", "unsigned short":
		ret = "uint16_t"
	case "qint32":
		ret = "int32_t"
	case "quint32":
		ret = "uint32_t"
	case "qlonglong", "qint64":
		ret = "int64_t"
	case "qulonglong", "quint64":
		ret = "uint64_t"
	case "qfloat16":
		ret = "_Float16" // No idea where this typedef comes from, but it exists
	case "qsizetype":
		ret = "size_t"
	case "QRgb":
		ret = "unsigned int"
	}

	if p.Pointer || p.ByRef {
		ret += "*"
	}

	return ret // ignore const
}

func emitReturnTypeCabi(p CppParameter) string {
	if p.ParameterType == "QString" {
		return "void" // Will be handled separately

	} else if _, ok := p.QListOf(); ok {
		return "void" // Will be handled separately

	} else if (p.Pointer || p.ByRef) && p.QtClassType() {
		return p.ParameterType + "*" // CABI type

	} else if p.QtClassType() && !p.Pointer {
		// Even if C++ returns by value, CABI is returning a heap copy (new'd, not malloc'd)
		return p.ParameterType + "*" // CABI type
		// return "void" // Handled separately with an _out pointer

	} else {
		return p.RenderTypeCpp()
	}
}

func emitParametersCabi(m CppMethod, selfType string) string {
	tmp := make([]string, 0, len(m.Parameters)+1)

	if !m.IsStatic && selfType != "" {
		tmp = append(tmp, selfType+" self")
	}

	for _, p := range m.Parameters {
		if p.ParameterType == "QString" {
			// The Go code has called this with two arguments: char* and len
			// Declare that we take two parameters
			tmp = append(tmp, "const char* "+p.ParameterName+", size_t "+p.ParameterName+"_Strlen")

		} else if t, ok := p.QListOf(); ok {

			if t.ParameterType == "QString" {
				// Combov
				tmp = append(tmp, "char** "+p.ParameterName+", int64_t* "+p.ParameterName+"_Lengths, size_t "+p.ParameterName+"_len")

			} else {
				// The Go code has called this with two arguments: T* and len
				// Declare that we take two parameters
				tmp = append(tmp, t.RenderTypeCpp()+"* "+p.ParameterName+", size_t "+p.ParameterName+"_len")
			}

		} else if (p.ByRef || p.Pointer) && p.QtClassType() {
			// Pointer to Qt type
			// Replace with taking our PQ typedef by value
			tmp = append(tmp, p.ParameterType+"* "+p.ParameterName)

		} else if p.QtClassType() {
			// Qt type passed by value
			// The CABI will unconditionally take these by pointer and dereference them
			// when passing to C++
			tmp = append(tmp, p.ParameterType+"* "+p.ParameterName)

		} else {
			// RenderTypeCpp renders both pointer+reference as pointers
			tmp = append(tmp, p.RenderTypeCpp()+" "+p.ParameterName)
		}
	}

	// If the return type is QString, we need to handle returns via extra CABI
	// parameters
	// Qt C++: memory is in QString RAII
	// CABI: memory is moved into C.malloc/C.free
	// Go: converted to native Go string
	if m.ReturnType.ParameterType == "QString" {
		tmp = append(tmp, "char** _out, size_t* _out_Strlen")

	} else if t, ok := m.ReturnType.QListOf(); ok {
		// +1 pointer indirection since it's a heap array
		// +1 pointer indirection for mutating remote parameter
		// = 3 for char*, 2 for most types

		// Maybe: +1 pointer indirection if we have to lift stack types to the heap

		if t.ParameterType == "QString" {
			// Combo
			tmp = append(tmp, "char*** _out, int64_t** _out_Lengths, size_t* _out_len")
		} else if t.QtClassType() && !t.Pointer {
			// QList<QByteArray> QByteArray::Split()
			// We need to pointer-ify each of the interior elements too
			tmp = append(tmp, t.RenderTypeCpp()+"*** _out, size_t* _out_len")
		} else {
			tmp = append(tmp, t.RenderTypeCpp()+"** _out, size_t* _out_len")
		}

	}

	return strings.Join(tmp, ", ")
}

func emitParametersCABI2CppForwarding(params []CppParameter) (preamble string, forwarding string) {
	tmp := make([]string, 0, len(params)+1)

	for _, p := range params {
		if p.ParameterType == "QString" {
			// The CABI has accepted two parameters - need to convert to one real QString
			// Create it on the stack
			preamble += "\tQString " + p.ParameterName + "_QString = QString::fromUtf8(" + p.ParameterName + ", " + p.ParameterName + "_Strlen);\n"
			tmp = append(tmp, p.ParameterName+"_QString")

		} else if listType, ok := p.QListOf(); ok {

			if listType.ParameterType == "QString" {

				// Combo (3 parameters)
				preamble += "\t" + p.ParameterType + " " + p.ParameterName + "_QList;\n"
				preamble += "\t" + p.ParameterName + "_QList.reserve(" + p.ParameterName + "_len);\n"
				preamble += "\tfor(size_t i = 0; i < " + p.ParameterName + "_len; ++i) {\n"
				preamble += "\t\t" + p.ParameterName + "_QList.push_back(QString::fromUtf8(" + p.ParameterName + "[i], " + p.ParameterName + "_Lengths[i]));\n"
				preamble += "\t}\n"
				tmp = append(tmp, p.ParameterName+"_QList")

			} else {

				// The CABI has accepted two parameters - need to convert to one real QList<>
				// Create it on the stack
				preamble += "\t" + p.ParameterType + " " + p.ParameterName + "_QList;\n"
				preamble += "\t" + p.ParameterName + "_QList.reserve(" + p.ParameterName + "_len);\n"
				preamble += "\tfor(size_t i = 0; i < " + p.ParameterName + "_len; ++i) {\n"
				preamble += "\t\t" + p.ParameterName + "_QList.push_back(" + p.ParameterName + "[i]);\n"
				preamble += "\t}\n"
				tmp = append(tmp, p.ParameterName+"_QList")
			}

		} else if p.IntType() {
			// Use the raw ParameterType to select an explicit integer overload
			// Don't use RenderTypeCpp() since it canonicalizes some int types for CABI
			castSrc := p.ParameterName
			castType := p.ParameterType
			if p.Pointer {
				castType += "*"
			}
			if p.ByRef { // e.g. QDataStream::operator>>() overloads
				castSrc = "*" + castSrc
				castType += "&" // believe it or not, this is legal

				if p.ParameterType == "qint64" || p.ParameterType == "quint64" {
					// QDataStream::operator>>()
					// CABI has these as int64_t* (long int) which fails a static_cast to qint64& (long long int&)
					// Hack a hard C-style cast
					castSrc = "(" + castType + ")(" + castSrc + ")"
				}
			}
			tmp = append(tmp, "static_cast<"+castType+">("+castSrc+")")

		} else if p.ByRef {
			// We changed RenderTypeCpp() to render this as a pointer
			// Need to dereference so we can pass as reference to the actual Qt C++ function
			//tmp = append(tmp, "*"+p.ParameterName)
			tmp = append(tmp, "*"+p.ParameterName)

		} else if p.QtClassType() && !p.Pointer {
			// CABI takes all Qt types by pointer, even if C++ wants them by value
			// Dereference the passed-in pointer
			tmp = append(tmp, "*"+p.ParameterName)

			// } else if p.QtClassType() && p.Pointer {
			// We need this static_cast<> anyway to convert from PQt (void*) to
			// the real Qt type
			// tmp = append(tmp, "static_cast<"+p.ParameterType+"*>("+p.ParameterName+")")

		} else {
			tmp = append(tmp, p.ParameterName)
		}
	}

	return preamble, strings.Join(tmp, ", ")
}

// getReferencedTypes finds all referenced Qt types in this file.
func getReferencedTypes(src *CppParsedHeader) []string {

	foundTypes := map[string]struct{}{}
	for _, c := range src.Classes {

		foundTypes[c.ClassName] = struct{}{}

		for _, ctor := range c.Ctors {
			for _, p := range ctor.Parameters {
				if p.QtClassType() {
					foundTypes[p.ParameterType] = struct{}{}
				}
				if t, ok := p.QListOf(); ok {
					foundTypes["QList"] = struct{}{} // FIXME or QVector?
					if t.QtClassType() {
						foundTypes[t.ParameterType] = struct{}{}
					}
				}
			}
		}
		for _, m := range c.Methods {
			for _, p := range m.Parameters {
				if p.QtClassType() {
					foundTypes[p.ParameterType] = struct{}{}
				}
				if t, ok := p.QListOf(); ok {
					foundTypes["QList"] = struct{}{} // FIXME or QVector?
					if t.QtClassType() {
						foundTypes[t.ParameterType] = struct{}{}
					}
				}
			}
			if m.ReturnType.QtClassType() {
				foundTypes[m.ReturnType.ParameterType] = struct{}{}
			}
			if t, ok := m.ReturnType.QListOf(); ok {
				foundTypes["QList"] = struct{}{} // FIXME or QVector?
				if t.QtClassType() {
					foundTypes[t.ParameterType] = struct{}{}
				}
			}
		}
	}

	// Some types (e.g. QRgb) are found but are typedefs, not classes
	for _, td := range src.Typedefs {
		delete(foundTypes, td.Alias)
	}

	// Convert to sorted list
	foundTypesList := make([]string, 0, len(foundTypes))
	for ft := range foundTypes {
		if strings.HasPrefix(ft, "QList<") || strings.HasPrefix(ft, "QVector<") {
			continue
		}
		if strings.HasSuffix(ft, "Private") { // qbrush.h finds QGradientPrivate
			continue
		}
		if ft == "QRgb" {
			continue
		}

		foundTypesList = append(foundTypesList, ft)
	}
	sort.Strings(foundTypesList)

	return foundTypesList
}

func emitBindingHeader(src *CppParsedHeader, filename string) (string, error) {
	ret := strings.Builder{}

	includeGuard := "GEN_" + strings.ToUpper(strings.Replace(filename, `.`, `_`, -1))

	ret.WriteString(`#ifndef ` + includeGuard + `
#define ` + includeGuard + `

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

`)

	foundTypesList := getReferencedTypes(src)

	ret.WriteString("#ifdef __cplusplus\n")

	for _, ft := range foundTypesList {
		if ft == "QList" || ft == "QString" { // These types are reprojected
			continue
		}
		ret.WriteString(`class ` + ft + ";\n")
	}

	ret.WriteString("#else\n")

	for _, ft := range foundTypesList {
		if ft == "QList" || ft == "QString" { // These types are reprojected
			continue
		}
		ret.WriteString(`typedef struct ` + ft + " " + ft + ";\n")
	}

	ret.WriteString("#endif\n")

	ret.WriteString("\n")

	for _, c := range src.Classes {

		for i, ctor := range c.Ctors {
			ret.WriteString(fmt.Sprintf("%s %s_new%s(%s);\n", c.ClassName+"*", c.ClassName, maybeSuffix(i), emitParametersCabi(ctor, "")))
		}

		for _, m := range c.Methods {
			ret.WriteString(fmt.Sprintf("%s %s_%s(%s);\n", emitReturnTypeCabi(m.ReturnType), c.ClassName, m.SafeMethodName(), emitParametersCabi(m, c.ClassName+"*")))
		}

		// delete
		ret.WriteString(fmt.Sprintf("void %s_Delete(%s* self);\n", c.ClassName, c.ClassName))

		ret.WriteString("\n")
	}

	ret.WriteString(
		`#ifdef __cplusplus
} /* extern C */
#endif 

#endif
`)
	return ret.String(), nil
}

func emitBindingCpp(src *CppParsedHeader, filename string) (string, error) {
	ret := strings.Builder{}

	ret.WriteString(`#include "gen_` + filename + `"
#include "` + filename + `"

`)

	for _, ref := range getReferencedTypes(src) {
		if ref[0] != 'Q' {
			continue
		}
		if ref == "QPlatformPixmap" {
			continue // This class doesn't have a <> version to include
		}
		ret.WriteString(`#include <` + ref + ">\n")
	}

	ret.WriteString("\n")

	for _, c := range src.Classes {

		for i, ctor := range c.Ctors {
			preamble, forwarding := emitParametersCABI2CppForwarding(ctor.Parameters)
			ret.WriteString(fmt.Sprintf(
				"%s* %s_new%s(%s) {\n"+
					"%s"+
					"\treturn new %s(%s);\n"+
					"}\n"+
					"\n",
				c.ClassName, c.ClassName, maybeSuffix(i), emitParametersCabi(ctor, ""),
				preamble,
				c.ClassName, forwarding,
			))
		}

		for _, m := range c.Methods {
			// Need to take an extra 'self' parameter

			shouldReturn := "return "
			afterCall := ""

			if m.ReturnType.ParameterType == "void" && !m.ReturnType.Pointer {
				shouldReturn = ""

			} else if m.ReturnType.ParameterType == "QString" {
				shouldReturn = "QString ret = "
				afterCall = "\t// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
				afterCall += "\tQByteArray b = ret.toUtf8();\n"
				afterCall += "\t*_out = static_cast<char*>(malloc(b.length()));\n"
				afterCall += "\tmemcpy(*_out, b.data(), b.length());\n"
				afterCall += "\t*_out_Strlen = b.length();\n"

			} else if t, ok := m.ReturnType.QListOf(); ok {

				if t.ParameterType == "QString" {
					// Combo
					// "char** _out, int64_t* _out_Lengths, size_t* _out_len")

					shouldReturn = m.ReturnType.ParameterType + " ret = "
					afterCall += "\t// Convert QStringList from C++ memory to manually-managed C memory\n"
					afterCall += "\tchar** __out = static_cast<char**>(malloc(sizeof(char*) * ret.length()));\n"
					afterCall += "\tint64_t* __out_Lengths = static_cast<int64_t*>(malloc(sizeof(int64_t) * ret.length()));\n"
					afterCall += "\tfor (size_t i = 0, e = ret.length(); i < e; ++i) {\n"
					afterCall += "\t\t// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
					afterCall += "\t\tQByteArray b = ret[i].toUtf8();\n"
					afterCall += "\t\t__out[i] = static_cast<char*>(malloc(b.length()));\n"
					afterCall += "\t\tmemcpy(__out[i], b.data(), b.length());\n"
					afterCall += "\t\t__out_Lengths[i] = b.length();\n"
					afterCall += "\t}\n"
					afterCall += "\t*_out = __out;\n"
					afterCall += "\t*_out_Lengths = __out_Lengths;\n"
					afterCall += "\t*_out_len = ret.length();\n"

				} else if !t.QtClassType() || (t.QtClassType() && t.Pointer) { // QList<int>, QList<QFoo*>

					shouldReturn = m.ReturnType.ParameterType + " ret = "
					afterCall += "\t// Convert QList<> from C++ memory to manually-managed C memory\n"
					afterCall += "\t" + t.RenderTypeCpp() + "* __out = static_cast<" + t.RenderTypeCpp() + "*>(malloc(sizeof(" + t.RenderTypeCpp() + ") * ret.length()));\n"
					afterCall += "\tfor (size_t i = 0, e = ret.length(); i < e; ++i) {\n"
					afterCall += "\t\t__out[i] = ret[i];\n"
					afterCall += "\t}\n"
					afterCall += "\t*_out = __out;\n"
					afterCall += "\t*_out_len = ret.length();\n"

				} else { // QList<QFoo>

					shouldReturn = m.ReturnType.ParameterType + " ret = "
					afterCall += "\t// Convert QList<> from C++ memory to manually-managed C memory of copy-constructed pointers\n"
					afterCall += "\t" + t.RenderTypeCpp() + "** __out = static_cast<" + t.RenderTypeCpp() + "**>(malloc(sizeof(" + t.RenderTypeCpp() + "**) * ret.length()));\n"
					afterCall += "\tfor (size_t i = 0, e = ret.length(); i < e; ++i) {\n"
					afterCall += "\t\t__out[i] = new " + t.ParameterType + "(ret[i]);\n"
					afterCall += "\t}\n"
					afterCall += "\t*_out = __out;\n"
					afterCall += "\t*_out_len = ret.length();\n"

				}

			} else if m.ReturnType.QtClassType() && !m.ReturnType.Pointer {
				shouldReturn = m.ReturnType.ParameterType + " ret = "
				afterCall = "\t// Copy-construct value returned type into heap-allocated copy\n"
				afterCall += "\treturn static_cast<" + m.ReturnType.ParameterType + "*>(new " + m.ReturnType.ParameterType + "(ret));\n"

			} else if m.ReturnType.Const {
				shouldReturn += "(" + emitReturnTypeCabi(m.ReturnType) + ") "
			}

			preamble, forwarding := emitParametersCABI2CppForwarding(m.Parameters)

			nativeMethodName := m.MethodName
			if m.OverrideMethodName != "" {
				nativeMethodName = m.OverrideMethodName
			}

			callTarget := "self->"
			if m.IsStatic {
				callTarget = c.ClassName + "::"
			}

			ret.WriteString(fmt.Sprintf(
				"%s %s_%s(%s) {\n"+
					"%s"+
					"\t%s%s%s(%s);\n"+
					"%s"+
					"}\n"+
					"\n",
				emitReturnTypeCabi(m.ReturnType), c.ClassName, m.SafeMethodName(), emitParametersCabi(m, c.ClassName+"*"),
				preamble,
				shouldReturn, callTarget, nativeMethodName, forwarding,
				afterCall,
			))
		}

		// Delete
		ret.WriteString(fmt.Sprintf(
			"void %s_Delete(%s* self) {\n"+
				"\tdelete self;\n"+
				"}\n"+
				"\n",
			c.ClassName, c.ClassName,
		))
	}

	return ret.String(), nil
}
