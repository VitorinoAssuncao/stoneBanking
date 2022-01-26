package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"strings"

	"github.com/lib/pq"
)

func (repository accountRepository) Create(ctx context.Context, newAccount account.Account) (account.Account, error) {
	const sqlQuery = `
	INSERT INTO
			accounts (name, cpf, secret, balance)
	VALUES
			($1, $2, $3, $4)
	RETURNING
			id, external_id, created_at
	`
	row := repository.db.QueryRow(
		ctx,
		sqlQuery,
		newAccount.Name,
		newAccount.CPF,
		newAccount.Secret,
		newAccount.Balance.ToInt(),
	)

	err := row.Scan(&newAccount.ID, &newAccount.ExternalID, &newAccount.CreatedAt)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" && strings.Contains(err.Error(), "accounts_cpf_key") {
				return account.Account{}, customError.ErrorAccountCPFExists
			}
		}

		return account.Account{}, err
	}

	return newAccount, nil
}
