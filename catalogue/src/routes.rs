use super::{App, types};
use rocket::{State, form::Form};
use rocket::response::content::Json;

#[get("/health")]
pub fn health() -> &'static str {
        "{\"status\": \"ok\"}"
}

#[post("/brand", data = "<brand>")]
pub async fn post_brand(app: &State<App>, brand: Form<types::BrandPostRequest>) -> Json<types::BrandPostResponse> {
        app.brands.insert_one(types::Brand{
                name: brand.name.clone(),
        }, None).await.unwrap();

        Json(types::BrandPostResponse{
                status: "ok".to_string(),
        })
}
