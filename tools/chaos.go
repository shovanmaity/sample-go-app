package tools

import (
	"os"

	harnesschaossdk "github.com/harness/fault-flag-go"
)

func init() {
	harnesschaossdk.ApplicationName = os.Getenv("APPLICATION_NAME")
	harnesschaossdk.EmissaryURL = os.Getenv("EMISSARY_URL")
	harnesschaossdk.Enable()
}
