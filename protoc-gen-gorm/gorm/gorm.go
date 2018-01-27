package gorm

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

const generatedCodeVersion = 4

func init() {
	generator.RegisterPlugin(new(grpc))
}

type gorm struct {
	gen *generator.Generator
}

func (g *gorm) Name() string {
	return "gorm"
}

func (g *gorm) Init(gen *generator.Generator) {
	g.gen = gen
}

func (g *gorm) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

func (g *gorm) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

func (g *gorm) P(args ...interface{}) { g.gen.P(args...) }

func (g *gorm) Generate(file *generator.FileDescriptor) {
	g.P("testing123")
}

func (g *gorm) GenerateImports(file *generator.FileDescriptor) {
}
