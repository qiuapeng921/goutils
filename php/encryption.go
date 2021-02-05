package php

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
)

// Md5 - Calculate the md5 hash of a string
func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Md5File - Calculates the md5 hash of a given file
func Md5File(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// JsonEncode - Returns the JSON representation of a value
func JsonEncode(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}

// JsonDecode - Decodes a JSON string
func JsonDecode(data []byte, v interface{}) error {

	return json.Unmarshal(data, v)
}

// Pack pack()
func Pack(order binary.ByteOrder, data interface{}) (string, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Unpack unpack()
func Unpack(order binary.ByteOrder, data string) (interface{}, error) {
	var result []byte
	r := bytes.NewReader([]byte(data))
	err := binary.Read(r, order, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Base64Encode - Encodes data with MIME base64
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode - Decodes data encoded with MIME base64
func Base64Decode(s string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	return string(bs), err
}
