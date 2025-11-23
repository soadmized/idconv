package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

const (
	randomCmd = "random"
	output    = `
 ----------------------------------------------------
|   uuid    |  %s  |
 ----------------------------------------------------
|  base64   |        %s        |
 ----------------------------------------------------
| rawBase64 |         %s         |
 ----------------------------------------------------
`
)

func Convert(input string) error {
	if input == randomCmd {
		id := uuid.New()

		printOutput(id)

		return nil
	}

	if id, err := uuid.Parse(input); err == nil {
		printAllArgs(id, uuidToB64(id), uuidToRawB64(id))

		return nil
	}

	if id, err := decodeUUID(input, base64.RawStdEncoding); err == nil {
		printAllArgs(id, uuidToB64(id), input)

		return nil
	}

	if id, err := decodeUUID(input, base64.StdEncoding); err == nil {
		printAllArgs(id, input, uuidToRawB64(id))

		return nil
	}

	return fmt.Errorf("input %s is not base64, rawBase64 or uuid", input)
}

func decodeUUID(input string, enc *base64.Encoding) (uuid.UUID, error) {
	b, err := enc.DecodeString(input)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.FromBytes(b)
}

func uuidToB64(id uuid.UUID) string {
	return base64.StdEncoding.EncodeToString(id[:])
}

func uuidToRawB64(id uuid.UUID) string {
	return base64.RawStdEncoding.EncodeToString(id[:])
}

func printOutput(id uuid.UUID) {
	printAllArgs(id, uuidToB64(id), uuidToRawB64(id))
}

func printAllArgs(id uuid.UUID, b64, rawB64 string) {
	fmt.Printf(output, id.String(), b64, rawB64)
}
