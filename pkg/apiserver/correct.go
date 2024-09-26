package apiserver

import (
	"errors"
	"lexicorn/pkg/core/correct"
	"lexicorn/pkg/core/lang"
	"lexicorn/pkg/util"
	"net/http"
)

type ICorrect interface {
	CorrectLangHandler(w http.ResponseWriter, r *http.Request)
}

type Correct struct {
	correctService correct.CorrectService
}

func NewCorrect(cs correct.CorrectService) ICorrect {
	return &Correct{
		correctService: cs,
	}
}

func (c *Correct) CorrectLangHandler(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Lang lang.Lang `json:"lang"`
		Text string    `json:"text"`
	}{}
	if !bindReqData(r, &reqData) {
		writeResponse(w, http.StatusBadRequest, nil)
		return
	}

	corrected, err := c.correctService.Correct(reqData.Text, reqData.Lang)
	if err != nil {
		if errors.Is(err, util.ErrInvalid) {
			writeResponse(w, http.StatusBadRequest, nil)
			return
		}

		writeResponse(w, http.StatusInternalServerError, H{
			"error": err.Error(),
		})
		return
	}

	writeResponse(w, http.StatusOK, H{
		"corrected": corrected,
	})
}
