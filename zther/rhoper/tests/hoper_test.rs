use rhoper::config;


#[hoper]
fn ddd(){
    let a:[String;5] = ["a".to_string(),"b".to_string(),"c".to_string(),"d".to_string(),"e".to_string()];
    config::parse_config(&a);
}

use hoper::math::add;

#[hoper]
fn add_two_a(){
    assert_eq!(4,add(1,3))
}
