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

Include example receipts and their expected point calculations as described in the API specification.

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

