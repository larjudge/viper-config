# Viper-Config

Viper-Config is a really simple demo application that uses viper to parse both config files and command line flags

Running with a flag: 
```zsh
go run main.go --intConfig=1 --stringConfig="hello world" --boolConfig=true
```

Running with config file:
```zsh
go run main.go --config "$PWD/config/local.yaml"
```

Running with both:
```zsh
go run main.go --config "$PWD/config/local.yaml" --stringConfig="yaml and flags"
```