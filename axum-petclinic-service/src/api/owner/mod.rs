use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    id : u64,
    first_name : String,
    last_name : String,
    address : String,
    city: String,
    telephone : String,
}