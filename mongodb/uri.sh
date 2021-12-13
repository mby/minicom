kubectl get secret example-mongodb-agent-password -n mongodb-operator -o json | jq -r '.data | with_entries(.value |= @base64d)'
