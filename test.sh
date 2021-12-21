curl -c cookie.txt -X POST -H "Content-Type: application/json" -d '{"email":"testtest@testmail.com", "name":"test tarou", "password": "paaaaassword"}' localhost:8080/girack/v1/register
curl -c cookie.txt -X POST -H "Content-Type: application/json" -d '{"email":"testtest@testmail.com", "password": "paaaaassword"}' localhost:8080/girack/v1/login
