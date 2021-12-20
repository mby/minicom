#[macro_use] extern crate rocket;
use mongodb::bson;

mod types;
mod config;
mod routes;

pub struct App {
    client: mongodb::Client,
    brands: mongodb::Collection<types::Brand>,
}

#[launch]
async fn rocket() -> _ {
    let env = std::env::var("ENV").unwrap_or("qa".to_string());
    let cfg = config::get_config(env);
    let client = mongo_connect(cfg).await.expect("failed to connect mongodb");

    let brands = client.database(cfg.mongo_db).collection(cfg.brand_collection);
    let app = App{client, brands};

    rocket::build().manage(app).mount("/", routes![
        routes::health,
        routes::post_brand,
    ])
}

async fn mongo_connect(cfg: &'static config::Config) -> Result::<mongodb::Client, mongodb::error::Error> {
    let mut client_options = mongodb::options::ClientOptions::parse(cfg.mongo_uri).await?;
    client_options.app_name = Some("catalogue".to_string());

    let client = mongodb::Client::with_options(client_options)?;
    client.database(cfg.mongo_db).run_command(bson::doc! {"ping": 1}, None).await?;

    Ok(client)
}
