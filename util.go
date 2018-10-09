package main

import (
	"archive/zip"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func EncodeToString(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

func DecodeFromString(data string) []byte {
	ret, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (dk DeviceCrypto) Print(f *os.File) {
	_, _ = f.WriteString("PublicKey : " + dk.PublicKey)
	_, _ = f.Write([]byte{0x0A, 0x0D})
	_, _ = f.WriteString("PrivateKey: " + dk.PrivateKey)
	_, _ = f.Write([]byte{0x0A, 0x0D})
	_ = f.Sync()
}

func (hash CanonicalHash) Stringify() string {
	key := hash.Key
	if len(key) == 0 {
		key = EmptyKey()
	}
	return fmt.Sprintf("%05d %s %s %s", hash.Sequence, key, hash.Gen, hash.Owner)
}

func (hash CanonicalHash) StringifyWithDigest() string {
	key := hash.Key
	if len(key) == 0 {
		key = EmptyKey()
	}
	return fmt.Sprintf("%05d %s %s %s %s", hash.Sequence, hash.Token, key, hash.Gen, hash.Owner)
}

func (hash CanonicalHash) Print() {
	key := hash.Key
	if len(key) == 0 {
		key = EmptyKey()
	}
	verified := "NotVerified"
	if hash.Verified {
		verified = "Verified"
	}
	fmt.Println(
		fmt.Sprintf("%05d", hash.Sequence),
		" ", hash.Token,
		" ", key,
		" ", hash.Gen,
		" ", hash.Owner,
		" ", verified)
}

func (tok *Token) Print() {
	for temp := tok.List.Back(); temp != nil; temp = temp.Prev() {
		temp.Value.(*CanonicalHash).Print()
	}
}

func (store *HashStore) Print(hash string) {
	fmt.Println("Token: ", hash)
	hashes := store.IndexToken[hash]
	if hashes != nil {
		for temp := hashes.Back(); temp != nil; temp = temp.Prev() {
			temp.Value.(*CanonicalHash).Print()
		}
	}
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	for _, file := range files {
		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = file
		header.Method = zip.Deflate
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		rc.Close()
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}
		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			_, err = io.Copy(outFile, rc)
			outFile.Close()
			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}
