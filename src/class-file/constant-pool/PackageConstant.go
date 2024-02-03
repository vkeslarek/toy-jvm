package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type PackageConstant struct {
	Name ConstantReference[*Utf8Constant]
}

func (constant *PackageConstant) ResolveReferences(cp *ConstantPool) {
	constant.Name.ResolveReferences(cp)
}

func (constant PackageConstant) String() string {
	return fmt.Sprintf("PackageConstant(Name: %s)", constant.Name.String())
}

func (constant PackageConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Package
}

func ParseConstantPackage(binaryReader *reader.BinaryReader, index int) *PackageConstant {
	return &PackageConstant{
		Name: ParseConstantReference[*Utf8Constant](binaryReader, index),
	}
}
