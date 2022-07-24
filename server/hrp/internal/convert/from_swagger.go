package convert

import (
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/test-instructor/cheetah/server/hrp"
	"github.com/test-instructor/cheetah/server/hrp/internal/builtin"
)

func LoadSwaggerCase(path string) (*hrp.TCase, error) {
	// load swagger file
	caseSwagger := new(spec.Swagger)
	err := builtin.LoadFile(path, caseSwagger)
	if err != nil {
		return nil, errors.Wrap(err, "load swagger file failed")
	}
	if caseSwagger.Definitions == nil {
		return nil, errors.New("invalid swagger case file, missing definitions")
	}

	// TODO: convert swagger to TCase
	return nil, nil
}
