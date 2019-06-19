use actix_web::{App, web};

use super::handler;
use actix_web::middleware::Logger;
use actix_session::CookieSession;
use actix_web::dev::Body;

pub(crate) fn new() -> Result<App<(), Body>, ()>{
    App::new()
        .data(state)
        .wrap(Logger::default())
        .wrap(CookieSession::signed(&[0; 32]).secure(false))
        .wrap(IdentityService::new(
            CookieIdentityPolicy::new(secret.as_bytes())
                .name("auth")
                .path("/")
                .domain(domain.as_str())
                .max_age_time(Duration::days(1))
                .secure(false), // this can only be true if you have https
        )).service(
        web::scope("/test")
            .service(web::resource("/index1").route(web::get().to(handler::index1)))
            .service(web::resource("/index2").route(web::get().to(handler::index2)))
            .service(web::resource("/index3").route(web::get().to(handler::index3)))
            .service(web::resource("/index5").route(web::get().to(handler::index5)))
    )
    // serve static files
    //.service(fs::Files::new("/", "./static/").index_file("index.html"))
}
