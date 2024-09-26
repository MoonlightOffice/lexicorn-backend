package correct

import "lexicorn/pkg/core/lang"

type CorrectService interface {
	// Correct or paraphrase given text based on given language.
	//
	// Err: util.ErrInvalid
	Correct(text string, lang lang.Lang) (string, error)
}
