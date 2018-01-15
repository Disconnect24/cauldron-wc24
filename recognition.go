package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"fmt"
	"crypto/cipher"
	"crypto/aes"
)

	// The content header is always a certain length, so the file needs to be larger.
	if 0x140 > len(file) {
		panic("File seems too small.")
	}

	contentHeader := ContentHeader{}
	// Read up to IV
	headerReadingBuf := bytes.NewBuffer(file[:64])
	err := binary.Read(headerReadingBuf, binary.BigEndian, &contentHeader)
	if err != nil {
		panic(err)
	}

	if ContentHeaderMagic != string(contentHeader.Magic[:]) {
		panic("Invalid magic!")
	}
	log.Print("Version detected: " + fmt.Sprint(contentHeader.Version))
	switch contentHeader.EncryptionType {
	case 0:
		panic("I don't support plaintext extraction, but I'm pretty sure you can yourself w/ hex editor. ;P")
		break
	case 1:
		break
	default:
		panic("Unknown WC24 container version.")
	}

	// if we're here, we need to decrypt.
	// let's get this party started!
	if 0x220 != len(pubk) {
		panic("pubk seems too small.")
	}
	pubkHeader := PubkFormat{}
	pubkReadingBuff := bytes.NewBuffer(pubk)
	err = binary.Read(pubkReadingBuff, binary.BigEndian, &pubkHeader)
	if err != nil {
		panic(err)
	}

	// Set up decryption using given IV and AES Key
	block, err := aes.NewCipher(pubkHeader.AES_key[:])
	if err != nil {
		panic(err)
	}

	decryptedBuf := make([]byte, len(file)-0x140)

	stream := cipher.NewOFB(block, contentHeader.IV[:])
	// Skip into 0x140 as that's where content starts
	stream.XORKeyStream(decryptedBuf, file[0x140:])
	return decryptedBuf
}
