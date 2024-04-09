# go-micro
Um software de comunicação entre microsserviços usando golang que possui:
  - Serviço de front-end para mostrar a parte interativa do site e conectar com outros 5 microsserviços;
  - Serviço de autenticação com um banco de dados Postgres;
  - Serviço de Logs com um banco de dados MongoDB;
  - Serviço que escutará mensagens vindas do RabbitMQ e fará ações com elas;
  - Serviço "Broker" a conexão entre o Producer e o Consumer;
  - Serviço de email que recebe uma mensagem JSON, formata ela para a estrutura apropiada e envia.
