use serde::Serialize;
use crate::repository::schema;

#[derive(Debug, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    pub id : u32,
    pub first_name : String,
    pub last_name : String,
    pub address : String,
    pub city: String,
    pub telephone : String,
}

impl From<schema::Owner> for Response {
    fn from(owner: schema::Owner) -> Self {
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