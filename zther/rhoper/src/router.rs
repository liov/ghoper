use actix_web::App;
use actix_web::error;
use actix_web::http::Method;
use actix_web::http::header;
use actix_web::middleware::cors::Cors;
use actix_web::middleware;
use actix_web::fs;
use actix_web::HttpRequest;
use actix_web::HttpResponse;
use actix_web::Body;



pub fn app() -> Result<App<()>, error::Error> {

    let app = App::new()
        .resource("/", |r| r.f(index));
    Ok(app)
}

fn index(_req: &HttpRequest) -> &'static str {
    "Hello world!"
}