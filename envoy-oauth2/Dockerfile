FROM curlimages/curl:7.74.0@sha256:a3e534fced74aeea171c4b59082f265d66914d09a71062739e5c871ed108a46e as pr14168

# https://github.com/envoyproxy/envoy/pull/14168
# https://github.com/veehaitch/envoy-oauth2-filter-google/blob/main/envoy/Dockerfile
RUN curl -o /tmp/envoy https://files.veehait.ch/envoy-pr14168-oauth_scope/envoy
RUN echo "2656ce3fb890588f3a0373c8009a01e1311a9d39234cee7102ebe985ceeea9d0  /tmp/envoy" | sha256sum -c -
RUN chmod +x /tmp/envoy

FROM envoyproxy/envoy:v1.16.2@sha256:8a1d010267ee291bb0e7f12bb24e43b255c771629d70482910bc1f13a56e6b2f

COPY --from=pr14168 /tmp/envoy /usr/local/bin/envoy
