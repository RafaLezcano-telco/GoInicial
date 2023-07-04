#Unit test

Validar que una funcionalidad ejecute lo esperado.

```go
func multiplicar (a,b int) int {
    return a*b
}
```

```go
func multiplicar(a,b int){
    r:= a*b
    fmt.Printlb(r) //Paquete pertenece al nucleo de go
}
```

#Integration test
Validamos la funcionalidad que depende de un servicio externo.

Por ejemplo: BD, servicio email, comunicaciones (slack, telegram)

Solo es integracion si utilizamos ese servicio externo.

```go
func ControllerSave (db databaseInterface, name string){
    db.Save(name)
}
```

Si me conecto a la BD, entonces estoy realizando un test de integracion

Si yo simulo una BD, entonces estoy realizando un test unitario
