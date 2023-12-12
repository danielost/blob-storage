package data

import (
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
)

type Infobuilder func(info Info, source keypair.Address) *xdrbuild.Transaction
