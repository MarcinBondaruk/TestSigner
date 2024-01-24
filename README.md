### Assumptions
Questions and answeres are expected to be strings. This can cover all cases like:
* question is an id, or text, url to image or something else
* answer could be an id, plaintext for open questins or list of answers for multi choice questions
* different set of questions are considered as a different test

Signature of JWT token is always considered valid due to lack of auth service
userId **is expected** to be UUID
example JWT `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwMThkM2M1YS02ZDJiLTc5ZWQtOTRhYi05MGU5ZDliYTUyNmIifQ.dvL7cL_KA-9eK9nTFArryTF-o1nxKoOWtF-xhzsfIyE`
