use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    id : u64,
    name: String,
    birthday: String,
    kind: String,
}

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Request {
    name: String,
    birthday: String,
    kind_id: u64,
    kind_name: String,
}
