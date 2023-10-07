package main

import (
	"fmt"
	"os"

	"github.com/digital_wallet_design/internals/repositiories"
	"github.com/digital_wallet_design/internals/services"
)

func main() {
	accountRepo := repositiories.NewAccountRepoImpl()
	accountService := services.NewAccountService(accountRepo)

	actResp, err := accountService.NewAccount("ABC", 100)
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(actResp.ToString())

	actResp, err = accountService.NewAccount("PQR", 100)
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(actResp.ToString())
	accountService.TransferMoney("ABC", "PQR", 50)
	resp, err := accountService.AccountStatement("ABC")
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.ToString())
	resp, err = accountService.AccountStatement("PQR")
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.ToString())
	accountService.TransferMoney("PQR", "ABC", 50)
	resp, err = accountService.AccountStatement("PQR")
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.ToString())
	resps, err := accountService.Overview()
	if err != nil {
		fmt.Println(" error = %+v\n", err)
		os.Exit(1)
	}
	for _, curResp := range resps {
		fmt.Println(curResp)
	}
}
