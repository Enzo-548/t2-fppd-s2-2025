/*
Servidor de Jogo:
• O servidor é responsável por gerenciar uma sessão de jogo e manter o
estado atual do jogo (ex: lista de jogadores, posição atual de cada
jogadores, número de vidas, etc.).
• O servidor NÃO deve manter uma cópia do mapa do jogo e a lógica de
movimentação de personagens e de funcionamento do jogo NÃO deve
ser movida para o servidor.
• O servidor não deve conter interface gráfica.
• As requisições recebidas pelos clientes e as respostas retornadas devem
ser impressas no terminal para permitir depuração durante do jogo.

Requisitos de Comunicação e Consistência
• Toda comunicação é iniciada pelos clientes. O servidor apenas
responde.
• Todas as chamadas de procedimento remoto devem ter tratamento de
erro com reexecução automática em caso de falha.
• É implementar garantia de execução única (exactly-once) dos comandos
enviados que modificam o estado do servidor:

		//cliente?
	▪ Cada comando pode incluir um sequenceNumber.

	▪ O servidor deve manter o controle de comandos processados por
	cliente para evitar reexecução em caso de retransmissão.
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//STRUCT DA SEÇÃO QUE O SERVIDOR IRA ARMAZENAR
//A IDEIA E ARMAZENAR TODAS AS INFORMAÇÕES RELEVANTES A SECAO NESSA STRUCT E
//ATUALIZAR A SECAO COM AS REQUISICOES
type Session struct{
	pos []int
	/*
	req1
	req2
	...
	*/
}

//STRUCTS DAS REQUISICOES QUE O SERVIDOR IRA LIDAR
//a ideia nao e clara, eu tava pensando de usar um struct assim pra rastrear as
//requisicoes para atualizar o struct session
type Requisitions struct{
	/*
	req1
	req2
	...
	*/
}

//numero maximo de jogadores (alteravel)
const n int = 5

//sessao que o servidor armazena
var sessaoServ Session
//ideia
//vetor de requisicoes;
//e possivel manter a rastreabilidade de todas as requisições baseadas
//no limite de jogadores max da sessao
var requisitionsServ [n]Requisitions

//metodos exportados
type Arith struct{}

//ATUALIZA A SECAO DO JOGO
//importante para atualizar a sessao do servidor
//A estrutura desses metodos é de um tipo para exportar, o primeiro pointer
//e dois parametros, o parametro para providenciado pelo caller(cliente) e o parametro de reply(servidor)
//neste caso ele ta mandandando a secao do cliente que vai ser usada para atualizar o servidor e
//o reply seria uma mensagem de string?(a ser definido)
func (a *Arith) atualizaSessao(sessionsCliente *Session, reply *string) error{
	//TODO
	*&sessaoServ = *sessionsCliente
	return nil
}

//PEGA A SESSAO DO JOGO
//PARA ATUALIZAR OS DADOS NO CLIENTE
//o pointer que ele recebe e para atualizar a sessao no cliente
//pode e deve ser alterado para outro tipo pois o cliente nao vai armazenar a sessao inteira
//tambem levar em consideracao o reply deste metodo, p metodo acima poderia retornar a sessao do cliente ja atualizada
func (a *Arith) pegaSessao(sessaoCliente *Session, reply *string) error{
	//TODO
	*&sessaoCliente = &sessaoServ
	return nil
}

//ideia
//FUNCAO PARA LIDAR COM CONFLITOS DE REQUISICOES(tipo conflito de posicao de jogadores)
func (a *Arith) sessaoException(){

}


//Main exemplo
func main(){
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil{
		fmt.Println("Erro ao registrar o serviço", err)
		return
	}
	//escuta na porta TCP 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil{
		fmt.Println("Erro ao escutar:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor RPC esperando chamadas na porta 1234...")
	rpc.Accept(listener)
}