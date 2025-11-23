# RemoteList RPC

Este projeto implementa um serviço de listas remotas utilizando RPC (Remote Procedure Call) em Go. O objetivo é permitir que múltiplos clientes possam criar, manipular e consultar listas de inteiros de forma concorrente, com persistência de dados e tolerância a falhas.

## Funcionalidades

- **Criação e manipulação de listas remotas:**  
  Os clientes podem criar listas identificadas por um ID, adicionar elementos (`append`), remover o último elemento (`remove`), consultar elementos por índice (`get`) e obter o tamanho da lista (`size`).

- **Persistência com snapshot e log:**  
  O estado das listas é salvo periodicamente em um arquivo de snapshot (`snapshot.json`). Todas as operações realizadas também são registradas em um arquivo de log (`operations.log`). Ao reiniciar o servidor, o snapshot é carregado e as operações do log são aplicadas para restaurar o estado mais recente.

- **Concorrência e segurança:**  
  O acesso às listas é protegido por exclusão mútua (mutex), garantindo que múltiplos clientes possam operar de forma segura e consistente.

- **Recuperação de falhas:**  
  Em caso de falha ou reinicialização do servidor, o sistema é capaz de recuperar o estado das listas a partir do snapshot e do log de operações.

## Estrutura do Projeto

- `remotelist_rpc_server.go`: Inicializa o servidor RPC e gerencia conexões dos clientes.
- `remotelist_rpc_client.go`: Cliente que interage com o servidor RPC.
- `pkg/remotelist.go`: Implementa a lógica das listas remotas e suas operações.
- `pkg/persistence.go`: Responsável pela persistência dos dados (snapshot e log).

---

# Discussão sobre Limitações

A solução desenvolvida apresenta características importantes, mas também limitações naturais de um sistema baseado em RPC síncrono e persistência local. A utilização de exclusão mútua garante consistência forte nas operações, porém impõe um gargalo claro: todas as requisições ao servidor competem pelo mesmo mutex. Isso significa que, à medida que mais clientes passam a acessar múltiplas listas simultaneamente, o throughput do sistema não aumenta; pelo contrário, tende a degradar. Portanto, a escalabilidade horizontal é bastante limitada, já que só existe uma instância capaz de atender as solicitações.

Em termos de disponibilidade, o sistema também depende diretamente da integridade de um único processo. Se o servidor cair, todas as operações deixam de ser atendidas. Embora exista um mecanismo de recuperação baseado em snapshot e log, ele não evita indisponibilidade, apenas reduz a perda de dados após a reinicialização. A falha do arquivo de snapshot, corrupção do log ou interrupção durante o salvamento são potenciais pontos críticos.

A consistência, por outro lado, é o ponto mais forte da solução: o uso de mutex garante que todas as operações são aplicadas de forma serializada, preservando o estado das listas. O mecanismo de snapshot seguido da aplicação do log mantém a coerência mesmo após reinicializações, mas não considera falhas durante a criação do snapshot, o que pode gerar inconsistência parcial caso a escrita seja interrompida.

A escalabilidade poderia ser melhorada através da remoção do mutex global, por exemplo substituindo por locks por lista, particionamento (sharding) do estado ou replicação entre múltiplos servidores. No entanto, qualquer melhoria desse tipo reduziria a consistência, seja permitindo condições de corrida entre partições, seja introduzindo replicação assíncrona, o que levaria a inconsistências temporárias. Assim, um sistema mais escalável inevitavelmente impactaria a garantia de consistência e/ou disponibilidade, refletindo o clássico trade-off do teorema CAP.