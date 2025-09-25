package app

//pacote externo que baixou do github
//diferencia o nome do pacote pelo que vem depois da útima barra
import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar com G maiusculo p/ que o main importe essa funcao
//retorno *cli.App (ponteiro)
//Sempre que a func estiver com letra maiuscula no inicio, coloque um comentario acima dessa func

// Gerar vai retornar a linha de comando pronta para ser executada
func Gerar() *cli.App {
	// p/criar uma aplicacao chame uma funcao deste pct
	app := cli.NewApp()
	//configure o nome e p/ q serve
	app.Name = "Aplicação de Linha de Comando"              //pode dar qlqr nome p/ aplicação
	app.Usage = "Busca IPs e Nomes de Servidor na internet" //oq essa aplicacao vai fazer

	flags := []cli.Flag{ //as flags estão aqui, numa variável externa
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	//criando o primeiro comando
	app.Commands = []cli.Command{ //é um slice, posso ter + de 1 comando na aplicacao. O command é do tipo struct
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps, //declarei esta func ali embaixo e chamei ela aqui
		},
		{
			Name:   "Servidores",
			Usage:  "Busca o nome do servidores na internet",
			Flags:  flags,            //aqui as flags são as variáveis q criou lá encima
			Action: buscarServidores, //Precisa criar uma funcao buscarServidores
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	//pegar o host vindo como flag e buscar os IPs
	//Primeiro salvar o valor do shost numa variavel
	host := c.String("host")

	ips, erro := net.LookupIP(host) //+1 pacote -> net
	if erro != nil {                //se der erro, trate!
		log.Fatal(erro) //Log tbm é um pacote
	}

	//variavel ips é um slace, posso iterar por ela(range)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host) //é chamado dentro do pct net (NS name server)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
