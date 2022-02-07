if [[ $1 == "available" ]]; then
    go run produce_messages.go \
        -topic platform.sources.status \
        -value '{"resource_type":"application","resource_id":"1", "status":"available","error":""}'
    else

    go run produce_messages.go \
        -topic platform.sources.status \
        -value '{"resource_type":"application","resource_id":"1", "status":"unavailable","error":"yeet"}'
fi
