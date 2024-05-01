mod route;

use serde::{Serialize};
use crate::repository::schema;

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    id : u32,
    first_name : String,
    last_name : String,
    address : String,
    city: String,
    telephone : String,
}

impl From<schema::Owner> for Response {
    fn from(owner: crate::repository::schema::Owner) -> Self {
        Self {
            id: owner.id,
            first_name: owner.first_name,
            last_name: owner.last_name,
            address: owner.address,
            city: owner.city,
            telephone: owner.telephone,
        }
    }
}