HOST_IP=$(hostname -I | cut -d \  -f 1)

# DB
export GT_DB_HOST=${HOST_IP}
export GT_DB_PORT="21001"
export GT_DB_USER="user"
export GT_DB_PASS="pass"
export GT_DB_NAME="db"
export GT_DB_DIALECT="postgres"

# Prometheus
export GT_PROM_PGW_PORT=21002

env | grep GT_ > .dev.env
