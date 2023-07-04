# Receipt Processor

## main.go

The main.go file contains the entry point of the server. It initializes the router from the generated code and starts the HTTP server on port 8080.

## routers.go

The routers.go file defines the routes and the router. It uses the Gorilla mux package to handle HTTP routing. The NewRouter function creates a new router and registers the routes defined in the routes variable. Each route has a name, HTTP method, URL pattern, and handler function.

## receipt_service.go

The receipt_service.go file contains the handler functions for the /receipts/process and /receipts/{id}/points endpoints. The ProcessReceiptHandler function processes a receipt JSON payload, calculates the points earned based on the receipt data, generates a receipt ID, and returns the ID as a JSON response. The GetPointsHandler function retrieves the points awarded for a given receipt ID and returns the points as a JSON response.

## receipt_processor.go

The receipt_processor.go file defines the ReceiptProcessor struct and its methods. The ReceiptProcessor represents the receipt processing logic. The ProcessReceipt method processes a receipt and calculates the number of points earned based on the specified rules. It uses helper methods to calculate points for different aspects of the receipt, such as the retailer name, total amount, items, etc.

## receipt_processor_test.go

The receipt_processor_test.go file contains unit tests for the ReceiptProcessor methods.

## How to run locally

Step 1: Clone the project

```bash
  git clone git@github.com:Sailesh2577/receipt-processor-challenge-solution.git
```

Step 2: Install go locally using this website https://go.dev/doc/install. Then go to the project directory

```bash
  cd receipt-processor-challenge-solution
```

```bash
  cd go-server-server-generated
```

Step 3: Now build and run the server

```bash
  go build
```

```bash
  ./go-server-server-generated
```

Step 4: The server has started. Now open another tab in the terminal. Navigate to go-server-server-generated folder again and run the following command(you will need to have installed curl. Click the link to install https://everything.curl.dev/get):

```bash
  curl -X POST -H "Content-Type: application/json" -d '{                    ─╯
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}' http://localhost:8080/receipts/process
```

Step 5: Now it will provide you with an ID. Copy that ID and paste it in the following command in the place of {id}:

```bash
curl -X GET http://localhost:8080/receipts/{id}/points
```

An alternative way to run this:

1. Follow above until step 3.
2. Open the browser at http://localhost:8080/.
3. Complete step 4 and an ID will be generated.
4. Then go to the browser again but this time add receipts/{id}/points in front of the http://localhost:8080/ and replace ID with the ID that was generated in the terminal.

## Resources

1. Used swagger.io website to get go-server-server-generated folder setup using api.yml.
2. Used https://go.dev/doc/install to install go
3. Used this website https://eli.thegreenplace.net/2021/rest-servers-in-go-part-3-using-a-web-framework/ to understand what are handlers and how to use them
4. Used this website to https://eli.thegreenplace.net/2021/rest-servers-in-go-part-4-using-openapi-and-swagger/ to understand Rest APIs and swagger.

## Given README.md

Build a webservice that fulfils the documented API. The API is described below. A formal definition is provided
in the [api.yml](./api.yml) file, but the information in this README is sufficient for completion of this challenge. We will use the
described API to test your solution.

Provide any instructions required to run your application.

Data does not need to persist when your application stops. It is sufficient to store information in memory. There are too many different database solutions, we will not be installing a database on our system when testing your application.

## Language Selection

You can assume our engineers have Go and Docker installed to run your application. Go is our preferred language, but it is not a requirement for this exercise.

If you are using a language other than Go, the engineer evaluating your submission may not have an environment ready for your language. Your instructions should include how to get an environment in any OS that can run your project. For example, if you write your project in Javascript simply stating to "run `npm start` to start the application" is not sufficient, because the engineer may not have NPM. Providing a docker file and the required docker command is a simple way to satisfy this requirement.

## Submitting Your Solution

Provide a link to a public repository, such as GitHub or BitBucket, that contains your code to the provided link through Greenhouse.

---

## Summary of API Specification

### Endpoint: Process Receipts

- Path: `/receipts/process`
- Method: `POST`
- Payload: Receipt JSON
- Response: JSON containing an id for the receipt.

Description:

Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt
was awarded.

How many points should be earned are defined by the rules below.

Reminder: Data does not need to survive an application restart. This is to allow you to use in-memory solutions to track any data generated by this endpoint.

Example Response:

```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Endpoint: Get Points

- Path: `/receipts/{id}/points`
- Method: `GET`
- Response: A JSON object containing the number of points awarded.

A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:

```json
{ "points": 32 }
```

---

# Rules

These rules collectively define how many points should be awarded to a receipt.

- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of `0.25`.
- 5 points for every two items on the receipt.
- If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
- 6 points if the day in the purchase date is odd.
- 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Examples

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

```text
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```

---

```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

```text
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
```

---

# FAQ

### How will this exercise be evaluated?

An engineer will review the code you submit. At a minimum they must be able to run the service and the service must provide the expected results. You
should provide any necessary documentation within the repository. While your solution does not need to be fully production ready, you are being evaluated so
put your best foot forward.

### I have questions about the problem statement

For any requirements not specified via an example, use your best judgment to determine the expected result.

### Can I provide a private repository?

If at all possible, we prefer a public repository because we do not know which engineer will be evaluating your submission. Providing a public repository
ensures a speedy review of your submission. If you are still uncomfortable providing a public repository, you can work with your recruiter to provide access to
the reviewing engineer.

### How long do I have to complete the exercise?

There is no time limit for the exercise. Out of respect for your time, we designed this exercise with the intent that it should take you a few hours. But, please
take as much time as you need to complete the work.
