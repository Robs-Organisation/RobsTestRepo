package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// CoverageCov default implementation.
func CoverageCov(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("coverage.html"))
}
