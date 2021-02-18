package types

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// PeggyDenomPrefix indicates the prefix for all assests minted by this module
	PeggyDenomPrefix = ModuleName

	// PeggyDenomSeparator is the separator for peggy denoms
	PeggyDenomSeparator = ""

	// ETHContractAddressLen is the length of contract address strings
	ETHContractAddressLen = 42

	// PeggyDenomLen is the length of the denoms generated by the peggy module
	PeggyDenomLen = len(PeggyDenomPrefix) + len(PeggyDenomSeparator) + ETHContractAddressLen
)

// EthAddrLessThan migrates the Ethereum address less than function
func EthAddrLessThan(e, o string) bool {
	return bytes.Compare([]byte(e)[:], []byte(o)[:]) == -1
}

// ValidateEthAddress validates the ethereum address strings
func ValidateEthAddress(a string) error {
	if a == "" {
		return fmt.Errorf("empty")
	}
	if !regexp.MustCompile("^0x[0-9a-fA-F]{40}$").MatchString(a) {
		return fmt.Errorf("address(%s) doesn't pass regex", a)
	}
	if len(a) != ETHContractAddressLen {
		return fmt.Errorf("address(%s) of the wrong length exp(%d) actual(%d)", a, len(a), ETHContractAddressLen)
	}
	return nil
}

/////////////////////////
//     ERC20Token      //
/////////////////////////

// NewERC20Token returns a new instance of an ERC20
func NewERC20Token(amount uint64, contract string) *ERC20Token {
	return &ERC20Token{Amount: sdk.NewIntFromUint64(amount), Contract: contract}
}

// PeggyCoin returns the peggy representation of the ERC20
func (e *ERC20Token) PeggyCoin() sdk.Coin {
	return sdk.NewCoin(PeggyDenom(e.Contract), e.Amount)
}

func PeggyDenom(tokenContract string) string {
	return fmt.Sprintf("%s%s%s", PeggyDenomPrefix, PeggyDenomSeparator, tokenContract)
}

// ValidateBasic permforms stateless validation
func (e *ERC20Token) ValidateBasic() error {
	if err := ValidateEthAddress(e.Contract); err != nil {
		return sdkerrors.Wrap(err, "ethereum address")
	}
	// TODO: Validate all the things
	return nil
}

// Add adds one ERC20 to another
// TODO: make this return errors instead
func (e *ERC20Token) Add(o *ERC20Token) *ERC20Token {
	if string(e.Contract) != string(o.Contract) {
		panic("invalid contract address")
	}
	sum := e.Amount.Add(o.Amount)
	if !sum.IsUint64() {
		panic("invalid amount")
	}
	return NewERC20Token(sum.Uint64(), e.Contract)
}

func PeggyDenomToERC20(denom string) (string, error) {
	fullPrefix := PeggyDenomPrefix + PeggyDenomSeparator
	if !strings.HasPrefix(denom, fullPrefix) {
		return "", fmt.Errorf("denom prefix(%s) not equal to expected(%s)", denom, fullPrefix)
	}
	contract := strings.TrimPrefix(denom, fullPrefix)
	err := ValidateEthAddress(contract)
	switch {
	case err != nil:
		return "", fmt.Errorf("error(%s) validating ethereum contract address", err)
	case len(denom) != PeggyDenomLen:
		return "", fmt.Errorf("len(denom)(%d) not equal to PeggyDenomLen(%d)", len(denom), PeggyDenomLen)
	default:
		return contract, nil
	}
}
