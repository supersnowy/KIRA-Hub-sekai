package signerkey

import (
	"errors"
	"time"

	"github.com/KiraCore/cosmos-sdk/codec"
	sdk "github.com/KiraCore/cosmos-sdk/types"
	"github.com/KiraCore/sekai/types"
)

// Keeper is an interface to keep signer keys
type Keeper struct {
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
}

// GetSignerKeys return SignerKeys by a curator
func (k Keeper) GetSignerKeys(ctx sdk.Context, curator sdk.AccAddress) []types.SignerKey {

	var signerKeys []types.SignerKey
	var queryOutput = []types.SignerKey{}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("signer_keys")) // TODO: should use iterator instead of this

	k.cdc.MustUnmarshalBinaryBare(bz, &signerKeys)

	for _, signerKey := range signerKeys {
		if signerKey.Curator.Equals(curator) {
			queryOutput = append(queryOutput, signerKey)
		}
	}

	return queryOutput
}

// NewKeeper is a utility to create a keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// UpsertSignerKey create signer key and put it into the keeper
func (k Keeper) UpsertSignerKey(ctx sdk.Context,
	pubKey [4096]byte,
	keyType types.SignerKeyType,
	Permissions []int,
	curator sdk.AccAddress) error {

	var newSignerKeys []types.SignerKey
	now := time.Now()
	unix := now.Unix() // TODO: this won't work as every node has little time differece in unix

	var signerKey = types.NewSignerKey(pubKey, keyType, unix, true, Permissions, curator)

	// Storage Logic
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte("signer_keys"))

	var signerKeys []types.SignerKey
	k.cdc.MustUnmarshalBinaryBare(bz, &signerKeys)
	// TODO: we need to create index that will help us quickly identify keys belonging to specific user.
	// TODO: must add a check to make sure that 2 accounts can't have the same sub-key
	// TODO: navigating around whole signer keys is inefficient, should update it to efficient and make it by sender
	for _, sk := range signerKeys {
		if sk.PubKey == pubKey {
			if keyType == sk.KeyType {
				return errors.New("keyType shouldn't be different for same pub key")
			}
			if sk.Curator.Equals(curator) {
				return errors.New("this key is owned by another curator already")
			}
			newSignerKeys = append(newSignerKeys, signerKey)
		} else if sk.ExpiryTime > unix { // TODO: this is not correct as unix is various per node
			newSignerKeys = append(newSignerKeys, sk)
		}
	}
	// TODO: easily query if sub-key x belongs to account y

	store.Set([]byte("signer_keys"), k.cdc.MustMarshalBinaryBare(newSignerKeys))
	// TODO: should add test for creating / updating after v0.0.5 release.
	return nil
}

// TODO: should add deleteSignerKey after discussion but this should create another directory under transactions folder?
