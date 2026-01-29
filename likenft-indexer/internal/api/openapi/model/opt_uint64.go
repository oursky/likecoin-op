package model

import (
	"strconv"

	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/openapi/api"
)

func MakeOptUint64(n *typeutil.Uint64) api.OptUint64 {
	if n == nil {
		return api.OptUint64{
			Set:   false,
			Value: api.Uint64(""),
		}
	}
	return api.NewOptUint64(api.Uint64(strconv.FormatUint(uint64(*n), 10)))
}
