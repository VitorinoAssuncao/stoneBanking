package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) GetByCPF(ctx context.Context, accountCPF string) (account.Account, error) {

	const sqlQuery = `
	SELECT 
		id, external_id, name, cpf, secret, balance, created_at
	FROM
		accounts
	WHERE
			cpf = $1
	`
	result := repository.db.QueryRow(
		sqlQuery,
		accountCPF,
	)

	var newAccount = account.Account{}

	err := result.Scan(&newAccount.ID, &newAccount.ExternalID, &newAccount.Name, &newAccount.CPF, &newAccount.Secret, &newAccount.Balance, &newAccount.CreatedAt)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
