curl "http://localhost:8080/stream/ListMessages" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "channel": "general"
}'