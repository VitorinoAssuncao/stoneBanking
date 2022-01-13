package input

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateLoginInputData(t *testing.T) {
	testCases := []struct {
		name    string
		input   input.CreateAccountVO
		want    bool
		wantErr error
	}{
		{
			name: "with the right input data, gonna return a token, and no error",
			input: input.CreateAccountVO{
				CPF:    "38343335812",
				Secret: "123456789",
			},
			wantErr: nil,
		},
		{
			name: "input data with cpf void, gonna return a error",
			input: input.CreateAccountVO{
				Secret: "123456789",
			},
			wantErr: customError.ErrorAccountCPFRequired,
		},
		{
			name: "input data with secret void, gonna return a error",
			input: input.CreateAccountVO{
				CPF: "12345678912",
			},
			wantErr: customError.ErrorAccountSecretRequired,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			err := ValidateLoginInputData(test.input)

			assert.Equal(t, err, test.wantErr)
		})
	}
}
