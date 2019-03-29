#[derive(Debug)]
pub struct User {
    pub  id: u64,
    pub  sex: Sex,
    pub name: String,
    pub password: String,
    pub email: String,
    pub  phone: String,
}

impl User {
    pub  fn get_name(&self) -> String {
        self.name.clone()
    }

    pub fn set_name(&mut self, name: String) {
        self.name = name
    }
}

#[derive(Debug)]
pub enum Sex{
    Boy,
    Girl,
}