// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package likenft_class

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ClassConfig is an auto generated low-level Go binding around an user-defined struct.
type ClassConfig struct {
	MaxSupply uint64
}

// ClassInput is an auto generated low-level Go binding around an user-defined struct.
type ClassInput struct {
	Name     string
	Symbol   string
	Metadata string
	Config   ClassConfig
}

// MsgNewClass is an auto generated low-level Go binding around an user-defined struct.
type MsgNewClass struct {
	Creator common.Address
	Input   ClassInput
}

// LikenftClassMetaData contains all meta data concerning the LikenftClass contract.
var LikenftClassMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structClassConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structClassInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"internalType\":\"structMsgNewClass\",\"name\":\"msgNewClass\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNftNoSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCompatibleWithSpotMints\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialMintExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialUpToTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpotMintTokenIdTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferWithMemo\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"metadata_list\",\"type\":\"string[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferWithMemo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"max_supply\",\"type\":\"uint64\"}],\"internalType\":\"structClassConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"internalType\":\"structClassInput\",\"name\":\"classInput\",\"type\":\"tuple\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161221d38038061221d83398101604081905261002e91610386565b805160208083015180519101516002610047838261055b565b506003610054828261055b565b50505f8055506001600160a01b03811661008757604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100908161017f565b506020810151517f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f2009081906100c5908261055b565b50602080830151015160018201906100dd908261055b565b5060208201516040015160028201906100f6908261055b565b50602082015160600151516003820180546001600160401b0319166001600160401b0390921691909117905561014c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6336101d0565b506101777f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab336101d0565b505050610615565b600980546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f828152600a602090815260408083206001600160a01b038516845290915281205460ff16610274575f838152600a602090815260408083206001600160a01b03861684529091529020805460ff1916600117905561022c3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610277565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156102b3576102b361027d565b60405290565b604051608081016001600160401b03811182821017156102b3576102b361027d565b604051602081016001600160401b03811182821017156102b3576102b361027d565b5f82601f83011261030c575f5ffd5b81516001600160401b038111156103255761032561027d565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103535761035361027d565b60405281815283820160200185101561036a575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f60208284031215610396575f5ffd5b81516001600160401b038111156103ab575f5ffd5b8201604081850312156103bc575f5ffd5b6103c4610291565b81516001600160a01b03811681146103da575f5ffd5b815260208201516001600160401b038111156103f4575f5ffd5b91909101908185036080811215610409575f5ffd5b6104116102b9565b83516001600160401b03811115610426575f5ffd5b610432888287016102fd565b82525060208401516001600160401b0381111561044d575f5ffd5b610459888287016102fd565b60208301525060408401516001600160401b03811115610477575f5ffd5b610483888287016102fd565b6040830152506020605f198301121561049a575f5ffd5b6104a26102db565b606094909401519391506001600160401b03841684146104c0575f5ffd5b928152606083015260208101919091529392505050565b600181811c908216806104eb57607f821691505b60208210810361050957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561055657805f5260205f20601f840160051c810160208510156105345750805b601f840160051c820191505b81811015610553575f8155600101610540565b50505b505050565b81516001600160401b038111156105745761057461027d565b6105888161058284546104d7565b8461050f565b6020601f8211600181146105ba575f83156105a35750848201515b5f19600385901b1c1916600184901b178455610553565b5f84815260208120601f198516915b828110156105e957878501518255602094850194600190920191016105c9565b508482101561060657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b611bfb806106225f395ff3fe6080604052600436106101c5575f3560e01c8063765a15bb116100f2578063c87b56dd11610092578063dcb49c7311610062578063dcb49c73146104ff578063e8a3d48514610512578063e985e9c514610526578063f2fde38b1461056d575f5ffd5b8063c87b56dd1461046f578063d53913931461048e578063d547741f146104c1578063d90794cf146104e0575f5ffd5b806395d89b41116100cd57806395d89b4114610416578063a217fddf1461042a578063a22cb4651461043d578063b88d4fde1461045c575f5ffd5b8063765a15bb146103bb5780638da5cb5b146103da57806391d14854146103f7575f5ffd5b8063248a9ca31161016857806347e633801161013857806347e63380146103365780636352211e1461036957806370a0823114610388578063715018a6146103a7575f5ffd5b8063248a9ca3146102b75780632f2ff15d146102e557806336568abe1461030457806342842e0e14610323575f5ffd5b8063095ea7b3116101a3578063095ea7b31461025557806317d70f7c1461026a57806318160ddd1461028d57806323b872dd146102a4575f5ffd5b806301ffc9a7146101c957806306fdde03146101fd578063081812fc1461021e575b5f5ffd5b3480156101d4575f5ffd5b506101e86101e33660046112ee565b61058c565b60405190151581526020015b60405180910390f35b348015610208575f5ffd5b5061021161059c565b6040516101f4919061133e565b348015610229575f5ffd5b5061023d610238366004611350565b61063d565b6040516001600160a01b0390911681526020016101f4565b61026861026336600461137d565b610676565b005b348015610275575f5ffd5b5061027f600c5481565b6040519081526020016101f4565b348015610298575f5ffd5b506001545f540361027f565b6102686102b23660046113a5565b610686565b3480156102c2575f5ffd5b5061027f6102d1366004611350565b5f908152600a602052604090206001015490565b3480156102f0575f5ffd5b506102686102ff3660046113df565b6107e0565b34801561030f575f5ffd5b5061026861031e3660046113df565b61080a565b6102686103313660046113a5565b610842565b348015610341575f5ffd5b5061027f7f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab81565b348015610374575f5ffd5b5061023d610383366004611350565b61085c565b348015610393575f5ffd5b5061027f6103a2366004611409565b610866565b3480156103b2575f5ffd5b506102686108a9565b3480156103c6575f5ffd5b506102686103d5366004611513565b6108bc565b3480156103e5575f5ffd5b506009546001600160a01b031661023d565b348015610402575f5ffd5b506101e86104113660046113df565b61099f565b348015610421575f5ffd5b506102116109c9565b348015610435575f5ffd5b5061027f5f81565b348015610448575f5ffd5b5061026861045736600461160f565b610a07565b61026861046a366004611648565b610a72565b34801561047a575f5ffd5b50610211610489366004611350565b610aad565b348015610499575f5ffd5b5061027f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b3480156104cc575f5ffd5b506102686104db3660046113df565b610ae5565b3480156104eb575f5ffd5b506102686104fa3660046116be565b610b09565b61026861050d36600461173d565b610c73565b34801561051d575f5ffd5b50610211610cd3565b348015610531575f5ffd5b506101e86105403660046117d0565b6001600160a01b039182165f90815260076020908152604080832093909416825291909152205460ff1690565b348015610578575f5ffd5b50610268610587366004611409565b610d1b565b5f61059682610d5d565b92915050565b5f516020611ba65f395f51905f5280546060919081906105bb906117f8565b80601f01602080910402602001604051908101604052809291908181526020018280546105e7906117f8565b80156106325780601f1061060957610100808354040283529160200191610632565b820191905f5260205f20905b81548152906001019060200180831161061557829003601f168201915b505050505091505090565b5f61064782610d91565b61065b5761065b6333d1c03960e21b610dd3565b505f908152600660205260409020546001600160a01b031690565b61068282826001610ddb565b5050565b5f61069082610e7c565b6001600160a01b0394851694909150811684146106b6576106b662a1148160e81b610dd3565b5f8281526006602052604090208054338082146001600160a01b038816909114176106f9576106e58633610540565b6106f9576106f9632ce44b5f60e11b610dd3565b8015610703575f82555b6001600160a01b038681165f9081526005602052604080822080545f19019055918716808252919020805460010190554260a01b17600160e11b175f85815260046020526040812091909155600160e11b8416900361078f57600184015f81815260046020526040812054900361078d575f54811461078d575f8181526004602052604090208490555b505b6001600160a01b0385168481887fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a4805f036107d7576107d7633a954ecd60e21b610dd3565b50505050505050565b5f828152600a60205260409020600101546107fa81610f0b565b6108048383610f15565b50505050565b6001600160a01b03811633146108335760405163334bd91960e11b815260040160405180910390fd5b61083d8282610fa6565b505050565b61083d83838360405180602001604052805f815250610a72565b5f61059682610e7c565b5f6001600160a01b038216610885576108856323d3ad8160e21b610dd3565b506001600160a01b03165f908152600560205260409020546001600160401b031690565b6108b1611011565b6108ba5f61103e565b565b336108cf6009546001600160a01b031690565b6001600160a01b03161415801561090d575061090b7f73e573f9566d61418a34d5de3ff49360f9c51fec37f7486551670290f6285dab3361099f565b155b1561092b57604051636609677b60e11b815260040160405180910390fd5b80515f516020611ba65f395f51905f529081906109489082611874565b506020820151600182019061095d9082611874565b50604082015160028201906109729082611874565b506040517fa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962905f90a15050565b5f918252600a602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f20180546060915f516020611ba65f395f51905f52916105bb906117f8565b335f8181526007602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b610a7d848484610686565b6001600160a01b0383163b1561080457610a998484848461108f565b610804576108046368d2bf6b60e11b610dd3565b5f818152600b60209081526040918290209151606092610acf9290910161192e565b6040516020818303038152906040529050919050565b5f828152600a6020526040902060010154610aff81610f0b565b6108048383610fa6565b33610b1c6009546001600160a01b031690565b6001600160a01b031614158015610b5a5750610b587f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a63361099f565b155b15610b7857604051636609677b60e11b815260040160405180910390fd5b5f547f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f203545f516020611ba65f395f51905f5291906001600160401b0316838115801590610be35750816001600160401b031681610bd76001545f540390565b610be191906119e2565b115b15610c0157604051636a29267160e01b815260040160405180910390fd5b610c0b878261116d565b5f5b81811015610c69575f610c2082866119e2565b9050878783818110610c3457610c346119f5565b9050602002810190610c469190611a09565b5f838152600b6020526040902091610c5f919083611a52565b5050600101610c0d565b5050505050505050565b610c7e858585610686565b82846001600160a01b0316866001600160a01b03167fbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b8585604051610cc4929190611b0b565b60405180910390a45050505050565b604051606090610d07907f99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f2029060200161192e565b604051602081830303815290604052905090565b610d23611011565b6001600160a01b038116610d5157604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610d5a8161103e565b50565b5f6001600160e01b03198216637965db0b60e01b148061059657506301ffc9a760e01b6001600160e01b0319831614610596565b5f5f54821015610dce575f5b505f8281526004602052604081205490819003610dc457610dbd83611b39565b9250610d9d565b600160e01b161590505b919050565b805f5260045ffd5b5f610de58361085c565b9050818015610dfd5750336001600160a01b03821614155b15610e2057610e0c8133610540565b610e2057610e206367d9dca160e11b610dd3565b5f8381526006602052604080822080546001600160a01b0319166001600160a01b0388811691821790925591518693918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a450505050565b5f81815260046020526040902054805f03610ee9575f548210610ea957610ea9636f96cda160e11b610dd3565b5b505f19015f818152600460205260409020548015610eaa57600160e01b81165f03610ed457919050565b610ee4636f96cda160e11b610dd3565b610eaa565b600160e01b81165f03610efb57919050565b610dce636f96cda160e11b610dd3565b610d5a8133611186565b5f610f20838361099f565b610f9f575f838152600a602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610f573390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610596565b505f610596565b5f610fb1838361099f565b15610f9f575f838152600a602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610596565b6009546001600160a01b031633146108ba5760405163118cdaa760e01b8152336004820152602401610d48565b600980546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b604051630a85bd0160e11b81525f906001600160a01b0385169063150b7a02906110c3903390899088908890600401611b4e565b6020604051808303815f875af19250505080156110fd575060408051601f3d908101601f191682019092526110fa91810190611b8a565b60015b611150573d80801561112a576040519150601f19603f3d011682016040523d82523d5f602084013e61112f565b606091505b5080515f03611148576111486368d2bf6b60e11b610dd3565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050949350505050565b610682828260405180602001604052805f8152506111bf565b611190828261099f565b6106825760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610d48565b6111c9838361121f565b6001600160a01b0383163b1561083d575f548281035b6111f15f86838060010194508661108f565b611205576112056368d2bf6b60e11b610dd3565b8181106111df57815f5414611218575f5ffd5b5050505050565b5f80549082900361123a5761123a63b562e8dd60e01b610dd3565b5f8181526004602090815260408083206001600160a01b0387164260a01b6001881460e11b1781179091558084526005909252822080546801000000000000000186020190559081900361129757611297622e076360e81b610dd3565b818301825b80835f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f5fa481816001019150810361129c57505f5550505050565b6001600160e01b031981168114610d5a575f5ffd5b5f602082840312156112fe575f5ffd5b8135611309816112d9565b9392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6113096020830184611310565b5f60208284031215611360575f5ffd5b5035919050565b80356001600160a01b0381168114610dce575f5ffd5b5f5f6040838503121561138e575f5ffd5b61139783611367565b946020939093013593505050565b5f5f5f606084860312156113b7575f5ffd5b6113c084611367565b92506113ce60208501611367565b929592945050506040919091013590565b5f5f604083850312156113f0575f5ffd5b8235915061140060208401611367565b90509250929050565b5f60208284031215611419575f5ffd5b61130982611367565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b038111828210171561145857611458611422565b60405290565b604051602081016001600160401b038111828210171561145857611458611422565b5f5f6001600160401b0384111561149957611499611422565b50604051601f19601f85018116603f011681018181106001600160401b03821117156114c7576114c7611422565b6040528381529050808284018510156114de575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f830112611504575f5ffd5b61130983833560208501611480565b5f60208284031215611523575f5ffd5b81356001600160401b03811115611538575f5ffd5b8201808403608081121561154a575f5ffd5b611552611436565b82356001600160401b03811115611567575f5ffd5b611573878286016114f5565b82525060208301356001600160401b0381111561158e575f5ffd5b61159a878286016114f5565b60208301525060408301356001600160401b038111156115b8575f5ffd5b6115c4878286016114f5565b6040830152506020605f19830112156115db575f5ffd5b6115e361145e565b9150606083013592506001600160401b0383168314611600575f5ffd5b91815260608201529392505050565b5f5f60408385031215611620575f5ffd5b61162983611367565b91506020830135801515811461163d575f5ffd5b809150509250929050565b5f5f5f5f6080858703121561165b575f5ffd5b61166485611367565b935061167260208601611367565b92506040850135915060608501356001600160401b03811115611693575f5ffd5b8501601f810187136116a3575f5ffd5b6116b287823560208401611480565b91505092959194509250565b5f5f5f604084860312156116d0575f5ffd5b6116d984611367565b925060208401356001600160401b038111156116f3575f5ffd5b8401601f81018613611703575f5ffd5b80356001600160401b03811115611718575f5ffd5b8660208260051b840101111561172c575f5ffd5b939660209190910195509293505050565b5f5f5f5f5f60808688031215611751575f5ffd5b61175a86611367565b945061176860208701611367565b93506040860135925060608601356001600160401b03811115611789575f5ffd5b8601601f81018813611799575f5ffd5b80356001600160401b038111156117ae575f5ffd5b8860208284010111156117bf575f5ffd5b959894975092955050506020019190565b5f5f604083850312156117e1575f5ffd5b6117ea83611367565b915061140060208401611367565b600181811c9082168061180c57607f821691505b60208210810361182a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561083d57805f5260205f20601f840160051c810160208510156118555750805b601f840160051c820191505b81811015611218575f8155600101611861565b81516001600160401b0381111561188d5761188d611422565b6118a18161189b84546117f8565b84611830565b6020601f8211600181146118d3575f83156118bc5750848201515b5f19600385901b1c1916600184901b178455611218565b5f84815260208120601f198516915b8281101561190257878501518255602094850194600190920191016118e2565b508482101561191f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b7f646174613a6170706c69636174696f6e2f6a736f6e3b757466382c000000000081525f5f835461195e816117f8565b6001821680156119755760018114611990576119c3565b60ff198316601b870152601b821515830287010193506119c3565b865f5260205f205f5b838110156119b8578154888201601b0152600190910190602001611999565b5050601b8287010193505b509195945050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610596576105966119ce565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112611a1e575f5ffd5b8301803591506001600160401b03821115611a37575f5ffd5b602001915036819003821315611a4b575f5ffd5b9250929050565b6001600160401b03831115611a6957611a69611422565b611a7d83611a7783546117f8565b83611830565b5f601f841160018114611aae575f8515611a975750838201355b5f19600387901b1c1916600186901b178355611218565b5f83815260208120601f198716915b82811015611add5786850135825560209485019460019092019101611abd565b5086821015611af9575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f81611b4757611b476119ce565b505f190190565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611b8090830184611310565b9695505050505050565b5f60208284031215611b9a575f5ffd5b8151611309816112d956fe99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f200a264697066735822122005b8bd1c7e0e07e5ba22977a5e0bec433868e824de8e7025ca0be2847f4422fb64736f6c634300081c0033",
}

