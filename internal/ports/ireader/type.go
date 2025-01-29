package ireader

import "github.com/sandronister/webcrawler/pkg/parser_html/types"

type Type interface {
	Read(request *types.PageDTO)
}
