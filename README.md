# MANGA BACKEND SERVICE 
![Coverage](https://img.shields.io/badge/Coverage-41.1%25-yellow)


## Feature What I Want Build 
1. RESTFULL API 
2. The Service Can Be Peer 2 Peed so the community can share the resource 
3. Auto Sync 
4. gRPC 


## Feature Need Have Be Done 
  - [ ] Manga 
  - [ ] User 
  - [ ] Chapter 
  - [ ] Upload Images Using Minio 
  - [ ] Intergrate Images Upload Using Minio 
  - [ ] Ads Management if needed
  - [ ] Docs Rest API 
  - [ ] Scrape Some Another comic site 


## Folder Struct
```
  ├── CODE_OF_CONDUCT.md
  ├── CONTRIBUTING.md
  ├── Dockerfile
  ├── LICENSE
  ├── Makefile
  ├── README.md
  ├── SECURITY.md
  ├── cmd
  │   └── main.go
  ├── config
  │   ├── config.go
  │   ├── config.json
  │   └── config_test.go
  ├── cover.cov
  ├── deployment
  │   ├── application
  │   ├── config.yaml
  │   ├── deployment.yaml
  │   ├── ingress.yaml
  │   ├── service-lb.yaml
  │   ├── service-mysql-lb.yaml
  │   └── service.yaml
  ├── go.mod
  ├── go.sum
  ├── handlers
  │   ├── chapter
  │   ├── http.go
  │   ├── manga
  │   ├── middleware
  │   └── users
  ├── internal
  │   ├── core
  │   ├── ports
  │   └── services
  ├── mocks
  │   ├── ChapterRepository.go
  │   ├── ChapterService.go
  │   ├── MangaRepository.go
  │   ├── MangaRoute.go
  │   ├── MangaService.go
  │   ├── UserRepository.go
  │   └── UserService.go
  ├── pkg
  │   ├── dbconn
  │   ├── migrations
  │   ├── password
  │   └── validation
  ├── repositories
  │   └── mysql
  └── utils
      ├── query.go
      └── query_test.go

```
