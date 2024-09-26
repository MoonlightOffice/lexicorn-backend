All endpoints accept POST method.

__/correct-lang__
```json
/* Request */
{
	"lang": "English",
	"text": "He have pen."
}

/* Response */

// 200
{
  "corrected": {
    "corrected": "He has a pen.",
    "fixed": true,
    "newPhrases": [
      {
        "phrase": "has",
        "description": "used to indicate possession; to be the owner of."
      },
      {
        "phrase": "a",
        "description": "used in English to indicate one or some, no matter how much or how little."
      }
    ]
  }
}

// 400
```

Example:

`curl -X POST localhost:8000/correct-lang -d '{"lang": "English", "text": "He have pen"}'`