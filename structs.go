package main

type ContentHeader struct {
	Magic          [4]byte
	Version        uint32
	_              [4]byte // Most likely filler.
	EncryptionType uint8
	_              [3]byte  // more padding
	_              [32]byte // Reserved?
	IV             [16]byte

	// Yes, there is more information available for a header.
	// We only need this much to grab the type and IV.
	// See also, https://wiibrew.org/wiki/WiiConnect24/WC24_Content
}

var ContentHeaderMagic = []byte("WC24")

type PubkFormat struct {
	RSAPublic [256]byte
	_         [256]byte // "RSA Reserved"
	AESKey    [16]byte
	_         [16]byte // "AES Reserved"
}
