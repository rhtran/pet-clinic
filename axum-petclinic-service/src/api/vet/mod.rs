use serde::{Deserialize, Serialize};
use crate::repository::schema;

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
    name: String,
}

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    id: u32,
    first_name: String,
    last_name: String,
    specialties: Vec<Specialty>,
}

impl From<schema::Vet> for Response {
    fn from(vet: crate::repository::schema::Vet) -> Self {
        Self {
            id: vet.id,
            first_name: vet.first_name,
            last_name: vet.last_name,
            specialties: vet.specialties.into_iter().map(|s| Specialty {
                name: s.name,
            }).collect(),
        }
    }
}