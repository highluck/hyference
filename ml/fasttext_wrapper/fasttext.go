package fasttext_wrapper

// #cgo LDFLAGS: -L${SRCDIR}/fastText/lib -lfasttext-wrapper -lstdc++ -lm -pthread
// #include <stdlib.h>
// int ft_load_model(char *path);
// int ft_predict(char *query, float *prob, char *buf, int buf_size);
// int ft_get_vector_dimension();
// int ft_get_sentence_vector(char* query_in, float* vector, int vector_size);
import "C"

import (
	"fmt"
	errors2 "github.com/hyference/internal/errors"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type Model struct {
	isInitialized bool
}

type Result struct {
	Label string
	Score float64
}

var re = regexp.MustCompile(`[\{\}\[\]\/?.,;:|\)*~!^\-_+<@\#$%&\\\=\(\'\"\n\r]+`)

func (r Result) GetReplaceLabel() string {
	str := strings.ReplaceAll(r.Label, "__label__", "")
	//str = strings.ReplaceAll(str, "/", "")
	return re.ReplaceAllString(str, "")
}

const UNKNOWN_ID = -1

func (r Result) GetReplaceLabelToId() int64 {
	str := strings.ReplaceAll(r.Label, "__label__", "")
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return UNKNOWN_ID
	}
	return parseInt
}

func New(file string) (*Model, error) {

	status := C.ft_load_model(C.CString(file))

	if status != 0 {
		return nil, fmt.Errorf("Cannot initialize model on `%s`", file)
	}

	return &Model{
		isInitialized: true,
	}, nil
}

func (m *Model) Predict(keyword string) (Result, error) {
	if !m.isInitialized {
		return Result{},
			errors2.New("fast text predict err")
	}

	resultSize := 32
	result := (*C.char)(C.malloc(C.ulong(resultSize)))
	query := C.CString(keyword)
	var cprob C.float

	status := C.ft_predict(
		query,
		&cprob,
		result,
		C.int(resultSize),
	)
	if status != 0 {
		return Result{},
			errors2.New("fast text predict err ")
	}

	label := C.GoString(result)
	prob := float64(cprob)

	C.free(unsafe.Pointer(result))
	C.free(unsafe.Pointer(query))

	return Result{Score: prob, Label: label}, nil
}

func (m *Model) GetSentenceVector(keyword string) ([]float64, error) {

	if !m.isInitialized {
		return nil, errors2.New("The FastText model needs to be initialized first. It's should be done inside the `New()` function")
	}

	vecDim := C.ft_get_vector_dimension()
	var cfloat C.float
	result := (*C.float)(C.malloc(C.ulong(vecDim) * C.ulong(unsafe.Sizeof(cfloat))))

	status := C.ft_get_sentence_vector(
		C.CString(keyword),
		result,
		vecDim,
	)

	if status != 0 {
		return nil, fmt.Errorf("Exception when predicting `%s`", keyword)
	}
	p2 := (*[1 << 30]C.float)(unsafe.Pointer(result))
	ret := make([]float64, int(vecDim))
	for i := 0; i < int(vecDim); i++ {
		ret[i] = float64(p2[i])
	}

	C.free(unsafe.Pointer(result))

	return ret, nil
}
