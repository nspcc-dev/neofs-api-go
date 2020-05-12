package service

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestSignedDataFromReader(t *testing.T) {
	// nil SignedDataReader
	_, err := SignedDataFromReader(nil)
	require.EqualError(t, err, ErrNilSignedDataReader.Error())

	rdr := &testSignedDataReader{
		testSignedDataSrc: new(testSignedDataSrc),
	}

	// make reader to return an error
	rdr.err = errors.New("test error")

	_, err = SignedDataFromReader(rdr)
	require.EqualError(t, err, rdr.err.Error())

	// remove the error
	rdr.err = nil

	// fill the data
	rdr.data = testData(t, 10)

	res, err := SignedDataFromReader(rdr)
	require.NoError(t, err)
	require.Equal(t, rdr.data, res)
}
