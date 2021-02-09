package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferentHash(t *testing.T) {
	fileName := "file1"
	url := fmt.Sprintf("http://qa-test.dev.imtlab.io/%s", fileName)
	hashToCompare := "eae50c4ed998d1b65dbf97e0071080a0"

	errFile := downloadFile(fileName, url)
	hash, errHash := hashFileMd5(fileName)

	assert.Nil(t, errFile)
	assert.Nil(t, errHash)
	assert.NotNil(t, hash)
	assert.NotEqualValues(t, hash, hashToCompare)
}

func TestDHash(t *testing.T) {
	fileName := "file2"
	url := fmt.Sprintf("http://qa-test.dev.imtlab.io/%s", fileName)
	hashToCompare := "601d80e3b514e23d8c5e5bbe60f289a5"

	errFile := downloadFile(fileName, url)
	hash, errHash := hashFileMd5(fileName)

	assert.Nil(t, errFile)
	assert.Nil(t, errHash)
	assert.NotNil(t, hash)
	assert.EqualValues(t, hash, hashToCompare)
}
