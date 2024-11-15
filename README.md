## The project is organized as follows:
```
receipt-processor-challenge/
├── cmd/
│   └── server/
│       └── main.go
├── pkg/
│   ├── handlers/
│   │   ├── get_handler.go
│   │   └── post_handler.go
│   └── models/
│       └── data_model.go
├── receipt-deployment.yaml
├── go.mod
└── README.md
```
* cmd/server/main.go: Entry point of the application; sets up the HTTP server and routes.
* pkg/handlers/get_handler.go: Contains the handler for the GET /receipts/{id}/points endpoint.
* pkg/handlers/post_handler.go: Contains the handler for the POST /receipts/process endpoint.
* pkg/models/data_model.go: Defines the data models and in-memory storage.
* receipt-deployment.yaml: Kubernetes deployment and service configuration.
* go.mod: Go module file specifying dependencies.
* README.md: Project documentation (this file).

## Prerequisites
Go: Version 1.22.1 or higher.
Docker: For containerization.
Kubernetes: For deployment (optional).

## Installation and Setup

Clone the Repository:
```
git clone https://github.com/arinjaysalgia/receipt-processor-challenge.git
cd receipt-processor-challenge
```
## Build the Application:
```
go build -o receipt-processor ./cmd/server
```

## Run the Application:
```
./receipt-processor
```
The server will start on port 8080.

## API Endpoints
1. Process Receipt

* Endpoint: POST /receipts/process
* Description: Accepts a receipt in JSON format and returns a unique receipt ID.
* Request Body:json
```
{
  "retailer": "string",
  "purchaseDate": "YYYY-MM-DD",
  "purchaseTime": "HH:MM",
  "items": [
    {
      "shortDescription": "string",
      "price": "string"
    }
  ],
  "total": "string"
}
```
Response:json
```
{
  "Id": "string"
}
```

2. Get Receipt Points
* Endpoint: GET /receipts/{id}/points
* Description: Returns the points awarded for the receipt corresponding to the provided ID.
* Response:json
```
{
  "points": integer
}
```

## Deployment with Docker
Build Docker Image:
```
docker build -t receipt-processor:latest .
```
Run Docker Container:
```
docker run -p 8080:8080 receipt-processor:latest
```
The service will be accessible at http://localhost:8080.

## Deployment with Kubernetes
1. Have a minikube environment setup and have docker running inside the minikube cluster:
```
eval $(minikube -p minikube docker-env)	
```
2.  Buld the docker image
```	
docker build -t receipt-processor .
```
3. Load the docker image to minikube
```
minikube image load receipt-processor:latest
```
4. Run the image by applying
```
k apply -f receipt-deployment.yaml 
```
5. Port forward the service to the host machine:
```
k port-forward service/receipt-processor-service 8080:8080 
```

6. Run the CURL commands from the host Machine:
```
curl --location http://localhost:8080/receipts/process \
  --header "Content-Type: application/json" \
  --data '{
    "retailer": "Costco",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Shaving Kit 14PK",
        "price": "16.49"
      }
    ],
    "total": "35.35"
  }'
```
7. To get the points for the last receipt
```
curl --location 'http://localhost:8080/receipts/{id}/points'
```


## Dependencies
* Gorilla Mux: Router for handling HTTP requests.
* Google UUID: For generating unique receipt IDs.
These dependencies are specified in the go.mod file.

## Limitation
* This API webserver doesn't store processed receipts persistently. That is a reboot will reset all the previous receipts.
* System also doesn't check if the same JSON payload has been processed before as it will process the same data over and over again and create and new receipt Id with POST call at /receipts/process endpoint.


## License
This project is licensed under the MIT License.

## Acknowledgments
This project is based on the receipt processor challenge by Fetch Rewards.
