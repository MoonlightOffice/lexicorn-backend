package correct

import (
	"lexicorn/pkg/core/correct"
	"lexicorn/pkg/core/lang"
	"lexicorn/pkg/impl/internal/ai"
	"lexicorn/pkg/util"
)

var CorrectService correct.CorrectService = correctService{}

type correctService struct{}

func (c correctService) Correct(text string, lang lang.Lang) (string, error) {
	if !lang.IsSupported() {
		return "", util.ErrInvalid
	}

	corrected, err := ai.CorrectEnglish(text)
	if err != nil {
		return "", util.ErrBuilder(err)
	}

	return corrected, nil
}
