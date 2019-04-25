use hoper::utils::tree::MyTree;
//只有库才能有test
#[test]
fn add_tree() {
        let x = std :: f64 :: consts :: PI;
        let mut t = MyTree::new();
        t.insert(x);
        println!("{:?}",t);
        t.insert(3f64);
        println!("{:?}",t);
        t.insert(5f64);
        println!("{:?}",t);
        t.insert(9f64);
        println!("{:?}",t);
        t.insert(66f64);
        println!("{:?}",t);
        t.insert(18f64);
        println!("{:?}",t);
        t.insert(12f64);
        println!("{:?}",t);
        t.insert(111f64);
        println!("{:?}",t);
        t.insert(12f64);
        println!("{:?}",t);
        t.insert(2f64);
        println!("{:?}",t);
        t.peek();
}
