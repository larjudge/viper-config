# Viper-Config

Viper-Config is a really simple demo application that uses viper to parse both config files and command line flags

Running with flags: 
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

We can also use short flags to avoid `spec.foo`
```zsh
go run main.go --config "$PWD/config/local.yaml" -s="short flag"
```

### Finally
If you pay close attention to the yaml file you can see that intConfig is commented out. This gives us the ability to not *have to* write full yaml files if we only want to change a couple of fields