# API-hackathon
API for getting the best developments in differents hackathons. Each development comes with its developers. A hackathon will be creted every 5 minutes (You can change that time in your enviroments)

## configuration
> You need to run the creation schema script and then, you can run the inserts script. Once you have the scripts you have to copy to your console the commands that are in 'enviroments', there you can change the string connection according your credentials of your database. You can change the hackathon creation period too
> Once you have all of that configured, you can run in your console:

```sh
$ go mod tidy
$ go run cmd/main.go
```

## valid user
- email: admin@admin.com
- password: admin

## endpoints

- [POST] localhost:8080/v1/login
- [GET] localhost:8080/v1/hackathons
- [GET] localhost:8080/v1/hackathons/top-ten

- ***important*** Please before using the API, read the documentation 
- [GET] http://localhost:8080/swagger/index.html *** documentation ***


## enviroments

```sh
export SQL_CONNECTION="root:@tcp(127.0.0.1:3306)/api_hackathon?parseTime=True"
export ActualizationPeriod="@every 5m"
export ACCESS_SECRET="someaccesssecret"
```