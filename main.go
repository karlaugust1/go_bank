package main

import (
	"fmt"
	"go_bank/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// conta := ContaCorrente{Titular: "karl", NumeroAgencia: 589, NumeroConta: 31654, Saldo: 111.11}
	// conta2 := ContaCorrente{Titular: "karl", NumeroAgencia: 589, NumeroConta: 31654, Saldo: 111.11}

	// fmt.Println(conta, conta2)
	// fmt.Println(conta == conta2)

	// 	var conta3 *ContaCorrente
	// 	conta3 = new(ContaCorrente)
	// 	conta3.Titular = "chris"
	// 	conta3.Saldo = 500

	// 	var conta4 *ContaCorrente
	// 	conta3 = new(ContaCorrente)
	// 	conta3.Titular = "chris"
	// 	conta3.Saldo = 500

	// fmt.Println(&conta3) //pegar o endereço da memoria
	// fmt.Println(*conta3) //pegar o conteúdo
	// fmt.Println(conta3)
	// fmt.Println(conta3 == conta4)

	// conta := ContaCorrente{}
	// conta.Titular = "Denise"
	// conta.Saldo = 500

	// saque := 200.
	// // conta.Saldo = conta.Saldo - saque
	// fmt.Println(conta.Sacar(saque))
	// fmt.Println(conta.Saldo)

	// fmt.Println(conta.Depositar(saque + 100))
	// fmt.Println(conta.Saldo)

	// -------------------
	// conta := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "denise",
	// 	CPF:       "123.132.123-88",
	// 	Profissao: "Desenvolvedor",
	// }, NumeroAgencia: 111, NumeroConta: 3333, Saldo: 100.}

	// cliente := clientes.Titular{"karl", "123.123.465-88", "lixeiro"}
	// conta2 := contas.ContaCorrente{Titular: cliente, NumeroAgencia: 222, NumeroConta: 4444}

	// status := conta2.Transferir(100., &conta)
	// fmt.Println(status)
	// fmt.Println(conta)
	// fmt.Println(conta2)

	// -------------------
	conta := contas.ContaPoupanca{}
	conta.Depositar(100)
	PagarBoleto(&conta, 55)

	fmt.Println(conta.ObterSaldo())

}
