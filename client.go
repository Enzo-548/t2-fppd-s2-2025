/*
Cliente do Jogo:
• O cliente possui a interface onde o jogador interage com o jogo e controla
toda a lógica de movimentação do personagem e de funcionamento do
jogo.
• Ele se conecta ao servidor para obter o estado atual do jogo (ex: lista de
jogadores e a posição de cada um) e envia comandos de movimento e
interação ao servidor.
• Deve possuir uma thread (goroutine) dedicada para buscar
periodicamente atualizações do estado do jogo no servidor e atualizar o
seu estado local.

Requisitos de Comunicação e Consistência
• Toda comunicação é iniciada pelos clientes. O servidor apenas
responde.
• Todas as chamadas de procedimento remoto devem ter tratamento de
erro com reexecução automática em caso de falha.
• É implementar garantia de execução única (exactly-once) dos comandos
enviados que modificam o estado do servidor:
▪ Cada comando pode incluir um sequenceNumber.
▪ O servidor deve manter o controle de comandos processados por
cliente para evitar reexecução em caso de retransmissão.

*/
package main

import(
	"fmt"
	"net/rpc"
	"os"
)


//main exemplo
//deve reconsilhar com a main rpc para interagir com o servidor
func main(){
	/*
	
	*/

	//conecta ao servidor rpc
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil{
		fmt.Println("Erro ao conectar ao servidor", err)
		return
	}

	defer client.Close()

	//prepara os argumentos para inicializar o jogo
	interfaceIniciar()
	defer interfaceFinalizar()

	// Usa "mapa.txt" como arquivo padrão ou lê o primeiro argumento
	mapaFile := "mapa.txt"
	if len(os.Args) > 1 {
		mapaFile = os.Args[1]
	}

	// Inicializa o jogo
	jogo := jogoNovo()
	if err := jogoCarregarMapa(mapaFile, &jogo); err != nil {
		panic(err)
	}

	// Desenha o estado inicial do jogo
	interfaceDesenharJogo(&jogo)

	// Loop principal de entrada
	for {
		evento := interfaceLerEventoTeclado()
		if continuar := personagemExecutarAcao(evento, &jogo); !continuar {
			break
		}
		interfaceDesenharJogo(&jogo)
	}


	//chama os metodos remotos
	/*
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil{
		fmt.Println("Erro na chamada pdc:", err)
		return
	}
	fmt.Println("Resultado da multiplicação: ", result)
	*/
}