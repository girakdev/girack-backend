# girack-backend
```
docker-compose up
```
#### Usage
ユーザー追加　POST /users  
全ユーザー取得 GET /users
ID指定でユーザー取得 GET /users/$id

curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d {IdealName: ss, RealName: uu} localhost:8080/users
