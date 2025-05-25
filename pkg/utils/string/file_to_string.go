package utilsString

import (
	"github.com/novando/go-ska/pkg/logger"
	"os"
)

// FileToString open up a file, and extract the strings inside
func FileToString(path string) string {
	str, err := os.ReadFile(path)
	if err != nil {
		logger.Call().Errorf("%v", err.Error())
	}
	return string(str)
}
