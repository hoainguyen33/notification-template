## Project Structure
```
. 
├── controller                  # all controller in here
    ├── static.go               # handle generator file (ex: qrcode, barcode)
    ├── template.go             # handle template file
    └── {file}.go               # handle all request
├── crontab                     # we save project crontab in here
    ├── crontab.go              # management all crontab
    └── {file}.go               # all file crontab
├── firebase                    # we save project firebase notification in here
    ├── firebase.go             # management all firebase
    └── {file}.go               # all file firebase
├── middleware                  # we save project middleware in here
    ├── auth.go                 # middleware authentication
    ├── base.go                 # middleware handle basesic
    └── cors.go                 # middleware cors domain policy
├── model                       # we save project struct/models in here
    └── {file}.go               # all file model
├── redis                       # we save project struct/models in here
    ├── redis.go                # management all redis
    └── {file}.go               # all file redis
├── repository                  # we save project struct/repository in here
    └── {file}.go               # all file repository
├── router                      # router layer directories. all server router logics is here 
    ├── api.go                  # all router generator auto
    ├── socket.go               # handle router connect socket
    └── api                     # a directory for custom http routers
        └── my_api.go           # handle router custom
├── service                     # request and connections to therd party servers is here
    ├── {file}.go               # handle all process
├── go.mod                      # go modules file
├── go.sum                      # go modules file
├── main.go                     # project started in this file
└── Readme.md                   # you read me here :)
```