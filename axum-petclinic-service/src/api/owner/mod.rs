use axum::{Json, Router};
use axum::extract::Path;
use axum::routing::get;

mod model;
mod service;

pub fn routes() -> Router {
    // let app_state = AppState { mc };
    Router::new().route("/owner/id/:id", get(get_owner))
}

async fn get_owner( Path(id): Path<u32>) -> Json<model::Response> {
    let owner = model::Response {
        id: id,
        first_name: "John".to_string(),
        last_name: "Doe".to_string(),
        address: "123 Main St".to_string(),
        city: "".to_string(),
        telephone: "123-456-7890".to_string(),
    };

    Json(owner)
}