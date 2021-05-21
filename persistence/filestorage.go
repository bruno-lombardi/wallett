package persistence

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Compress(compressor string, data []byte) ([]byte, error) {
	buffer := &bytes.Buffer{}
	var writer io.Writer
	var err error = nil

	switch compressor {
	case "flate":
		writer, err = flate.NewWriter(buffer, flate.DefaultCompression)
	case "gzip":
		writer = gzip.NewWriter(buffer)
	}
	_, err = writer.Write(data)
	if c, ok := writer.(io.Closer); ok {
		c.Close()
	}

	return buffer.Bytes(), err
}

func Decompress(compressor string, compressed []byte) ([]byte, error) {
	var reader io.Reader
	var err error = nil

	switch compressor {
	case "flate":
		reader = flate.NewReader(bytes.NewReader(compressed))
	case "gzip":
		reader, err = gzip.NewReader(bytes.NewReader(compressed))
	}

	data, err := ioutil.ReadAll(reader)
	return data, err
}

func EncodeToBytes(e interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(e)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}

func ReadAndDecodeFile(path string, e interface{}) error {
	if file, err := os.Open(path); err != nil {
		return err
	} else {
		defer file.Close()
		compressedBytes, err := ioutil.ReadAll(file)
		byteEncodedData, err := Decompress("flate", compressedBytes)

		decoder := gob.NewDecoder(bytes.NewReader(byteEncodedData))
		err = decoder.Decode(e)
		return err
	}
}

func WriteAndEncodeFile(path string, e interface{}) error {
	if file, err := os.Create(path); err != nil {
		return err
	} else {
		defer file.Close()
		byteEncodedData := EncodeToBytes(e)
		compressedBytes, err := Compress("flate", byteEncodedData)
		_, err = file.Write(compressedBytes)
		return err
	}
}
