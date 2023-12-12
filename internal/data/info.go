package data

import "gitlab.com/tokend/regources"

type Info interface {
	Info() (*regources.Info, error)
}
