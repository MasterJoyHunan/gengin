package {{.pkgName}}

type Server struct {
    Name string
    Host string
    Port string
}

type Jwt struct {
    Secret string
    Expire int
}

type Log struct {
    Dir   string
    Level string
}

type Mysql struct{
    Host string
    Port int
    User string
    Pwd  string
    Db   string
}

type Redis struct{
    Host string
    Port int
    User string
    Pwd  string
    Db   string
}