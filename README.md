# testMicroservice
Some dumb testing to learn..

A microservice with routing, to handle users (CR but not UD)

JWToken security first naïve build : 
/users page secured with specific handler, looking for a "security" cookie, containing an encrypted token
/login page check query string parameters as a credential couple. If couple is validated, response has a cookie with an encrypted token inside

Tried many time to use claims objects from dgrijalva/jwt-go, but yet no way to parse claims from a decrypted token ! 
