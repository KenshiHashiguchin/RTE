package model

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

/**
ユーザ署名検証
*/
type userSignature struct {
	address   string
	signature string
	token     string
}

//const rowMessage = "Sign in(token:%s)"
const rowMessage = "Signature for login authentication(token:%s)"

func NewUserSignature(addr, sig, token string) userSignature {
	return userSignature{addr, sig, token}
}

func (userSignature userSignature) Verify() error {
	msg := accounts.TextHash([]byte(fmt.Sprintf(rowMessage, userSignature.token)))

	sig, err := hexutil.Decode(userSignature.signature)
	if err != nil {
		return err
	}
	// revert an offset of Recovery identifier
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return errors.New("invalid signature recovery id")
	}
	sig[crypto.RecoveryIDOffset] -= 27

	pubKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if strings.ToLower(userSignature.address) != strings.ToLower(recoveredAddr.String()) {
		return errors.New("signed by different address")
	}
	return nil
}
