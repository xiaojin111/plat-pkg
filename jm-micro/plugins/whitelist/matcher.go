package whitelist

import (
	"fmt"

	"github.com/gobwas/glob"
)

func Match(pattern, service, method string) bool {
	if service == "" || method == "" {
		return false
	}

	g := glob.MustCompile(pattern)
	s := fmt.Sprintf("%s/%s", service, method)
	return g.Match(s)
}
