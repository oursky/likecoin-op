package evm_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeMintNFTsRequestBody(t *testing.T) {
	Convey("MakeMintNFTsRequestBody", t, func() {
		fromTokenId := big.NewInt(0)
		tos := []common.Address{
			common.HexToAddress("0x0"),
		}
		memos := []string{
			"memo1",
		}
		metadataList := []string{
			"metadata1",
		}
		_, err := evm.MakeMintNFTsRequestBody(
			"0x0",
			fromTokenId,
			tos,
			memos,
			metadataList,
		)
		So(err, ShouldBeNil)
	})
}
