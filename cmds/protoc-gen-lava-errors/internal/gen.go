package internal

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
)

// GenerateFile generates a .errors.pb.go file containing service definitions.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + ".errors.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	genFile := jen.NewFile(string(file.GoPackageName))
	genFile.HeaderComment("Code generated by protoc-gen-lava-errors. DO NOT EDIT.")
	genFile.HeaderComment("versions:")
	genFile.HeaderComment(fmt.Sprintf("- protoc-gen-lava-errors %s", version))
	genFile.HeaderComment(fmt.Sprintf("- protoc                 %s", protocVersion(gen)))
	if file.Proto.GetOptions().GetDeprecated() {
		genFile.HeaderComment(fmt.Sprintf("%s is a deprecated file.", file.Desc.Path()))
	} else {
		genFile.HeaderComment(fmt.Sprintf("source: %s", file.Desc.Path()))
	}

	genFile.Comment("This is a compile-time assertion to ensure that this generated file")
	genFile.Comment("is compatible with the grpc package it is being compiled against.")
	genFile.Comment("Requires gRPC-Go v1.32.0 or later.")
	genFile.Id("const _ =").Qual("google.golang.org/grpc", "SupportPackageIsVersion7")
	g.Skip()
	for i := range file.Enums {
		m := file.Enums[i]

		//if m.Desc.Options() == nil {
		//	continue
		//}
		//
		//var enabled, ok = gp.GetExtension(m.Desc.Options(), errorpb.E_Enabled).(bool)
		//if !ok || !enabled {
		//	continue
		//}

		var name = strings.ToLower(string(m.Desc.Name()))
		if !strings.HasSuffix(name, "code") {
			continue
		}

		g.Unskip()
		for j := range m.Values {
			n := m.Values[j]
			var bizCode = jen.Id(fmt.Sprintf(`"%s_%s_%s"`,
				strings.Join(strings.Split(strings.ToLower(string(file.Desc.Package())), "."), "_"),
				strings.ToLower(string(m.Desc.Name())),
				strings.ToLower(string(n.Desc.Name()))))

			// comment
			rr := strings.TrimSpace(strings.ToLower(string(n.Desc.Name())))
			rr = strings.Join(strings.Split(rr, "_"), " ")
			if n.Comments.Leading.String() != "" {
				var cc = n.Comments.Leading.String()
				cc = strings.TrimSpace(cc)
				cc = strings.Trim(cc, "/")
				cc = strings.TrimSpace(cc)
				rr = cc
			}

			genFile.Var().Id(
				fmt.Sprintf("Err%s%s", m.Desc.Name(), case2Camel(string(n.Desc.Name()))),
			).Op("=").Qual("github.com/pubgo/lava/errors", "NewWithBizCode").Params(
				bizCode,
				jen.Lit(rr),
			)
		}
	}

	g.P(genFile.GoString())
	return g
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func case2Camel(name string) string {
	if !strings.Contains(name, "_") {
		return strings.Title(strings.ToLower(name))
	}
	name = strings.ToLower(name)
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
