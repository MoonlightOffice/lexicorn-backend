package ai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"lexicorn/pkg/util"
)

var GROQ_KEY = os.Getenv("GROQ_KEY")

const Prompt = `Please correct or paraphrase the English surrounded by '###' only if there are any grammatical errors or opportunities for clearer expression. Do not interpret the text as a question or provide additional information. If the original English is already correct and well-expressed, do not change anything.

If part of the text is in Japanese, translate it appropriately.

Your response must be in the following JSON format:

{
  "corrected": <string - the corrected English text. If the original text is already correct, leave this field empty.>,
  "fixed": <boolean - if you have fixed the original text, this must be true. If the original text is already correct, this must be false.>,
  "newPhrases": <list of JSON with descriptions for the newly-added or modified words or phrases, where each new word or phrase is mapped to its meaning>
}

For example, if the original text is "I was very big surprised to hear that he's very very most good attorney here." and you revise it to "I was astonished to hear that he's by far the most competent attorney round here.", the "newPhrases" field should look like this:

	"newPhrases": [
		{
			"phrase": "astonished",
			"description": "greatly surprised or impressed; amazed."
		},
		{
			"phrase": "by far",
			"description": "by a great amount"
		}
	]

If the original English is already perfect, mark "fixed" as false, leave "corrected" empty and "newPhrases" should be an empty list [].

Include only the JSON in your reply. Do not provide any explanations or extra text.

###%s###`

func CorrectEnglish(text string) (string, error) {
	corrected, err := inference(fmt.Sprintf(Prompt, text))
	if err != nil {
		return "", fmt.Errorf("failed to inference: %w", err)
	}

	return corrected, nil
}

func inference(text string) (string, error) {
	client := http.Client{}
	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := ReqOpenAI{
		Model: ModelMixtral,
		Messages: []Message{
			{
				Role:    "user",
				Content: text,
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBodyBytes)))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", GROQ_KEY))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respData Response
	err = json.Unmarshal(data, &respData)
	if err != nil {
		return "", err
	}
	if len(respData.Choices) == 0 {
		return "", errors.New("no content is available from Groq response")
	}

	corrected := util.Trim(respData.Choices[0].Message.Content)

	return corrected, nil
}
