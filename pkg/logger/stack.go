package logger

import (
	"fmt"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"strings"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func StackFromError(err error, ds ...int) [][]string {
	var depthStart int
	if len(ds) == 0 {
		depthStart = 1
	} else {
		depthStart = ds[0]
	}
	var out [][]string
	var depthEnd int
	depth := 0
	depthLen := 5
	depthEnd = depthStart + depthLen

	common := func(pError stackTracer) {
		st := pError.StackTrace()
		if depth != 0 {
			depthEnd = depth
		}

		for i := depthStart; i < depthEnd; i++ {
			valued := fmt.Sprintf("%+v", st[i])
			valued = strings.Replace(valued, "\t", "", -1)
			stack := strings.Split(valued, "\n")
			out = append(out, stack)
		}
	}

	if err2, ok := err.(stackTracer); ok {
		common(err2)
	}

	if err2, ok := errors.Cause(err).(stackTracer); ok {
		common(err2)
	}

	return out
}
func GetStackField(err error) logger.Fields {
	return logger.Fields{
		"stack": StackFromError(errors.WithStack(err)),
		"err":   err.Error(),
	}
}
