# chat microservices

My contact: [telegram](https://t.me/cs_and_dev)

Homework form course ["Microservices, as in BigTech"](https://olezhek28.courses/)

## Generating `.env` file for `docker-compose`
```bash
cat <<EOF > .env
POSTGRES_DB=<your_database_name>
POSTGRES_USER=<your_username>
POSTGRES_PASSWORD=<your_password>
EOF
```