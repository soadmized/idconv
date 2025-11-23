package cmd

import (
	"encoding/base64"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_decodeUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		enc     *base64.Encoding
		want    uuid.UUID
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "decode base64",
			input:   "cfRVuteIRnO7acfaTVoHwg==",
			enc:     base64.StdEncoding,
			want:    uuid.MustParse("71f455ba-d788-4673-bb69-c7da4d5a07c2"),
			wantErr: assert.NoError,
		},
		{
			name:    "decode raw base64",
			input:   "rTvwwl3ER8KCCRnqcsuMzQ",
			enc:     base64.RawStdEncoding,
			want:    uuid.MustParse("ad3bf0c2-5dc4-47c2-8209-19ea72cb8ccd"),
			wantErr: assert.NoError,
		},
		{
			name:    "decode uuid",
			input:   "71f455ba-d788-4673-bb69-c7da4d5a07c2",
			enc:     base64.StdEncoding,
			want:    uuid.Nil,
			wantErr: assert.Error,
		},
		{
			name:    "decode invalid string",
			input:   "dfgsdghdfghdfgh",
			enc:     base64.StdEncoding,
			want:    uuid.Nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := decodeUUID(tt.input, tt.enc)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "convert uuid",
			input:   "2a9b3ce1-8dd7-4dc7-9e72-523f1d189755",
			wantErr: assert.NoError,
		},
		{
			name:    "convert b64",
			input:   "Fhzz7an2RyaooXIfe2Saxw==",
			wantErr: assert.NoError,
		},
		{
			name:    "convert raw b64",
			input:   "FrvKxa32ShCAeO/9PyVRSQ",
			wantErr: assert.NoError,
		},
		{
			name:    "random cmd",
			input:   randomCmd,
			wantErr: assert.NoError,
		},
		{
			name:    "convert invalid string",
			input:   "test___",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.wantErr(t, Convert(tt.input))
		})
	}
}
