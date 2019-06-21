
use chrono::{Local, NaiveDateTime};

use uuid::{Uuid};

pub struct User {
    pub  id: Uuid,
    pub  sex: Sex,
    pub name: String,
    pub email: String,
    pub  phone: String,
    pub password: String,
    pub created_at: NaiveDateTime,
}

impl User {
    pub  fn get_name(&self) -> String {
        self.name.clone()
    }

    pub fn set_name(&mut self, name: String) {
        self.name = name
    }

    pub fn with_details(email: String, password: String) -> Self {
        User {
            id: Uuid::new_v5(&&Uuid::NAMESPACE_OID,"id".as_bytes()),
            sex: Sex::Boy(0),
            name: "".to_string(),
            email,
            phone: "".to_string(),
            password,
            created_at: Local::now().naive_local(),
        }
    }
}

#[derive(Debug, Serialize)]
pub enum Sex{
    Boy(i32),
    Girl(i32),
}



