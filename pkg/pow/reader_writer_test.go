package pow

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWriter_ByteWriteWorkflow(t *testing.T) {
	pkgData := ([]byte)("efwefwefwef\nwrfw\twefwefwq#$!EF23r3")
	nextData := ([]byte)("3e2r3")

	buf := &bytes.Buffer{}
	writer := wrapWriter(buf)
	err := writer.writeNext(pkgData)
	assert.NoError(t, err)
	err = writer.writeNext(nextData)
	assert.NoError(t, err)

	reader := wrapReader(buf)
	recvData, err := reader.readNext()
	assert.NoError(t, err)
	assert.Equal(t, pkgData, recvData)

	recvData, err = reader.readNext()
	assert.NoError(t, err)
	assert.Equal(t, nextData, recvData)
}

func TestReadWriter_ObjectWriteWorkflow(t *testing.T) {
	pkgData := struct {
		Value []byte `json:"value"`
	}{Value: ([]byte)("1234567")}

	buf := &bytes.Buffer{}
	writer := wrapWriter(buf)
	err := writer.write(pkgData)
	assert.NoError(t, err)

	reader := wrapReader(buf)
	tmp := struct {
		Value []byte `json:"value"`
	}{}
	assert.NoError(t, reader.read(&tmp))
	assert.Equal(t, pkgData, tmp)
}
