use sqlx::{Pool, Postgres};
use sqlx::postgres::PgPoolOptions;
use crate::error::Error;

pub async fn create_connection() -> Result<Pool<Postgres>, Error> {
    let conn_string = "postgresql://postgres:mysecretpassword@localhost:5432/petclinic?sslmode=disable";
    PgPoolOptions::new()
        .min_connections(1)
        .max_connections(5)
        .connect(conn_string)
        .await
        .map_err(|e| {
            Error::new("create_connection", &format!("{:?}", e))
        })
}