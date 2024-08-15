# Api-go 🗒️

<div align="justify">
Este projeto é uma aplicação web simples para gerenciar clientes. A aplicação permite adicionar, editar, visualizar e deletar clientes, com suporte para persistência de dados em arquivo JSON.
</div>

### 💻 Sobre o projeto
---

<div align="justify">
O objetivo principal desta aplicação é fornecer uma ferramenta simples e eficaz para gerenciar informações de clientes, permitindo a adição, edição, visualização e exclusão de registros de maneira intuitiva e acessível. A aplicação foi projetada para ser facilmente utilizada em ambientes pequenos e médios que necessitam de uma solução prática para o gerenciamento de dados de clientes, com uma interface amigável e persistência de dados local.

#### 👁️‍🗨️ Funcionalidades Principais

🔹 Cadastro de Clientes: Preencha o formulário com nome, sobrenome, email, endereço, telefone, cidade e estado para adicionar um novo cliente.

🔹 Edição de Clientes: Clique no ícone de lápis ao lado de um cliente para editar suas informações.

🔹 Exclusão de Clientes: Clique no ícone de lixeira para deletar um cliente da lista.

🔹 Visualização de Clientes: Veja a lista completa de clientes cadastrados.

🔹 Feedback em Tempo Real: Receba mensagens de confirmação no rodapé da página para cada operação realizada.

🔹 Data e Hora: Exibição da data e hora atual no rodapé da aplicação.



</div>

#### 🛠 Tecnologias utilizadas

O projeto é desenvolvido utilizando as seguintes tecnologias e gems:

🔹 Go (Golang): Linguagem de programação utilizada no backend.

🔹 Gorilla Mux: Pacote para roteamento de URL.

🔹 Bootstrap 4.5.2: Framework CSS para o design responsivo da aplicação.

🔹 JavaScript Fetch API: Para realizar requisições AJAX e manipular o DOM dinamicamente.

🔹 HTML/CSS/JavaScript: Utilizados para a interface do usuário.

🔹 JSON: Utilizado para armazenar e manipular os dados dos clientes.

---

#### 💡 Veja!


<br>
🔹Video de demonstração

https://github.com/user-attachments/assets/5ebbd99c-e4ba-4890-a313-6ff43097cbe3




<br>

### Como Rodar a Aplicação
---

Pré-requisitos

 🔹 Go: Certifique-se de ter o Go instalado em sua máquina.
 
 🔹 Git: Para clonar o repositório.



### 📋 Instalação
---

Para executar a aplicação localmente em seu ambiente de desenvolvimento, siga estas etapas:

🔹 Clone o repositório:
  ```bash
git clone git@github.com:AngeloSouza1/api-go.git
```
🔹 Abra o diretório do projeto

```bash
cd api-go
```

🔹 Execute a Aplicação:

```bash
 go run main.go
```
 Acesse a Aplicação:

🔹 Abra seu navegador e vá para http://localhost:8080 para acessar a interface da aplicação.

---

### 🚀 Estrutura do Projeto

Api-Go
   - main.go         
   - clientes.json   
   - index.html      
   - README.md       


API Endpoints

   🔹 GET /clientes: Retorna a lista de todos os clientes.
   
   🔹 POST /clientes: Adiciona um novo cliente.
   
   🔹 GET /clientes/{id}: Retorna um cliente específico pelo ID.
   
   🔹 PUT /clientes/{id}: Atualiza um cliente existente.
   
   🔹 DELETE /clientes/{id}: Remove um cliente pelo ID.


---

###  👁️‍🗨️ Contribuição

Contribuições são bem-vindas! Se você quiser contribuir para o projeto, siga estas etapas:

🔹Faça um fork do projeto.

🔹Crie uma nova branch com a sua feature: git checkout -b minha-feature

🔹Faça commit das suas alterações: git commit -m 'Adicionar nova feature'

🔹Faça push para a branch: git push origin minha-feature

🔹Envie um pull request.

---
### Licença
Este projeto é licenciado sob a MIT License.












