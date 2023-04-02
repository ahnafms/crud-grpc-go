# CRUD GRPC Go
## Preparation
Make sure you have installed the applications listed below :
* Go (https://go.dev/)
* Proto compiler for Go (https://grpc.io/docs/languages/go/quickstart/)

## How To Run
1. Git clone this repository
2. Go to the proto directory and run this code in your terminal
```cmd
protoc *.proto --go_out=. --go-grpc_out=.
```
3. Then you'll have the grpc code generated in pkg/protobuf directory.
4. Setup mysql database configurations in pkg/database/db.go, if you have your own setup feel free to change it based on your own configurations.
5. After that run the application. Make sure to run it in root directory
```Go
go run cmd/app/main.go
```
6. Last but not least, import your proto file to POSTMAN

## Make a request
1. After importing proto file to POSTMAN, you'll have the all the MovieServices automatically generated
<img width="1148" alt="Screenshot 2023-04-02 at 18 49 20" src="https://user-images.githubusercontent.com/108605109/229351163-a01d4c87-aae2-48c6-a747-455b39ea20e2.png">

2. Choose the service that you'd like to test, ex : CreateMovies
<img width="1148" alt="Screenshot 2023-04-02 at 18 49 20" src="https://user-images.githubusercontent.com/108605109/229351412-d957b254-4934-45fd-90f2-b97bf82781bd.png">

3. Click on Example Message and POSTMAN automatically generate example message that you can directly send

4. Example response if success
<img width="1147" alt="Screenshot 2023-04-02 at 19 04 27" src="https://user-images.githubusercontent.com/108605109/229351652-1e3418c8-36d2-4b6c-bbde-c09a426440d3.png">
