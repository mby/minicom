use lazy_static::lazy_static;
use std::collections::HashMap;

pub struct Config {
        pub mongo_uri: &'static str,
        pub mongo_db: &'static str,
        pub brand_collection: &'static str,
}

lazy_static! {
        static ref CONFIGS: HashMap<&'static str, Config> = {
                let mut map = HashMap::new();
            
                map.insert("qa", Config {
                    mongo_uri: "mongodb://localhost:27017", // TODO: mongodb
                    mongo_db: "catalogue-qa",
                    brand_collection: "brands",
                });
            
                map.insert("prod", Config {
                    mongo_uri: "mongodb://localhost:27017", // TODO: mongodb
                    mongo_db: "catalogue",
                    brand_collection: "brands",
                });
            
                map
        };
}

pub fn get_config(env: String) -> &'static Config {
    CONFIGS.get(env.as_str()).unwrap()
}
