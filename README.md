# Microservices training Go

The goal of this project is to create 3 microservices, one for clients, one for invoices and one for products in Golang. Each microservice will be in it's own folder with the corresponding name.

## Clients microservice

This endpoint will manage clients and will interact with other microservices. It will expose the endpoint in port **5000**
The endpoints to create for this service are:

### Get all clients

**GET**

**/clients**

*Response:*
```
[
	{
		"Id": 1,
		"Name": "Google"
	},
	...
]
```
### Get client by id

**GET**

**/clients/{id}**

*Response:*
```
{
	"Id": 1,
	"Name": "Google",
	"HasDebt": false
}
```

*Special considerations*

This endpoint will use the invoice endpoint (GET /invoices/client/{id}) to check if the client has debt and will include that information in the response

### Create client

**POST**

**/clients**

*Request data*
```
{
	"Name": "Google"
}
```

*Response:*
```
{
	"Id": 1,
	"Name": "Google",
	"HasDebt": false
}
```

*Special considerations*

The id of the client will be autoincremental determined by the endpoint or storage

### Update client

**PUT**

**/clients/{id}**

*Request data*
```
{
	"Name": "Google"
}
```

*Response:*
```
{
	"Id": 1,
	"Name": "Google",
	"HasDebt": false
}
```

### Delete client

**DELETE**

**/clients/{id}**

## Tips and special considerations
- Use GORM [example usage](https://github.com/mludeiro/go-micro)
- Use in memory database [example usage](https://github.com/mludeiro/go-micro)

## Next steps
- Unit tests
- Docker usage