To start
    1)docker build -t golang-api .
    2)docker run -p 8080:8000 golang-api

API Endpoits:
    localhost:8080/account (POST)
    localhost:8080/account (PUT)
    localhost:8080/account (DELETE)
    localhost:8080/account (GET)
    localhost:8080/account/{code} (GET) (Employee Code)


    Main structure:
    {
        "name": "Gurgen",
        "age": 22,
        "employeeCode": 2332
    }