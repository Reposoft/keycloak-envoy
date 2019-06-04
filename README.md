# Keycloak login flow with Envoy proxy as API Gateway

Scenario: An SPA is hosted together with REST and Websocket endpoints behind Envoy.
Endpoints can be protected with regular Bearer tokens, but we won't get a log in flow for users with Envoy's stock auth.

## Using Keycloak Gatekeeper

Our first approach is to use Gatekeeper as https://www.ory.sh/docs/oathkeeper/#access-control-decision-api though it wasn't designed that way

We should also allow access to a bearer token for use with Live
https://www.keycloak.org/docs/latest/securing_apps/index.html#endpoints

### TODO

How do we preserve the request path when users get a roundtrinp to Keycloak?
