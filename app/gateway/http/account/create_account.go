package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"stoneBanking/app/gateway/http/account/vo/input"
	validations "stoneBanking/app/gateway/http/account/vo/input/validations"
	"stoneBanking/app/gateway/http/account/vo/output"
	"stoneBanking/app/gateway/http/response"
)

//@Summary Create a account
//@Description With the data received, validate her and if all is correct, and don't exist a account with that document, create a new account
//@Accept json
//@Produce json
//@Param account body input.CreateAccountVO true "Account Creation Data"
//@Success 200 {object} output.AccountOutputVO
//@Failure	400 {object} response.OutputError
//@Failure 500 {object} response.OutputError
//@Router /accounts [POST]
func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	const operation = "Gateway.Rest.Account.Create"
	resp := response.NewResponse(w)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}
	defer r.Body.Close()

	controller.log.LogInfo(operation, "unmarshal the data to a internal object")
	var accountInput = input.CreateAccountVO{}
	err = json.Unmarshal(reqBody, &accountInput)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	accountInput.CPF = accountInput.CPF.TrimCPF()

	controller.log.LogInfo(operation, "begin the validation of the input data")
	accountInput, err = validations.ValidateAccountInput(accountInput)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.BadRequest(response.NewError(err))
		return
	}

	accountData := accountInput.ToEntity()
	account, err := controller.usecase.Create(r.Context(), accountData)
	if err != nil {
		controller.log.LogError(operation, err.Error())
		resp.InternalError(response.NewError(err))
		return
	}

	accountOutput := output.ToAccountOutput(account)

	resp.Created(accountOutput)
}
