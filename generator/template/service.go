package template

import (
	"github.com/vetcher/go-astra/types"
)

func constructorName(p *types.Interface) string {
	return "New" + p.Name
}
