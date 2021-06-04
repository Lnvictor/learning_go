package main

import (
	"fmt"
	"restaurante-app/res"
	"os"
)

func main(){
	for{
		fmt.Println(" =========Bem vindo ao restaurante do Vitão!!!==========")	
		fmt.Println("|\tOPÇÕES\t\t\t\t\t\t|")
		fmt.Println("|\t1. Visualizar pedidos\t\t\t\t|")
		fmt.Println("|\t2. Fazer um pedido\t\t\t\t|")
		fmt.Println("|\t3. Remover pedido mais antigo\t\t\t|")
		fmt.Println(" =======================================================")

		var choice string
		fmt.Scanln(&choice)

		if choice == "4"{
			fmt.Println("Terminando o programa")
			os.Exit(0)
		}

		paths := map[string]func(){
			"1": res.GetPedidos,
			"2": res.MakePedido,
			"3": res.RemovePedido,
		}

		paths[choice]()
	}
}
