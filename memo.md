# get
curl localhost:8888/todos

# post
curl -X POST localhost:8888/todos -H "Content-Type: application/json" -d '{"title":"買い物","context":"寿司を買う","limit_date":"2000-08-28 19:02:00"}'

# put
curl -X PUT localhost:8888/todos/1 -H "Content-Type: application/json" -d '{"title":"洗濯","context":"洗濯機をまわす","limit_date":"2001-08-28 08:02:00"}'

# delete
curl -X DELETE localhost:8888/todos/1