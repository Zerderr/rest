curl -v -X GET "http://localhost:9001/university?id=4"

curl -v -X PUT "http://localhost:9001/university" -d "{ \"id\": 2, \"name\": \"new\" }"