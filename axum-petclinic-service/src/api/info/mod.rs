use axum::{Json, Router};
use axum::routing::get;
use derive_builder::Builder;
use serde::{Serialize};
use validator::Validate;
use crate::error::Error;

#[derive(Builder, Debug, Serialize, Validate)]
#[serde(rename_all = "camelCase")]
struct InfoResponse {
    #[validate(length(min = 1, message = "Version must be at least 1 character long"))]
    version: String,
    #[validate(length(min = 1, message = "Name must be at least 1 character long"))]
    name: String,
    description: String,
    author: String,
    license: String,
}

pub fn routes() -> Router {
    // let app_state = AppState { mc };
    Router::new().route("/info", get(info))
}

async fn info() -> Json<InfoResponse> {
    let info = InfoResponseBuilder::default()
        .version("0.1.0".to_string())
        .name("Rust API".to_string())
        .description("Rust API using Axum".to_string())
        .author("Ethan J. Tran".to_string())
        .license("MIT".to_string())
        .build()
        .unwrap();

    Json(info)
}