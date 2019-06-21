fn main(){
    fn a(a:i32)->i32{
        println!("a:{:p}",&a);
        a
    }
    let b = a ;
    let b = a(1);
    println!("b:{:p}",&b);
}
