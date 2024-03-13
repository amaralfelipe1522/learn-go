# Canais

- utilizado para transmitir dados entre Goroutines;
- Canais são o jeito certo de fazer sincronização e código concorrente;
- Servem para coordenar e orquestrar;
- É como numa corrida de revezamento, o corredor com o bastão deve entregar ao próximo, ou seja, não pode numa unica função pegar e entregar o dado. Para isso deve se criar Goroutines;

## Canais direcionais

- Define quais canais apenas recebem ou enviam informações;
- Mecanismos de checagem de tipagem impedirão que um dado seja escrito em um channel de leitura
- É possível converter um canal bidirecional para algum específico, ao contrario, não
