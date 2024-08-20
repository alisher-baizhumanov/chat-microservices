# chat microservices

My contact: [telegram](https://t.me/cs_and_dev)

## Migrations
1) pre requirements:
    ```bash
    make install
    ```
2) create migration:
    ```bash
    make migration-create MIGRATION_NAME=<migration_name>
    ```
3) setting `.env` file:
    ```bash
    cp .env.example .env
    ```
   Open `.env` file and set environments.