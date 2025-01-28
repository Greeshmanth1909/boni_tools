#!/bin/bash 
apiKey="bd7ad5c1681b74b335f2960bce87805f"
secret="bfd826b5e8"
curl -i \
-X GET \
-H 'Accept:application/json' \
-H 'Api-key:'$apiKey'' \
-H 'X-Signature:'$(echo -n ${apiKey}${secret}$(date +%s)|sha256sum|awk '{ print $1}')'' \
https://api.test.hotelbeds.com/hotel-api/1.0/status
