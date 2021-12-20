use rocket::{form::FromForm, response::Responder};

#[derive(FromForm)]
pub struct BrandPostRequest {
        pub name: String,
}

#[derive(Responder)]
pub struct BrandPostResponse {
        pub status: String,
}
