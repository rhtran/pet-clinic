use serde::{Deserialize, Serialize};
use crate::api::pet;

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    id : u64,
    visit_date: String,
    description: String,
    pet: pet::Response,
}