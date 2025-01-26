package ireader

import "github.com/sandronister/webcrawler/internal/dto"

type Type interface {
	Read(request *dto.PageDTO)
}
