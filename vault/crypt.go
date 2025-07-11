/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/cybervidia/naka/model"
	"github.com/pterm/pterm"
	"golang.org/x/crypto/argon2"
)

func Unlock(record *model.SecretEntry) {

	pwdInput := pterm.DefaultInteractiveTextInput.WithMask("中")

	pwdDaChiedereAUser, err := pwdInput.Show("中put your seal here")
	if err != nil {
		panic(err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(record.Password)
	if err != nil {
		panic(err)
	}

	iv, err := base64.StdEncoding.DecodeString(record.IV)
	if err != nil {
		panic(err)
	}

	salt, err := base64.StdEncoding.DecodeString(record.Salt)
	if err != nil {
		panic(err)
	}

	key := deriveKey(pwdDaChiedereAUser, salt)

	fmt.Println("unlock key:", key)

	plain, err := decryptAESGCM(ciphertext, iv, key)
	if err != nil {
		panic(err)
	}
	// record.Password = base64.StdEncoding.EncodeToString(plain)
	record.Password = string(plain)
	// record.Password = plain
}

func Lock(record *model.SecretEntry) {

	//chiedo la password all'utente e la copro con 中 per non farla vedere
	pwdInput := pterm.DefaultInteractiveTextInput.WithMask("中")
	pwdDaChiedereAUser, err := pwdInput.Show("中put your seal here")

	if err != nil {
		panic(err)
	}

	salt := make([]byte, 16) // 16 bytes = 128 bit
	_, err = rand.Read(salt)
	if err != nil {
		panic("Errore nella generazione del salt: " + err.Error())
	}

	// === CREA LA CHIAVE AES A 256 BIT ===
	key := deriveKey(pwdDaChiedereAUser, []byte(salt))

	fmt.Println("lock key:", key)
	// === CIFRA LA PASSWORD === qui cifro la password,
	// "record.Password" passando anche la key
	// mi restituisce un array di dati anzi 2,
	//  il testo cifrato e l'initialization vector
	//  tutti e 2 sotto forma di byte
	ciphertext, iv, err := encryptAESGCM([]byte(record.Password), key)
	if err != nil {
		panic(err)
	}

	//faccio l'enconding da binario a base64
	//  per pwd,IV e sale
	// e li metto nel modello (ricorda è un puntatore)
	record.Password = base64.StdEncoding.EncodeToString(ciphertext)
	record.IV = base64.StdEncoding.EncodeToString(iv)
	record.Salt = base64.StdEncoding.EncodeToString(salt)
}

// === qui ci sono metodi di utilità
// === deriveKey prende la password segreta dell'utente,
// === prende il sale (stringa casuale) e genera la chiave

func deriveKey(pass string, salt []byte) []byte {
	key := argon2.IDKey([]byte(pass), salt, 1, 64*1024, 4, 32)
	return key
}

func encryptAESGCM(plaintext []byte, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	iv := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	ciphertext := aesgcm.Seal(nil, iv, plaintext, nil)
	return ciphertext, iv, nil
}

func decryptAESGCM(ciphertext, iv, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, iv, ciphertext, nil)
}
