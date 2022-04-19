# betnomi

THE PLATFORM


The platform is the minimal version of Betnomi. Betnomi consists of some
microservices written in golang and communicates using grpc over nats. We want you to
develop 2 services reachable over envoy and communicate with each other over nats. The
db should be postgres.

Overall high level architecture should look like this;

Envoy should use prot descriptors to communicate GRPC service and provide http
endpoint
(https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transc
oder_filter)
Services must communicate with each other using nats topics, not direct calls to
each other at all.
Each service should have their own DB (for example, if you need users in
transactions, you can keep list of ids etc whatever you need for easy ownership relations)
And the service should have 4 endpoints
/user/login (should return a token, it does not have to do authentication)
/user/balance (should return user balance)
/transactions/up (should add 1 transaction record and update user balance, amount should
come in body)
/transactions/down (should add 1 transaction record that decreases user balance, amount
should come in body)

