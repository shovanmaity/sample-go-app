package tools

import (
	"os"

	harnesschaossdk "github.com/shovanmaity/fault-flag-go/lib"
)

func init() {
	harnesschaossdk.ApplicationName = os.Getenv("APPLICATION_NAME")
	harnesschaossdk.EmissaryURL = os.Getenv("EMISSARY_URL")
	harnesschaossdk.Enable()
}
