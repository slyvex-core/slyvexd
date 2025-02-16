package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// SlyvexMainnetPrivate is the version that is used for
// slyvex mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var SlyvexMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// SlyvexMainnetPublic is the version that is used for
// slyvex mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var SlyvexMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// SlyvexTestnetPrivate is the version that is used for
// slyvex testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var SlyvexTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// SlyvexTestnetPublic is the version that is used for
// slyvex testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var SlyvexTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// SlyvexDevnetPrivate is the version that is used for
// slyvex devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var SlyvexDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// SlyvexDevnetPublic is the version that is used for
// slyvex devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var SlyvexDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// SlyvexSimnetPrivate is the version that is used for
// slyvex simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var SlyvexSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// SlyvexSimnetPublic is the version that is used for
// slyvex simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var SlyvexSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case SlyvexMainnetPrivate:
		return SlyvexMainnetPublic, nil
	case SlyvexTestnetPrivate:
		return SlyvexTestnetPublic, nil
	case SlyvexDevnetPrivate:
		return SlyvexDevnetPublic, nil
	case SlyvexSimnetPrivate:
		return SlyvexSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case SlyvexMainnetPrivate:
		return true
	case SlyvexTestnetPrivate:
		return true
	case SlyvexDevnetPrivate:
		return true
	case SlyvexSimnetPrivate:
		return true
	}

	return false
}
