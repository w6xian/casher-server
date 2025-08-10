package version

import (
	"fmt"
	"runtime"
	"strings"
)

// Version is the service current released version.
// Semantic versioning: https://semver.org/
var Version = "0.0.9"

// DevVersion is the service current development version.
var DevVersion = "0.0.9"

func GetCurrentVersion(mode string) string {
	if mode == "dev" || mode == "demo" {
		return DevVersion
	}
	return fmt.Sprintf("%s (built w/%s)", Version, runtime.Version())
}

func GetMinorVersion(version string) string {
	versionList := strings.Split(version, ".")
	if len(versionList) < 3 {
		return ""
	}
	return versionList[0] + "." + versionList[1]
}

func GetSchemaVersion(version string) string {
	minorVersion := GetMinorVersion(version)
	return minorVersion + ".0"
}

func String(mode string) string {
	return GetCurrentVersion(mode)
}

type SortVersion []string

func (s SortVersion) Len() int {
	return len(s)
}

func (s SortVersion) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
