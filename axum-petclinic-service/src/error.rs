use axum::{
    http::StatusCode,
};

#[derive(Clone, Debug)]
pub enum Error {
    Sqlx(StatusCode, String),
    NotFound,
    InternalServerError(String),
}
//
// impl From<sqlx::Error> for Error {
//     fn from(err: sqlx::Error) -> Error {
//         match err {
//             sqlx::Error::RowNotFound => Error::NotFound,
//             _ => Error::Sqlx(
//                 StatusCode::INTERNAL_SERVER_ERROR,
//                 err.to_string(),
//             ),
//         }
//     }
// }
//
// impl IntoResponse for Error {
//     fn into_response(self) -> Response {
//         match self {
//             Error::Sqlx(code, body) => (code, body).into_response(),
//             Error::NotFound => StatusCode::NOT_FOUND.into_response(),
//         }
//     }
// }