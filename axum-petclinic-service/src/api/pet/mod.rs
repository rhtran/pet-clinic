mod service;

use serde::{Deserialize, Serialize};
use crate::repository::schema;

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    id : u32,
    name: String,
    birthday: String,
    kind: String,
}

impl From<schema::Pet> for Response {
    fn from(pet: crate::repository::schema::Pet) -> Self {
        Self {
            id: pet.id,
            name: pet.name,
            birthday: pet.birthday,
            kind: pet.kind.name,
        }
    }
}

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Request {
    name: String,
    birthday: String,
    kind_id: u64,
    kind_name: String,
}
