// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package linuxdmi

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dc3-network/ghw/v99/pkg/context"
	"github.com/dc3-network/ghw/v99/pkg/linuxpath"
	"github.com/dc3-network/ghw/v99/pkg/util"
)

func Item(ctx *context.Context, value string) string {
	paths := linuxpath.New(ctx)
	path := filepath.Join(paths.SysClassDMI, "id", value)

	b, err := os.ReadFile(path)
	if err != nil {
		ctx.Warn("Unable to read %s: %s\n", value, err)
		return util.UNKNOWN
	}

	return strings.TrimSpace(string(b))
}
