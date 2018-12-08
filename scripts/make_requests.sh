#!/bin/bash

SERVICE_PORT=${SERVICE_PORT:-8080}
TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidWlkIjoidXNlci1pZCIsImNoaXBzIjoxMjM0NSwiYmV0Ijo1MDAsImlhdCI6MTU0NDI0ODk3NywiZXhwIjoyODA2NTUyOTc3fQ.r5o41ZBolMHc4aMOrkH_1x_w6zq1FMX0jW1_vyEEsqw'
curl -X POST http://localhost:${SERVICE_PORT}/api/machines/atkins-diet/spins -H "Authorization: ${TOKEN}" \
