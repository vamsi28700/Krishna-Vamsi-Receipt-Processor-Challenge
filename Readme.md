# Receipt Processor API with Membership-based Points System

## File Structure
```
receipt-processor/
├── main.go
├── handlers/
│   └── receipt_handler.go
├── models/
│   └── receipt.go
├── utils/
│   └── points_calculator.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Initial Setup if done from start
  1. Initialize mod file
      `go mod init`
  2. Initializing main
      `go run main.go`

  3. If missed any packages run
      `go mod download`


## Running the Application

1. Building the docker container
`docker build -t receipt-processor .`

2. Running docker container
`docker run -p 8080:8080 receipt-processor`

3. Getting container id:
`docker ps`

4. (Optional) To clear cache
`docker builder prune`

## Process Receipts (API Endpoints)

After setting up docker, execute this all in the same or different command prompt in the system.

    1. Getting receipt id for first example:

`curl -X POST -H "Content-Type: application/json" -d "{\"retailer\": \"Target\", \"purchaseDate\": \"2022-01-01\", \"purchaseTime\": \"13:01\", \"items\": [{\"shortDescription\": \"Mountain Dew 12PK\", \"price\": \"6.49\"}, {\"shortDescription\": \"Emils Cheese Pizza\", \"price\": \"12.25\"}, {\"shortDescription\": \"Knorr Creamy Chicken\", \"price\": \"1.26\"}, {\"shortDescription\": \"Doritos Nacho Cheese\", \"price\": \"3.35\"}, {\"shortDescription\": \"Klarbrunn 12-PK 12 FL OZ\", \"price\": \"12.00\"}], \"total\": \"35.35\"}" http://localhost:8080/receipts/process`


  1.1 getting points using receipt id:

`curl http://localhost:8080/receipts/{id}/points`


    2. Getting receipt id for second example:

`curl -X POST ^ -H "Content-Type: application/json" ^ -d "{\"retailer\": \"M^&M Corner Market\", \"purchaseDate\": \"2022-03-20\", \"purchaseTime\": \"14:33\", \"items\": [{\"shortDescription\": \"Gatorade\",\"price\": \"2.25\"},{\"shortDescription\": \"Gatorade\",\"price\": \"2.25\"},{\"shortDescription\": \"Gatorade\",\"price\": \"2.25\"},{\"shortDescription\": \"Gatorade\",\"price\": \"2.25\"}], \"total\": \"9.00\"}" ^ http://localhost:8080/receipts/process`

  2.1 getting points for second example using receipt id:

`curl http://localhost:8080/receipts/{id}/points`



    3. getting receipt id for third example:
`curl -X POST -H "Content-Type: application/json" -d "{\"retailer\": \"Walgreens\", \"purchaseDate\": \"2022-01-02\", \"purchaseTime\": \"08:13\", \"items\": [{\"shortDescription\": \"Pepsi - 12-oz\", \"price\": \"1.25\"}, {\"shortDescription\": \"Dasani\", \"price\": \"1.40\"}], \"total\": \"2.65\"}" http://localhost:8080/receipts/process`

  3.1 getting points for thrid example using receipt id:

`curl http://localhost:8080/receipts/{id}/points`
