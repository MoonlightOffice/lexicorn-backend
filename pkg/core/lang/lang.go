package lang

type Lang string

const (
	English  Lang = "English"
	Japanese Lang = "Japanese"
)

func (l Lang) IsSupported() bool {
	switch l {
	case English:
		return true
	case Japanese:
		return true
	default:
		return false
	}
}
