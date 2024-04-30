use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Request {
    first_name: String,
    last_name: String,
    specialties: Vec<Specialty>,
}

#[derive(Debug, Deserialize, Serialize)]
#[serde(rename_all = "camelCase")]
struct Specialty {
    id: u64,
    name: String,
}

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    id: u64,
    first_name: String,
    last_name: String,
    specialties: Vec<Specialty>,
}