

### BLOG APP


This codebase implements the clean architecture [uncle BOB]and the use of different repository such as mysql, firebase. It aims to understand the full-fledged go-lang application with RESTFUL API. 

Objectives of applications:
- user authentication
- post CRUD operations

# How it works
```
.
├── common                    
│   ├── database.go           
│   ├── firebase.go 
│   └── utils.go
├── config
│   └── config.go
├── go.mod                    
├── go.sum                    
├── main.go                   
├── Makefile
├── middleware                
│   └── routes.go            
├── post
│   ├── model
│   │   └── postModel.go
│   ├── repo
│   │   └── postRepo.go
│   ├── route
│   │   └── postRoute.go
│   └── usecase
│       └── postUsecase.go
├── public
│   └── favicon.ico
├── README.Md
├── runner.conf
├── serviceAccountKey.json            //firebase service account key
├── tmp
│   └── runner-build
└── user
    ├── model
    │   └── userModel.go
    ├── repository
    ├── route
    │   └── userRoute.go
    └── usecase
        └── userUsecase.go

```

# Getting started

## Install the Golang
https://golang.org/doc/install

## Environment Con- fig
make sure your bashrc holds the proper gopath and goroot.


## Install Fresh
I used  Fresh for hot-reload 

https://github.com/pilu/fresh
```
go get -u github.com/pilu/fresh
```

## Makefile and start dev-server
I have used makefile as service to run and build the package. 
```
➜  make serve - It run system with hot reload
➜  make dev  -It runs system  

```

## Todo
- Add test module. 
