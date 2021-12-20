use rocket::response::Responder;

#[derive(Responder)]
pub struct BrandPostRequest {
        pub name: String,
}

#[derive(Responder)]
pub struct BrandPostResponse {
        pub status: String,
}
