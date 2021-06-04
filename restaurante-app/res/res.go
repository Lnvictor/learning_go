package res

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
	"strconv"
)


type Pedido struct{
	numero int64
	author string
	prato string
	mesa int64
}

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

/*
Lê pedidos de um arquivo e printa na tela
*/
func GetPedidos(){
	pedidos := carregaPedidos()

	for _, pedido := range pedidos{
		fmt.Println("Pedido nº: " + strconv.Itoa(int(pedido.numero)) + "\n\t Nome: "+ pedido.author + "\n\t Prato: " + pedido.prato + "\n\t Mesa:" + strconv.Itoa(int(pedido.mesa)))
	}
}

/*
Cria pedido e escreve no arquivo de pedidos
*/
func MakePedido(){
	pedidos := carregaPedidos()
	var current int64

	if len(pedidos) > 0{
			last := pedidos[len(pedidos) -1]
		current = last.numero + 1
	}else{
		current = 1
	}

	var clientName, prato string
	var numMesa int64

	fmt.Println("Qual o seu nome?")
	fmt.Scanln(&clientName)

	fmt.Println("Digite o prato:")
	fmt.Scanln(&prato)

	fmt.Println("Digite o numero da mesa")
	fmt.Scanln(&numMesa)

	writePedido(Pedido{current, clientName, prato, numMesa})
}

/*
Remove pedido do arquivo de pedidos
*/
func RemovePedido(){
	pedidos := carregaPedidos()
	file, err := os.Create("./data/pedidos.txt")
	check(err)

	for index, pedido := range pedidos{
		if index > 0{
			line := strconv.Itoa(int(pedido.numero)) + "," + pedido.author + "," + pedido.prato + "," + strconv.Itoa(int(pedido.mesa)) + "\n"
			n, err := file.WriteString(line)
			check(err)
			fmt.Println(n)
			writer := bufio.NewWriter(file)
			writer.Flush()
		}
	}
}

func carregaPedidos() []Pedido{
	file, err := os.Open("./data/pedidos.txt")

	check(err)

	defer file.Close()
	var pedidos []Pedido
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		infos := strings.Split(line, ",")
		numero, err := strconv.ParseInt(infos[0], 10, 64)
		check(err)
		author := infos[1]
		prato := infos[2]
		mesa, err := strconv.ParseInt(infos[3], 10, 64)

		check(err)

		pedidos = append(pedidos, Pedido{numero, author, prato, mesa})
	}

	return pedidos
}

func writePedido(pedido Pedido) bool{
	file, err := os.OpenFile("./data/pedidos.txt", os.O_APPEND|os.O_WRONLY, 0664)

	check(err)

	defer file.Close()

	writer := bufio.NewWriter(file)
	var line string
	line = strconv.Itoa(int(pedido.numero)) + "," + pedido.author + "," + pedido.prato + "," + strconv.Itoa(int(pedido.mesa)) + "\n"
	n, err := file.WriteString(line)
	check(err)
	fmt.Println(n)
	writer.Flush()
	return true
}
