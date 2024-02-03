package version

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type JavaVersion struct {
	Major uint16
	Minor uint16
}

var (
	JavaVersionUndefined = JavaVersion{0, 0}
	JavaVersion1_0_2     = JavaVersion{45, 3}
	JavaVersion1_1       = JavaVersion{45, 3}
	JavaVersion5_0       = JavaVersion{49, 0}
	JavaVersion6         = JavaVersion{50, 0}
	JavaVersion7         = JavaVersion{51, 0}
	JavaVersion8         = JavaVersion{52, 0}
	JavaVersion9         = JavaVersion{53, 0}
	JavaVersion10        = JavaVersion{54, 0}
	JavaVersion11        = JavaVersion{55, 0}
	JavaVersion12        = JavaVersion{56, 0}
	JavaVersion13        = JavaVersion{57, 0}
	JavaVersion14        = JavaVersion{58, 0}
	JavaVersion15        = JavaVersion{59, 0}
	JavaVersion16        = JavaVersion{60, 0}
	JavaVersion17        = JavaVersion{61, 0}
	JavaVersion18        = JavaVersion{62, 0}
	JavaVersion19        = JavaVersion{63, 0}
	JavaVersion20        = JavaVersion{64, 0}
	JavaVersion21        = JavaVersion{65, 0}
)

func (v JavaVersion) String() string {
	return fmt.Sprintf("JavaVersion(Major: %d, Minor: %d)", v.Major, v.Minor)
}

func ParseJavaVersion(binaryReader *reader.BinaryReader) JavaVersion {
	return JavaVersion{
		Minor: binaryReader.ReadUint16("$.Minor"),
		Major: binaryReader.ReadUint16("$.Major"),
	}
}
