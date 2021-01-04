package conda

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	Newline                 = "\n"
	defaultRobocorpLocation = "$HOME/.robocorp"
	binSuffix               = "/bin"
)

var (
	FileExtensions = []string{""}
	Shell          = []string{"bash", "--noprofile", "--norc", "-i"}
)

func MicromambaLink() string {
	return "https://downloads.robocorp.com/micromamba/stable/linux64/micromamba"
}

func ExpandPath(entry string) string {
	intermediate := os.ExpandEnv(entry)
	result, err := filepath.Abs(intermediate)
	if err != nil {
		return intermediate
	}
	return result
}

func CondaEnvironment() []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("MAMBA_ROOT_PREFIX=%s", RobocorpHome()))
	return env
}

func BinMicromamba() string {
	return ExpandPath(filepath.Join(BinLocation(), "micromamba"))
}

func CondaPaths(prefix string) []string {
	return []string{ExpandPath(prefix + binSuffix)}
}

func IsPosix() bool {
	return true
}

func IsWindows() bool {
	return false
}

func HasLongPathSupport() bool {
	return true
}

func ValidateLocations() bool {
	return true
}

func EnforceLongpathSupport() error {
	return nil
}
