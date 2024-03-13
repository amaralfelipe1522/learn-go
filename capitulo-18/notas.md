# Concorrência x Paralelismo

- Concorrência é um padrão, uma forma de pensar em executar múltiplas tarefas sem que uma dependa da outra, tanto faz se a arquitetura utilizada tiver 1 ou vários núcleos.
- DOS não permitia tarefas em concorrência. Windows trouxe essa possibilidade.
- Paralelismo ocorre quando o processador possui mais de um núcleo, sendo possível a execução de mais de uma tarefa concorrente ao mesmo tempo.
- A linguagem GO utiliza o máximo de núcleos disponível na arquitetura durante o Run Time, portanto, sempre que possível executará os processos de forma paralela.
- Goroutine (_go func_) determina que a função em questão será executada de forma concorrente e fora de nosso controle. Assim como não temos mais controle de uma flecha a partir do momento que a disparamos.
- sync.WaiGroups serve para que a func main() aguarde o termino do Goroutine a partir de 3 funções principais:
  - func Add: Quantidade de Goroutines;
  - func Done: Encerra;
  - func Wait: Aguarda todo mundo terminar.
- Share by communicating:
  - Programação concorrente em multiplos ambientes pode ser problematico em relação ao uso de variáveis compartilhadas entre elas;
  - O compartilhamento de variáveis em GO é feito a partir de channels;
  - Condição de corrida é uma falha no processo em que o resultado é INESPERADAMENTE dependente de outro processo;
- Yield é quando um unico núcleo consegue executar de forma concorrente múltiplas tarefas, passando a impressão de ser em paralelo. Para tratar isso utilizasse Mutex ou Channel:
  - Mutex para impedir que função B utilize a variável compartilhada que está em processo de leitura e escrita pela função A;
  - Channels
