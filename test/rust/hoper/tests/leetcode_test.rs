use hoper::leetcode::*;

#[test]
fn two_sum_test(){
    let nums = vec![0,4,3,0];
    let target = 0;
    assert_eq!(two_sum(nums,target),[0,3])
}
