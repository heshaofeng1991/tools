package repository

import (
	"fmt"
	"log"
	"runtime/debug"

	ent "combo/ent"

	"github.com/pkg/errors"
)

func RecoverError(r interface{}) (err error) {
	log.Println("panic: ", r, debug.Stack())

	switch v := r.(type) {
	case error:
		err = v
	case string:
		err = errors.New(v)
	default:
		msg := fmt.Sprintf("unknown type err: %+v", v)
		err = errors.New(msg)
	}

	return
}

type TransactionDoFunc func(tx *ent.Tx) (bool, error)
