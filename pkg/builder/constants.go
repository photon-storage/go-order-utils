package builder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/photon-storage/go-order-utils/pkg/eip712"
)

var (
	PROTOCOL_NAME    = crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	PROTOCOL_VERSION = crypto.Keccak256Hash([]byte("2"))
)

var (
	ORDER_STRUCTURE = []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Uint256, // salt
		eip712.Address, // maker
		eip712.Address, // signer
		eip712.Uint256, // tokenId
		eip712.Uint256, // makerAmount
		eip712.Uint256, // takerAmount
		eip712.Uint8,   // side
		eip712.Uint8,   // signatureType
		eip712.Uint256, // timestamp
		eip712.Bytes32, // metadata
		eip712.Bytes32, // builder
	}
)

var (
	OrderContentsType    = []byte("Order(uint256 salt,address maker,address signer,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint8 side,uint8 signatureType,uint256 timestamp,bytes32 metadata,bytes32 builder)")
	ORDER_STRUCTURE_HASH = crypto.Keccak256Hash(OrderContentsType)
)
