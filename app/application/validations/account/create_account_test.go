package validations

import (
	"stoneBanking/app/gateway/web/account/vo/input"
	validations "stoneBanking/app/gateway/web/account/vo/input/validations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateAccountInput(t *testing.T) {
	testCases := []struct {
		name    string
		input   input.CreateAccountVO
		want    input.CreateAccountVO
		wantErr bool
	}{
		{
			name: "dados de entrada corretos, irá passar por todas as validações",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			want: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			wantErr: false,
		},
		{
			name: "dados de entrada com nome vazio, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com cpf vazio, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com cpf incorreto, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395021",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com cpf de tamanho incorreto, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "105453950",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com cpf de teste, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "11111111111",
				Secret:  "123456",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com senha vazia, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "",
				Balance: 0,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
		{
			name: "dados de entrada com saldo negativo, irá retornar dados de input vazios, e erro",
			input: input.CreateAccountVO{
				Name:    "Joao da Silva",
				CPF:     "10545395020",
				Secret:  "123456",
				Balance: -50,
			},
			want:    input.CreateAccountVO{},
			wantErr: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			got, err := validations.ValidateAccountInput(test.input)

			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
