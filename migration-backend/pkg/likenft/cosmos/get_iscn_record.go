package cosmos

import (
	"encoding/json"
	"net/url"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

func (a *LikeNFTCosmosClient) GetISCNRecord(
	iscnIdPrefix string,
	iscnVersionAtMint string,
) (*model.ISCN, error) {
	u, err := url.Parse("/iscn/records/id")

	if err != nil {
		return nil, err
	}

	if iscnVersionAtMint == "" {
		iscnVersionAtMint = "1"
	}

	queryString := u.Query()
	queryString.Set("iscn_id", iscnIdPrefix)
	queryString.Set("iscn_version_at_mint", iscnVersionAtMint)
	u.RawQuery = queryString.Encode()

	base, err := url.Parse(a.NodeURL)
	if err != nil {
		return nil, err
	}

	resp, err := a.HTTPClient.Get(base.ResolveReference(u).String())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response model.ISCN
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
