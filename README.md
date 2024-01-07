# Receipt API

This repository contains the source code for a simple Receipt API that processes receipts and awards points based on specified rules.

## API Endpoints

### Process Receipts

- **Endpoint:** /receipts/process
- **Method:** POST
- **Payload:** Receipt JSON
- **Response:** JSON containing an id for the receipt.

### Get Points

- **Endpoint:** /receipts/{id}/points
- **Method:** GET
- **Response:** A JSON object containing the number of points awarded.

## Rules

The rules for awarding points to a receipt are outlined in the API specification.

- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of 0.25.
- 5 points for every two items on the receipt.
- If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
- 6 points if the day in the purchase date is odd.
- 10 points if the time of purchase is after 2:00 pm and before 4:00 pm.

## Examples

### Example 1

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
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
}

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

### Example 2
```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}

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


## How to Run

### Using Docker

1. Install Docker on your machine.
2. Clone this repository:

    ```bash
    git clone https://github.com/your-username/receipt-api.git
    ```

3. Navigate to the project directory:

    ```bash
    cd receipt-api
    ```

4. Build the Docker image:

    ```bash
    docker build -t receipt-api .
    ```

5. Run the Docker container:

    ```bash
    docker run -p 8080:8080 receipt-api
    ```

### Local Development

1. Install Go on your machine.
2. Clone this repository:

    ```bash
    git clone https://github.com/your-username/receipt-api.git
    ```

3. Navigate to the project directory:

    ```bash
    cd receipt-api
    ```

4. Build and run the Go application:

    ```bash
    go build -o main .
    ./main
    ```

