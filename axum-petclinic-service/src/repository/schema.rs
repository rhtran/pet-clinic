use chrono::{ NaiveDateTime };
use sqlx::FromRow;

#[derive(Debug, Clone, FromRow)]
pub struct Owner {
    pub id: u32,
    pub first_name: String,
    pub last_name: String,
    pub address: String,
    pub city: String,
    pub telephone: String,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Pet {
    pub id: u32,
    pub name: String,
    pub birthday: String,
    pub owner: Owner,
    pub kind: Kind,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Vet {
    pub id: u32,
    pub first_name: String,
    pub last_name: String,
    pub specialties: Vec<Specialty>,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Visit {
    pub id: u32,
    pub visit_date: String,
    pub description: String,
    pub pet: Pet,
    pub owner: Owner,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Specialty {
    pub id: u32,
    pub name: String,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Kind {
    pub id: u32,
    pub name: String,
    pub base: Base,
}

#[derive(Debug, Clone, FromRow)]
pub struct Base {
    pub created_at: NaiveDateTime,
    pub update_at : NaiveDateTime,
    pub deleted_at: NaiveDateTime,
}
