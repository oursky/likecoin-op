package filecoinoffer

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	gopiece "github.com/storacha/go-piece/pkg/piece"
	"github.com/web3-storage/go-ucanto/core/ipld"
	"github.com/web3-storage/go-ucanto/did"
	"github.com/web3-storage/go-ucanto/ucan"
)

const Ability = "filecoin/offer"

type Caveat struct {
	Content ipld.Link
	Piece   gopiece.PieceLink
}

var _ ucan.CaveatBuilder = (*Caveat)(nil)

func (c *Caveat) Build() (map[string]datamodel.Node, error) {
	data := map[string]datamodel.Node{}

	b := basicnode.Prototype.Link.NewBuilder()
	err := b.AssignLink(c.Content)
	if err != nil {
		return nil, err
	}
	data["content"] = b.Build()

	b = basicnode.Prototype.Link.NewBuilder()
	err = b.AssignLink(c.Piece.Link())
	if err != nil {
		return nil, err
	}
	data["piece"] = b.Build()

	return data, nil
}

func NewCapability(space did.DID, nb *Caveat) ucan.Capability[ucan.MapBuilder] {
	return ucan.NewCapability(Ability, space.String(), ucan.MapBuilder(nb))
}
