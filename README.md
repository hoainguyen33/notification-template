## Project Structure
```
. 
├── cmd                         # start project
    ├── init.go                 # file init application
    ├── main.go                 # file start application
├── config                      # config of application
    ├── config.go               # config with file, env, fangs
    └── config.yaml             # config file
├── common                      # common in project
    ├── flags                   # flags cli app
    ├── errors                  # errors used
    └── {constant}.go           # constants used
├── internal                    # internal project
    ├── delivery                # contact domain and producer, consumer kafka
    ├── domain                  # contact repositorys and services
    ├── repository              # contact database, socket, ...
    └── model                   # model struct
├── pkg                         # all package application used
    └── {package}               # new package
├── proto                       # all proto file
    └── {proto}.pb              # proto API of grpc service
├── ssl                         # ssl key
    ├── instruction.sh          # file create ssl key grpc server
    └── {public_key}.bem        # public keys of services grpc
├── go.mod                      # go modules file
├── go.sum                      # go modules file
├── main.go                     # main file
└── Readme.md                   # you read me here :)
```
