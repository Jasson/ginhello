# ginhello
本项目借鉴了  https://github.com/Jasson/golang-gin-restfulAPI-example-app
用 Redis 替换了MongoDB。

用户注册：
curl -X POST \
  http://127.0.0.1:8000/api/v1/user/register \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: 127.0.0.1:8000' \
  -H 'Postman-Token: e6dee36e-6a89-4266-8172-a18b662a6b47,24278f50-6b46-4d68-a8b2-c55d6df52ed5' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 99' \
  -H 'cookie: mysession=MTU2MjkxMjIzMXxOd3dBTkVFMVFUZEhVVWcyVjAxVFRGUXlUVTlGV1RaWVZrWmFWVXBOVDFOQlF6VkNORXRCVVV4Q1MwOVZRMGhCTWpOSlR6WlFVVkU9fD_1RxA923zXozWHkPJ4D3d1tliyrFNCDyS56GF1cv-C' \
  -b mysession=MTU2MjkxMjIzMXxOd3dBTkVFMVFUZEhVVWcyVjAxVFRGUXlUVTlGV1RaWVZrWmFWVXBOVDFOQlF6VkNORXRCVVV4Q1MwOVZRMGhCTWpOSlR6WlFVVkU9fD_1RxA923zXozWHkPJ4D3d1tliyrFNCDyS56GF1cv-C \
  -d '{
    "Username": "jason1@lenovo.com",
    "Password": "jass",
    "sex": "1",
    "role": "user"
}'
成功：返回 
{
    "code": 200,
    "timestamp": "2019-07-15 09:57:24"
}

用户登录
curl -X POST \
  http://127.0.0.1:8000/api/v1/user/login \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: text/plain' \
  -H 'Host: 127.0.0.1:8000' \
  -H 'Postman-Token: a7dddf7b-ef95-4667-ab1b-11e5d9334c63,537cf5bc-743d-48bf-8189-6b3a2e38da92' \
  -H 'User-Agent: PostmanRuntime/7.13.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 52' \
  -d '{"Username":"jason1@lenovo.com",
 "Password":"jass"}'
 
 成功返回：
 {
    "code": 200,
    "expire": "2019-07-15T10:59:09+08:00",
    "message": "login success!",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMxNTk1NDksImlkIjp7Il9pZCI6IjVkMmJkZDg0NTUyNTBiY2VkOGIxYmRiYSIsInVzZXJuYW1lIjoiamFzb24xQGxlbm92by5jb20iLCJwYXNzd29yZCI6Imphc3MiLCJzZXgiOiIxIiwidXBkYXRlVGltZSI6MTU2MzE1NTg0NCwicm9sZSI6InVzZXIifSwib3JpZ19pYXQiOjE1NjMxNTU5NDl9.vFArR3plTk4QamHDdSO2JmlFEsVot-cEOF0u1I9TPn4"
}
