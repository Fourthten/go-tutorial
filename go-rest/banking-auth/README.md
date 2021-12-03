## Source
```bash
https://github.com/ashishjuyal/banking-auth
```

## API
```python
# POST http://localhost:8181/auth/register
# POST http://localhost:8181/auth/login
    {
        "username":  "admin",
        "password": "abc123"
    }
# POST http://localhost:8181/auth/refresh
    {
        "access_token":  "aaaa.bbbb.cccc",
        "refresh_token": "aaaa.bbbb.cccc"
    }
# GET http://localhost:8181/auth/verify?token={token}&routeName={routeNm}&customer_id={custId}&account_id={accId}
```

Run ./start.sh to download the dependencies and run the the application