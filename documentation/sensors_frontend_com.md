Authentication of farms against the farm endpoint
Request of JSON web token

The farm requires a valid farm identifier and farm secret to successfully receive a JSON web token. The identifier consists of an UUID and the secret is an alphanumeric string, special characters are permitted. To start the authentication, the farm has to send its identifier and secret to the "/authentication/gettoken/farm" endpoint. The endpoint requires the HTTP method POST and the content type "application/x-www-form-urlencoded". The required form fields are called identifier and secret. A possible request could look like "identifier=&secret=", while the and is replaced with its respective values. If the authentication was successful the endpoint responds with HTTP status code 200 and the content type "application/jwt", otherwise a HTTP status code 401 is send.
Authentication against the device endpoint

After a successful acquisition of a JSON web token, it must be send with each subsequent request. Therefore, the HTTP header "Authorization" must be added and the value must be of form "Bearer ", while is replaced with the actual token.

The default validity of a token is restricted to 60 seconds . After this period a new token must be acquired.
Authentication of users against the user endpoint
Request of JSON web token

The user requires a valid email and password to successfully receive a JSON web token. The email consists of a valid formatted email address and the password is an alphanumeric string, special characters are permitted. To start the authentication, the user has to send its email and password to the "/authentication/gettoken/user" endpoint. The endpoint requires the HTTP method POST and the content type "application/x-www-form-urlencoded". The required form fields are called email and password. A possible request could look like "email=&password=", while the and is replaced with its respective values. If the authentication was successful the endpoint responds with HTTP status code 200 and the content type "application/jwt", otherwise a HTTP status code 401 is send.
Authentication against the user endpoint

After a successful acquisition of a JSON web token, it must be send with each subsequent request. Therefore, the HTTP header "Authorization" must be added and the value must be of form "Bearer ", while is replaced with the actual token.

The default validity of a token is restricted to 30 seconds . After this period a new token must be acquired.