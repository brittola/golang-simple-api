# golang-simple-api
API simples, com banco de dados fictício, para conhecer comandos básicos do framework Gin

A API funciona com um CRUD simples e algumas respostas/status baseados em checagem de dados.
Para alimentar o "banco de dados", basta enviar um JSON para a rota post /tweets, respeitando o formato da struct do arquivo /api/entities/Tweet.go.
