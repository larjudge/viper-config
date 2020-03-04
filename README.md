# Viper-Config

Viper-Config is a really simple demo application that uses viper to parse both config files and command line flags

Running flags: 
```zsh
go run main.go --spec.intConfig=1 --spec.stringConfig="hello world" --spec.boolConfig=true
```

Running with config file:
```zsh
go run main.go --config "$PWD/config/local.yaml"
```

Running with both:
```zsh
go run main.go --config "$PWD/config/local.yaml" --spec.stringConfig="yaml and flags"
```