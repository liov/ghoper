///两数之和
///
use std::collections::HashMap;
//暴力法52ms，2MB
//一遍hash版本，0ms，2.7MB
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map: HashMap<i32,usize> = HashMap::new();
    let mut index = 0;
    while index < nums.len() {
        if let Some(j) = map.get(&(target-nums[index])){
            return vec![*j as i32,index as i32]
        }
        map.insert(nums[index],index);
        index =index+1
    }
    panic!("不存在")
}
