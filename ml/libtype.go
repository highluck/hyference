package ml

type LibType string
type LibTypeMap map[string]LibType

var LibTypes = LibTypeMap{
	"fasttext": FastText,
}

const FastText = LibType("fasttext")
const UnKnown = LibType("unknown")
