
Simple get all
```curl http://localhost:8080/books```

Get specific book by id
```curl http://localhost:8080/books/4```

Post/create book
```
curl http://localhost:8080/books \
  --include \
  --header "Content-Type: applicaiton/json" \
  --request "POST" \
  --data '{"id": "4", "title": "Neuromancer", "author": "William Gibson", "price": 9.99}'
```