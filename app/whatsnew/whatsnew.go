package whatsnew

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config/profile"
)

func Display() {
	var (
		version string
		b       []byte
	)

	f, err := os.OpenFile(profile.ModulePath()+"/version", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		goto changelog
	}

	b, err = io.ReadAll(f)
	f.Close()
	if err != nil {
		goto changelog
	}

	version = string(bytes.TrimSpace(b))

	if version == app.Version {
		return
	}

changelog:
	fmt.Printf("Welcome to murex %d.%d.%d\nChangelog: https://murex.rocks/docs/CHANGELOG.html\n",
		app.Major, app.Minor, app.Revision)

	f, err = os.OpenFile(profile.ModulePath()+"/version", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	f.WriteString(app.Version)
	f.Close()
}
