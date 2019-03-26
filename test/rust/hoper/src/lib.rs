pub mod utils;

#[cfg(test)]
mod tests {
    use crate::math::add;

    #[test]
    fn add_two_a(){
        assert_eq!(4,add(1,3))
    }
}

pub mod math{
    pub fn add(a:i32,b:i32) -> i32{a+b}
}
