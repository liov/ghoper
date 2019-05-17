use hoper::leetcode::*;

#[test]
fn two_sum_test(){
    let nums = vec![0,4,3,0];
    let target = 0;
    assert_eq!(two_sum2(nums,target),[0,3])
}

#[test]
fn add_two_numbers_test(){
    let mut l1= Box::new(ListNode::new(2));
    l1.push(4);
    l1.push(3);
    let mut l2=Box::new(ListNode::new(5));
    l2.push(6);
    l2.push(4);
    let mut result=Box::new(ListNode::new(7));
    result.push(0);
    result.push(8);
    assert_eq!(add_two_numbers(Some(l1),Some(l2)),Some(result))
}

#[test]
fn multiply_test(){
    let s1 = String::from("136");
    let s2 = String::from("5261939");
    assert_eq!(multiply(s1,s2),  String::from("715623704"))
}

#[test]
fn multiply_test2(){
    let s1 = String::from("5");
    let s2 =s1.as_bytes();
    let b1:[u8;1]=[53];
    assert_eq!(*s2,b1 )
}

