# POC go-chi
POC de um cadastro de produtos usando go-chi. Objetivo é testar algumas necessidades que tenho de APIs que monto no dia a dia.

# Run

``` 
go run src/main.go 
```

# Docker

- Build

``` 
 docker build -t go-chi-poc:01 .
```

- Run 

``` 
docker run -d -p 3000:3000 go-chi-poc:01
```

# Testes API

- Insert Produto

``` 
Url: http://localhost:3000/product
Method: POST
Authirization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2VudElkIjoxMjEyfQ.uk6SVBcL-YVQWQbVFzvDZU9siiuNXjJq1dcEL6HZPtc
Content-Type: application/json

Body:{
  "name": "TV Philco",
  "obs": "Tem na loja"
}

Response:{
"statusCod": 200,
"message": "Produto <id produto> salvo com sucesso"
}
```


- Get Produto

``` 
Url: http://localhost:3000/product/<id produto>
Method: GET
Authirization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2VudElkIjoxMjEyfQ.uk6SVBcL-YVQWQbVFzvDZU9siiuNXjJq1dcEL6HZPtc
Content-Type: application/json

Response:{
"name": "TV Philco",
"obs": "Tem na loja"
}
```

# Falta Testar

- Implementar um Middleware de Logs no meu padrão

# Referencias

- https://github.com/go-chi/chi
- https://github.com/go-chi/chi/blob/master/middleware/basic_auth.go
- https://github.com/unrolled/render