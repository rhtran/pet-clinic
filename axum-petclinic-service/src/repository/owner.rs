use axum::http::StatusCode;
use crate::db;
use crate::error::Error;
use crate::repository::schema::Owner;

// pub trait Repositorier {
//     fn new(&self) -> Self;
//     fn list_owners(&self) -> Result<Vec<Owner>, Error>;
//     // fn find_owner_by_id(&self, id: i32) -> Result<Owner, Error>;
//     // fn create_owner(&self, owner: Owner) -> Result<Owner, Error>;
//     // fn update_owner(&self, id: i32, owner: Owner) -> Result<Owner, Error>;
//
// }

#[derive(Debug, Clone)]
pub struct Repository {
}

impl Repository {
    pub async fn new(&self) -> Self {
        Self {}
    }
    pub async fn list_owners(&self) -> Result<Vec<Owner>, Error> {
        let db = db::create_connection();
        let owners = sqlx::query("SELECT id, first_name FROM owners")
            .fetch_all(db)
            .map(|result| {
                result
                    .into_iter()
                    .map(|row| Owner {
                        id: row.get(0),
                        first_name: row.get(1),
                        last_name: row.get(2),
                        address: row.get(3),
                        city: row.get(4),
                        telephone: row.get(5),
                    })
                    .collect()
            });
        Ok(owners)
    }

    // fn find_owner_by_id(&self, id: i32) -> Result<Owner, Error> {
    //     let owner = sqlx::query("SELECT * FROM owners WHERE id = $1")
    //         .bind(id)
    //         .fetch_one(&self.db)
    //         .map_err(|e| {
    //             Error::new("find_owner_by_id", &format!("{:?}", e))
    //         })?;
    //     Ok(owner)
    // }
    //
    // fn create_owner(&self, owner: Owner) -> Result<Owner, Error> {
    //     let owner = sqlx::query("INSERT INTO owners (first_name, last_name, address, city, telephone) VALUES ($1, $2, $3, $4, $5) RETURNING *")
    //         .bind(&owner.first_name)
    //         .bind(&owner.last_name)
    //         .bind(&owner.address)
    //         .bind(&owner.city)
    //         .bind(&owner.telephone)
    //         .fetch_one(&self.db)
    //         .map_err(|e| {
    //             Error::new("create_owner", &format!("{:?}", e))
    //         })?;
    //     Ok(owner)
    // }

    // fn update_owner(&self, id: i32, owner: Owner) -> Result<Owner, Error> {
    //     let owner = sqlx::query("UPDATE owners SET first_name = $1, last_name = $2, address = $3, city = $4, telephone = $5 WHERE id = $6 RETURNING *")
    //         .bind(&owner.first_name)
    //         .bind(&owner.last_name)
    //         .bind(&owner.address)
    //         .bind(&owner.city)
    //         .bind(&owner.telephone)
    //         .bind(id)
    //         .fetch_one(&self.db)
    //         .map_err(|e| {
    //             Error::new("update_owner", &format!("{:?}", e))
    //         })?;
    //     Ok(owner)
    // }
}