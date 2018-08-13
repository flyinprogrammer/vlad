package grifts

import (
	"github.com/flyinprogrammer/vlad/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
