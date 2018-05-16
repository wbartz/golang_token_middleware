# golang_token_middleware
Middelware para validação de token

Middleware utilizado em todas as conexões http para validar o token JWT.

## Packages utilizados
- github.com/kpango/glg
- github.com/dgrijalva/jwt-go
- github.com/gorilla/mux

## Exemplo de uso
>router := mux.NewRouter()
>
>router.Handle("/", negroni.New(
>    negroni.HandlerFunc(validateTokenMiddleware),
 >   negroni.Wrap(http.HandlerFunc(dashboard.home)),
>)).Methods("GET")