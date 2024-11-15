The project is organized as follows:

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
├── deployment.yaml
├── go.mod
└── README.md

cmd/server/main.go: Entry point of the application; sets up the HTTP server and routes.
pkg/handlers/get_handler.go: Contains the handler for the GET /receipts/{id}/points endpoint.
pkg/handlers/post_handler.go: Contains the handler for the POST /receipts/process endpoint.
pkg/models/data_model.go: Defines the data models and in-memory storage.
deployment.yaml: Kubernetes deployment and service configuration.
go.mod: Go module file specifying dependencies.
README.md: Project documentation (this file).
Prerequisites
Go: Version 1.22.1 or higher.
Docker: For containerization.
Kubernetes: For deployment (optional).
Installation and Setup
Clone the Repository:

bash
git clone https://github.com/arinjaysalgia/receipt-processor-challenge.git
cd receipt-processor-challenge
Build the Application:

bash
go build -o receipt-processor ./cmd/server
Run the Application:

bash
./receipt-processor
The server will start on port 8080.

API Endpoints
1. Process Receipt
Endpoint: POST /receipts/process

Description: Accepts a receipt in JSON format and returns a unique receipt ID.

Request Body:

json
 
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
Response:

json
{
  "Id": "string"
}
2. Get Receipt Points
Endpoint: GET /receipts/{id}/points

Description: Returns the points awarded for the receipt corresponding to the provided ID.

Response:

json

{
  "points": integer
}

Deployment with Docker
Build Docker Image:

bash

docker build -t receipt-processor:latest .
Run Docker Container:

bash
docker run -p 8080:8080 receipt-processor:latest
The service will be accessible at http://localhost:8080.

Deployment with Kubernetes
Apply Deployment and Service Configuration:

bash
 
kubectl apply -f receipt-deployment.yaml
Access the Service:

Retrieve the NodePort:

bash
kubectl get services receipt-processor-service
Access the service at http://<NodeIP>:<NodePort>.

Dependencies
Gorilla Mux: Router for handling HTTP requests.
Google UUID: For generating unique receipt IDs.
These dependencies are specified in the go.mod file.

License
This project is licensed under the MIT License.

Acknowledgments
This project is based on the receipt processor challenge by Fetch Rewards.