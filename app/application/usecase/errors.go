package usecase

import "errors"

var (
	errorAccountCPFExists  = errors.New("já existe uma conta cadastrada com este CPF")
	errorCreateAccount     = errors.New("erro ao criar nova conta")
	errorAccountIDNotFound = errors.New("conta não localizada, favor validar o ID informado")
)
