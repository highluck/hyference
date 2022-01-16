package ml

type MlLibType string
type MlLibTypeMap map[string]MlLibType

var LibTypes = MlLibTypeMap{
	"fasttext": FastText,
}

const FastText = MlLibType("fasttext")
const UnKnown = MlLibType("unknown")
