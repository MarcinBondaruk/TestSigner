### Overview
I implemented loose layerd architecture. Directory structure is flat - i dont really like this but it required much less thinking and i focused on delivering. I like explicit Layers Application/Domain/Infra

Please remember to start up DB `docker-compose up -d test_signer_db`

### Assumptions
Questions and answeres are expected to be strings. This can cover all cases like:
* question is an id, or text, url to image or something else
* answer could be an id, plaintext for open questins or list of answers for multi choice questions
* different set of questions are considered as a different test

Signature of JWT token is always considered valid, if it has userId in its claims, due to lack of auth service
userId **is expected** to be UUID
example JWT `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwMThkM2M1YS02ZDJiLTc5ZWQtOTRhYi05MGU5ZDliYTUyNmIifQ.dvL7cL_KA-9eK9nTFArryTF-o1nxKoOWtF-xhzsfIyE`


### API
JWT should be passed in Authorization header as bearer token
`localhost:8080/api/v1/tests/sign` - sign a JWT expects json body [{"question": "myquestion", "answer": "users answer"}, ...]
`localhost:8080/api/v1/tests?user_id=someUUID&signature=someSignature` - retrieve by userId and signature. Endpoint returns error when userId or signature is missing. This endpoint **DOES NOT EXPECT** JWT token. thats what i figured from the task description

### Some known issues
**Forgot to make db calls concurrent**
Error handling might not be perfect, for example if signature already exists (meaning this test by this user was aleady signed) i dont really remap infrastructural error to application layer one.

There might be some cleanup needed



