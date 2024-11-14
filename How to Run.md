Steps to run:
1. Have a minikube environment setup and have docker running inside the mikuce cluster:
	eval $(minikube -p minikube docker-env)	
2.  Buld the docker image
	docker build -t receipt-processor .
3. Load the docker image
	minikube image load receipt-processor:latest
4. Run the image by applying
    k apply -f receipt-deployment.yaml 
5. Get your minikube ip
	minikube ip
6. RUN THE CURL COMMANDS:
	curl --location 'http://192.168.49.2:30080/receipts/process' --header 'Content-Type: application/json' --data '{
  "retailer": "Costco",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Shaving Kit 14PK",
      "price": "16.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}'

curl --location 'http://<minikube-ip>:30080/receipts/{id}/points'

k port-forward service/receipt-processor-service 30080:8080 