// LikenftClassABI is the input ABI used to generate the binding from.
// Deprecated: Use LikenftClassMetaData.ABI instead.
var LikenftClassABI = LikenftClassMetaData.ABI

// LikenftClassBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LikenftClassMetaData.Bin instead.
var LikenftClassBin = LikenftClassMetaData.Bin

// DeployLikenftClass deploys a new Ethereum contract, binding an instance of LikenftClass to it.
func DeployLikenftClass(auth *bind.TransactOpts, backend bind.ContractBackend, msgNewClass MsgNewClass) (common.Address, *types.Transaction, *LikenftClass, error) {
	parsed, err := LikenftClassMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LikenftClassBin), backend, msgNewClass)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LikenftClass{LikenftClassCaller: LikenftClassCaller{contract: contract}, LikenftClassTransactor: LikenftClassTransactor{contract: contract}, LikenftClassFilterer: LikenftClassFilterer{contract: contract}}, nil
}

// LikenftClass is an auto generated Go binding around an Ethereum contract.
type LikenftClass struct {
	LikenftClassCaller     // Read-only binding to the contract
	LikenftClassTransactor // Write-only binding to the contract
	LikenftClassFilterer   // Log filterer for contract events
}

// LikenftClassCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikenftClassCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikenftClassTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikenftClassFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikenftClassSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikenftClassSession struct {
	Contract     *LikenftClass     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LikenftClassCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikenftClassCallerSession struct {
	Contract *LikenftClassCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LikenftClassTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikenftClassTransactorSession struct {
	Contract     *LikenftClassTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LikenftClassRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikenftClassRaw struct {
	Contract *LikenftClass // Generic contract binding to access the raw methods on
}

// LikenftClassCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikenftClassCallerRaw struct {
	Contract *LikenftClassCaller // Generic read-only contract binding to access the raw methods on
}

// LikenftClassTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikenftClassTransactorRaw struct {
	Contract *LikenftClassTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikenftClass creates a new instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClass(address common.Address, backend bind.ContractBackend) (*LikenftClass, error) {
	contract, err := bindLikenftClass(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikenftClass{LikenftClassCaller: LikenftClassCaller{contract: contract}, LikenftClassTransactor: LikenftClassTransactor{contract: contract}, LikenftClassFilterer: LikenftClassFilterer{contract: contract}}, nil
}

// NewLikenftClassCaller creates a new read-only instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassCaller(address common.Address, caller bind.ContractCaller) (*LikenftClassCaller, error) {
	contract, err := bindLikenftClass(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikenftClassCaller{contract: contract}, nil
}

// NewLikenftClassTransactor creates a new write-only instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassTransactor(address common.Address, transactor bind.ContractTransactor) (*LikenftClassTransactor, error) {
	contract, err := bindLikenftClass(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransactor{contract: contract}, nil
}

// NewLikenftClassFilterer creates a new log filterer instance of LikenftClass, bound to a specific deployed contract.
func NewLikenftClassFilterer(address common.Address, filterer bind.ContractFilterer) (*LikenftClassFilterer, error) {
	contract, err := bindLikenftClass(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikenftClassFilterer{contract: contract}, nil
}

// bindLikenftClass binds a generic wrapper to an already deployed contract.
func bindLikenftClass(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikenftClassMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikenftClass *LikenftClassRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikenftClass.Contract.LikenftClassCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikenftClass *LikenftClassRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.Contract.LikenftClassTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikenftClass *LikenftClassRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikenftClass.Contract.LikenftClassTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikenftClass *LikenftClassCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikenftClass.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikenftClass *LikenftClassTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikenftClass *LikenftClassTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikenftClass.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LikenftClass.Contract.DEFAULTADMINROLE(&_LikenftClass.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LikenftClass.Contract.DEFAULTADMINROLE(&_LikenftClass.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) MINTERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.MINTERROLE(&_LikenftClass.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) MINTERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.MINTERROLE(&_LikenftClass.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) UPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassSession) UPDATERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.UPDATERROLE(&_LikenftClass.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) UPDATERROLE() ([32]byte, error) {
	return _LikenftClass.Contract.UPDATERROLE(&_LikenftClass.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikenftClass.Contract.BalanceOf(&_LikenftClass.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LikenftClass.Contract.BalanceOf(&_LikenftClass.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassSession) ContractURI() (string, error) {
	return _LikenftClass.Contract.ContractURI(&_LikenftClass.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) ContractURI() (string, error) {
	return _LikenftClass.Contract.ContractURI(&_LikenftClass.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.GetApproved(&_LikenftClass.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.GetApproved(&_LikenftClass.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LikenftClass.Contract.GetRoleAdmin(&_LikenftClass.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LikenftClass *LikenftClassCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LikenftClass.Contract.GetRoleAdmin(&_LikenftClass.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LikenftClass.Contract.HasRole(&_LikenftClass.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LikenftClass.Contract.HasRole(&_LikenftClass.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikenftClass.Contract.IsApprovedForAll(&_LikenftClass.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LikenftClass.Contract.IsApprovedForAll(&_LikenftClass.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassSession) Name() (string, error) {
	return _LikenftClass.Contract.Name(&_LikenftClass.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) Name() (string, error) {
	return _LikenftClass.Contract.Name(&_LikenftClass.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassSession) Owner() (common.Address, error) {
	return _LikenftClass.Contract.Owner(&_LikenftClass.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LikenftClass *LikenftClassCallerSession) Owner() (common.Address, error) {
	return _LikenftClass.Contract.Owner(&_LikenftClass.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.OwnerOf(&_LikenftClass.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LikenftClass *LikenftClassCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LikenftClass.Contract.OwnerOf(&_LikenftClass.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikenftClass.Contract.SupportsInterface(&_LikenftClass.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LikenftClass *LikenftClassCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LikenftClass.Contract.SupportsInterface(&_LikenftClass.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassSession) Symbol() (string, error) {
	return _LikenftClass.Contract.Symbol(&_LikenftClass.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikenftClass *LikenftClassCallerSession) Symbol() (string, error) {
	return _LikenftClass.Contract.Symbol(&_LikenftClass.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_LikenftClass *LikenftClassCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_LikenftClass *LikenftClassSession) TokenId() (*big.Int, error) {
	return _LikenftClass.Contract.TokenId(&_LikenftClass.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_LikenftClass *LikenftClassCallerSession) TokenId() (*big.Int, error) {
	return _LikenftClass.Contract.TokenId(&_LikenftClass.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _LikenftClass.Contract.TokenURI(&_LikenftClass.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_LikenftClass *LikenftClassCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _LikenftClass.Contract.TokenURI(&_LikenftClass.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_LikenftClass *LikenftClassCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikenftClass.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_LikenftClass *LikenftClassSession) TotalSupply() (*big.Int, error) {
	return _LikenftClass.Contract.TotalSupply(&_LikenftClass.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_LikenftClass *LikenftClassCallerSession) TotalSupply() (*big.Int, error) {
	return _LikenftClass.Contract.TotalSupply(&_LikenftClass.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.Approve(&_LikenftClass.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.Approve(&_LikenftClass.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.GrantRole(&_LikenftClass.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.GrantRole(&_LikenftClass.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadata_list) returns()
func (_LikenftClass *LikenftClassTransactor) Mint(opts *bind.TransactOpts, to common.Address, metadata_list []string) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "mint", to, metadata_list)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadata_list) returns()
func (_LikenftClass *LikenftClassSession) Mint(to common.Address, metadata_list []string) (*types.Transaction, error) {
	return _LikenftClass.Contract.Mint(&_LikenftClass.TransactOpts, to, metadata_list)
}

// Mint is a paid mutator transaction binding the contract method 0xd90794cf.
//
// Solidity: function mint(address to, string[] metadata_list) returns()
func (_LikenftClass *LikenftClassTransactorSession) Mint(to common.Address, metadata_list []string) (*types.Transaction, error) {
	return _LikenftClass.Contract.Mint(&_LikenftClass.TransactOpts, to, metadata_list)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceOwnership(&_LikenftClass.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LikenftClass *LikenftClassTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceOwnership(&_LikenftClass.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceRole(&_LikenftClass.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LikenftClass *LikenftClassTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RenounceRole(&_LikenftClass.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RevokeRole(&_LikenftClass.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LikenftClass *LikenftClassTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.RevokeRole(&_LikenftClass.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_LikenftClass *LikenftClassTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_LikenftClass *LikenftClassSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom0(&_LikenftClass.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LikenftClass.Contract.SafeTransferFrom0(&_LikenftClass.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.Contract.SetApprovalForAll(&_LikenftClass.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LikenftClass *LikenftClassTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LikenftClass.Contract.SetApprovalForAll(&_LikenftClass.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferFrom(&_LikenftClass.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferOwnership(&_LikenftClass.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferOwnership(&_LikenftClass.TransactOpts, newOwner)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassTransactor) TransferWithMemo(opts *bind.TransactOpts, from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "transferWithMemo", from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferWithMemo(&_LikenftClass.TransactOpts, from, to, _tokenId, memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address from, address to, uint256 _tokenId, string memo) payable returns()
func (_LikenftClass *LikenftClassTransactorSession) TransferWithMemo(from common.Address, to common.Address, _tokenId *big.Int, memo string) (*types.Transaction, error) {
	return _LikenftClass.Contract.TransferWithMemo(&_LikenftClass.TransactOpts, from, to, _tokenId, memo)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassTransactor) Update(opts *bind.TransactOpts, classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.contract.Transact(opts, "update", classInput)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassSession) Update(classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.Contract.Update(&_LikenftClass.TransactOpts, classInput)
}

// Update is a paid mutator transaction binding the contract method 0x765a15bb.
//
// Solidity: function update((string,string,string,(uint64)) classInput) returns()
func (_LikenftClass *LikenftClassTransactorSession) Update(classInput ClassInput) (*types.Transaction, error) {
	return _LikenftClass.Contract.Update(&_LikenftClass.TransactOpts, classInput)
}

// LikenftClassApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LikenftClass contract.
type LikenftClassApprovalIterator struct {
	Event *LikenftClassApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassApproval represents a Approval event raised by the LikenftClass contract.
type LikenftClassApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LikenftClassApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassApprovalIterator{contract: _LikenftClass.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LikenftClassApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassApproval)
				if err := _LikenftClass.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) ParseApproval(log types.Log) (*LikenftClassApproval, error) {
	event := new(LikenftClassApproval)
	if err := _LikenftClass.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LikenftClass contract.
type LikenftClassApprovalForAllIterator struct {
	Event *LikenftClassApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassApprovalForAll represents a ApprovalForAll event raised by the LikenftClass contract.
type LikenftClassApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikenftClass *LikenftClassFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LikenftClassApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassApprovalForAllIterator{contract: _LikenftClass.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikenftClass *LikenftClassFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LikenftClassApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassApprovalForAll)
				if err := _LikenftClass.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LikenftClass *LikenftClassFilterer) ParseApprovalForAll(log types.Log) (*LikenftClassApprovalForAll, error) {
	event := new(LikenftClassApprovalForAll)
	if err := _LikenftClass.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the LikenftClass contract.
type LikenftClassConsecutiveTransferIterator struct {
	Event *LikenftClassConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassConsecutiveTransfer represents a ConsecutiveTransfer event raised by the LikenftClass contract.
type LikenftClassConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_LikenftClass *LikenftClassFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*LikenftClassConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassConsecutiveTransferIterator{contract: _LikenftClass.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_LikenftClass *LikenftClassFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *LikenftClassConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassConsecutiveTransfer)
				if err := _LikenftClass.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_LikenftClass *LikenftClassFilterer) ParseConsecutiveTransfer(log types.Log) (*LikenftClassConsecutiveTransfer, error) {
	event := new(LikenftClassConsecutiveTransfer)
	if err := _LikenftClass.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the LikenftClass contract.
type LikenftClassContractURIUpdatedIterator struct {
	Event *LikenftClassContractURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassContractURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassContractURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassContractURIUpdated represents a ContractURIUpdated event raised by the LikenftClass contract.
type LikenftClassContractURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_LikenftClass *LikenftClassFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*LikenftClassContractURIUpdatedIterator, error) {

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &LikenftClassContractURIUpdatedIterator{contract: _LikenftClass.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_LikenftClass *LikenftClassFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *LikenftClassContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassContractURIUpdated)
				if err := _LikenftClass.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractURIUpdated is a log parse operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_LikenftClass *LikenftClassFilterer) ParseContractURIUpdated(log types.Log) (*LikenftClassContractURIUpdated, error) {
	event := new(LikenftClassContractURIUpdated)
	if err := _LikenftClass.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LikenftClass contract.
type LikenftClassOwnershipTransferredIterator struct {
	Event *LikenftClassOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassOwnershipTransferred represents a OwnershipTransferred event raised by the LikenftClass contract.
type LikenftClassOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikenftClass *LikenftClassFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LikenftClassOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassOwnershipTransferredIterator{contract: _LikenftClass.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikenftClass *LikenftClassFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LikenftClassOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassOwnershipTransferred)
				if err := _LikenftClass.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LikenftClass *LikenftClassFilterer) ParseOwnershipTransferred(log types.Log) (*LikenftClassOwnershipTransferred, error) {
	event := new(LikenftClassOwnershipTransferred)
	if err := _LikenftClass.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LikenftClass contract.
type LikenftClassRoleAdminChangedIterator struct {
	Event *LikenftClassRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleAdminChanged represents a RoleAdminChanged event raised by the LikenftClass contract.
type LikenftClassRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LikenftClass *LikenftClassFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LikenftClassRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleAdminChangedIterator{contract: _LikenftClass.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LikenftClass *LikenftClassFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleAdminChanged)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LikenftClass *LikenftClassFilterer) ParseRoleAdminChanged(log types.Log) (*LikenftClassRoleAdminChanged, error) {
	event := new(LikenftClassRoleAdminChanged)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LikenftClass contract.
type LikenftClassRoleGrantedIterator struct {
	Event *LikenftClassRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleGranted represents a RoleGranted event raised by the LikenftClass contract.
type LikenftClassRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LikenftClassRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleGrantedIterator{contract: _LikenftClass.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleGranted)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) ParseRoleGranted(log types.Log) (*LikenftClassRoleGranted, error) {
	event := new(LikenftClassRoleGranted)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LikenftClass contract.
type LikenftClassRoleRevokedIterator struct {
	Event *LikenftClassRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassRoleRevoked represents a RoleRevoked event raised by the LikenftClass contract.
type LikenftClassRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LikenftClassRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassRoleRevokedIterator{contract: _LikenftClass.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LikenftClassRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassRoleRevoked)
				if err := _LikenftClass.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LikenftClass *LikenftClassFilterer) ParseRoleRevoked(log types.Log) (*LikenftClassRoleRevoked, error) {
	event := new(LikenftClassRoleRevoked)
	if err := _LikenftClass.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LikenftClass contract.
type LikenftClassTransferIterator struct {
	Event *LikenftClassTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassTransfer represents a Transfer event raised by the LikenftClass contract.
type LikenftClassTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LikenftClassTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransferIterator{contract: _LikenftClass.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LikenftClassTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassTransfer)
				if err := _LikenftClass.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LikenftClass *LikenftClassFilterer) ParseTransfer(log types.Log) (*LikenftClassTransfer, error) {
	event := new(LikenftClassTransfer)
	if err := _LikenftClass.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikenftClassTransferWithMemoIterator is returned from FilterTransferWithMemo and is used to iterate over the raw logs and unpacked data for TransferWithMemo events raised by the LikenftClass contract.
type LikenftClassTransferWithMemoIterator struct {
	Event *LikenftClassTransferWithMemo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LikenftClassTransferWithMemoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikenftClassTransferWithMemo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LikenftClassTransferWithMemo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LikenftClassTransferWithMemoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikenftClassTransferWithMemoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikenftClassTransferWithMemo represents a TransferWithMemo event raised by the LikenftClass contract.
type LikenftClassTransferWithMemo struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferWithMemo is a free log retrieval operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_LikenftClass *LikenftClassFilterer) FilterTransferWithMemo(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LikenftClassTransferWithMemoIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.FilterLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LikenftClassTransferWithMemoIterator{contract: _LikenftClass.contract, event: "TransferWithMemo", logs: logs, sub: sub}, nil
}

// WatchTransferWithMemo is a free log subscription operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_LikenftClass *LikenftClassFilterer) WatchTransferWithMemo(opts *bind.WatchOpts, sink chan<- *LikenftClassTransferWithMemo, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LikenftClass.contract.WatchLogs(opts, "TransferWithMemo", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikenftClassTransferWithMemo)
				if err := _LikenftClass.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferWithMemo is a log parse operation binding the contract event 0xbd5c95affecf80a51b513bb4eddd42724421b80ef31b07cee1b5b25d8ce5a05b.
//
// Solidity: event TransferWithMemo(address indexed from, address indexed to, uint256 indexed tokenId, string memo)
func (_LikenftClass *LikenftClassFilterer) ParseTransferWithMemo(log types.Log) (*LikenftClassTransferWithMemo, error) {
	event := new(LikenftClassTransferWithMemo)
	if err := _LikenftClass.contract.UnpackLog(event, "TransferWithMemo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
