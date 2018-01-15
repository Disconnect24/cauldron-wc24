package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
)

func Extract(file []byte, pubk []byte) ([]byte, error) {
	// The content header is always a certain length, so the file needs to be larger.
	if 0x140 > len(file) {
		return nil, errors.New("file seems too small")
	}

	contentHeader := ContentHeader{}
	// Read up to IV
	headerReadingBuf := bytes.NewBuffer(file[:64])
	err := binary.Read(headerReadingBuf, binary.BigEndian, &contentHeader)
	if err != nil {
		return nil, err
	}

	// This'll return 0 if they're exactly the same.
	if bytes.Compare(ContentHeaderMagic, contentHeader.Magic[:]) != 0 {
		return nil, errors.New("invalid magic")
	}
	log.Print("Version detected: " + fmt.Sprint(contentHeader.Version))
	switch contentHeader.EncryptionType {
	case 0:
		// there's no encryption to do here. just return 0x140 on
		return file[0x140:], nil
		break
	case 1:
		break
	default:
		return nil, errors.New("unknown encryption type given")
	}

	// if we're here, we need to decrypt.
	// let's get this party started!
	if 0x220 != len(pubk) {
		return nil, errors.New("pubk seems too small")
	}
	pubkHeader := PubkFormat{}
	pubkReadingBuff := bytes.NewBuffer(pubk)
	err = binary.Read(pubkReadingBuff, binary.BigEndian, &pubkHeader)
	if err != nil {
		return nil, err
	}

	// Set up decryption using given IV and AES Key
	block, err := aes.NewCipher(pubkHeader.AESKey[:])
	if err != nil {
		return nil, err
	}

	decryptedBuf := make([]byte, len(file)-0x140)

	stream := cipher.NewOFB(block, contentHeader.IV[:])
	// Skip into 0x140 as that's where content starts
	stream.XORKeyStream(decryptedBuf, file[0x140:])
	return decryptedBuf, nil
}
