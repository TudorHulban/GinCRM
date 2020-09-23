# GinCRM

## Persisted structures definition should reside in /pck/persistence package. 

## Test User creation
```
curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/login
```

