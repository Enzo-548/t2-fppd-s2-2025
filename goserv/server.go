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
	"io"
	"log"
	"net"
	"time"
)

//CONTEUDO DA MENSAGEM
//-- ALTERAR A CADA ATRIBUTO ADICIONADO
//digamos que eu quero adicionar outro atributo pra mandar pro servidor por exemplo
type Payload struct{
	minhaPos PosJogador
}

//STRUCT PARA TRANSPORTE DE MENSAGEM, mais versatil que um canal
//DE ONDE QUE VEIO A MENSAGEM E O CONTEUDO DA MENSAGEM(info da sessao)
type Message struct{
	from string
	payload []byte
	//payload Payload
}

//STRUCT DE DEF Do SERVIDOR
type Server struct{
	listenAddr string
	ln net.Listener
	//canal para desconectar do servidor
	quitch chan struct{}
	msgch chan Message
	//peerMap [netAddr] rastreamentode conexões
}

//Criacao de server
func NewServer(listenAddr string) *Server{
	return &Server{
	listenAddr: listenAddr,
	quitch: make(chan struct{}),
	msgch: make(chan Message, 10),
	}
}

//Init de server
func (a *Server) Start() error{
	ln, err := net.Listen("tcp", a.listenAddr)
	if err != nil{
	return err
	}
	defer ln.Close()	
	a.ln = ln
	
	go a.acceptLoop()

	<-a.quitch
	close(a.msgch)

	return nil
}

//Ciclo de aceitacao de conexão
func (a *Server) acceptLoop(){
	for{
	conn, err := a.ln.Accept()
	if err != nil{
		fmt.Println("accept error: ", err)
		continue
		}

	fmt.Println("new connection to the server accepted: ", conn.RemoteAddr())
	
	go a.readLoop(conn)
	}
}

//Leitura de mensagem
func (a *Server) readLoop(conn net.Conn){
	defer conn.Close()
	for{
	//ALTERAR AQUI PARA A LEITURA DA MENSAGEM
	buf := make([]byte, 2048)

		n,err := conn.Read(buf)
		if err != nil{
		fmt.Println("read error: ", err)
			if err == io.EOF{
				break
			}
		continue
		}
		a.msgch <- Message{
			from:	 conn.RemoteAddr().String(),
			payload: buf[:n],
		}

		conn.Write([]byte("thank you for your message!"))
	}			
}

type PosJogador struct{
	posX int;
	posY int;
}

//STRUCT DA SEÇÃO QUE O SERVIDOR IRA ARMAZENAR
//A IDEIA E ARMAZENAR TODAS AS INFORMAÇÕES RELEVANTES A SECAO NESSA STRUCT E
//ATUALIZAR A SECAO COM AS REQUISICOES
type Session struct{
	listPosJogador []PosJogador
	//Informações relevantes
}

//STRUCTS DAS REQUISICOES QUE O SERVIDOR IRA LIDAR
//CORPO DAS REQUISICOES QUE O CLIENTE IRA PEDIR
//ex: minha posicao!, meu mapa!, onde estão os outros players?, etc.
type Requisition struct{
	listPosJogadores []PosJogador
	//+informações relevantes
}

//numero maximo de jogadores (alteravel)
const n int = 5

//sessao que o servidor armazena
var sessaoServ Session
//ideia
//vetor de requisicoes;
//e possivel manter a rastreabilidade de todas as requisições baseadas
//no limite de jogadores max da sessao
var requisitionsServ [n]Requisition

//metodos exportados usados pelo cliente
//a logica vai ser implementada em servidor
//ex: quero atualizar a minha secao!
type ClientMet struct{}

//ATUALIZA A SECAO DO JOGO
//importante para atualizar a sessao do servidor
//A estrutura desses metodos é de um tipo para exportar, o primeiro pointer
//e dois parametros, o parametro para providenciado pelo caller(cliente) e o parametro de reply(servidor)
//neste caso ele ta mandandando a secao do cliente que vai ser usada para atualizar o servidor e
//o reply seria uma mensagem de string?(a ser definido)
func (a *ClientMet) AtualizaSessao(sessionsCliente *Session, reply *string) error{
	//TODO
	return nil
}

//PEGA A SESSAO DO JOGO
//PARA ATUALIZAR OS DADOS NO CLIENTE
//o pointer que ele recebe e para atualizar a sessao no cliente
//pode e deve ser alterado para outro tipo pois o cliente nao vai armazenar a sessao inteira
//tambem levar em consideracao o reply deste metodo, p metodo acima poderia retornar a sessao do cliente ja atualizada
func (a *ClientMet) PegaSessao(sessaoCliente *Session) error{
	//TODO
	fmt.Println("buscando atualizações do servidor em", time.Now().Format("15:04:05.000"))
	return nil
}

//ideia
//FUNCAO PARA LIDAR COM CONFLITOS DE REQUISICOES(tipo conflito de posicao de jogadores)
func (a *ClientMet) SessaoException(){

}


//Main exemplo
func main(){
	server := NewServer(":3000")

	go func(){
	for msg := range server.msgch{
		fmt.Printf("received message from connection(%s):%s\n", msg.from, string(msg.payload))
		}
	}()

	log.Fatal(server.Start())
}