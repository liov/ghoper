#![feature(nll)]
use std::cmp::PartialOrd;
use std::fmt::Debug;

//本人二叉树
#[derive(Debug)]
pub struct MyTree<T: PartialOrd + Debug> {
    pub data:Option<Box<T>>,
    pub left:Option<Box<MyTree<T>>>,
    pub right:Option<Box<MyTree<T>>>,
}

impl<T: PartialOrd + Debug> MyTree<T> {
    pub fn insert(&mut self, data: T) {
        match self.data {
            Some(ref mut rdata) =>
               if data < **rdata {
                   match self.left {
                       Some(ref mut left) =>
                       left.insert(data),
                       None =>
                           {
                               self.left = Some(Box::new(MyTree::new()));
                               if let Some(ref mut left) = self.left {
                                   left.insert(data)
                               }
                           },
                   }


               }else {
                   match self.right {
                       Some(ref mut right) =>
                           right.insert(data),
                       None =>
                           {
                               self.right = Some(Box::new(MyTree::new()));
                               if let Some(ref mut right) = self.right {
                                   right.insert(data)
                               }
                           },
                   }
               }
            None => self.data = Some(Box::new(data)),
        }
    }

    pub fn new() -> MyTree<T> {
        MyTree {
            data: None,
            left: None,
            right:None,
        }
    }
    pub fn peek(&self) {
        match self.data {
            Some(ref data) =>{
                println!("{:?}",**data);
                match self.left {
                    Some(ref left) =>
                    left.peek(),
                    None =>{},
                }
                match self.right {
                    Some(ref right) =>
                        right.peek(),
                    None =>{},
                }
            }
            None => {},
        }
    }
}


struct Node<T: PartialOrd + Debug> {
    elem: T,
    left: Tree<T>,
    right: Tree<T>,
}

struct Tree<T: PartialOrd + Debug> {
    data: Option<Box<Node<T>>>,
}

impl<T: PartialOrd + Debug> Tree<T> {
    pub fn insert(&mut self, data: T) {
        match self.data {
            Some(ref mut node) => if node.elem > data {
                node.left.insert(data)
            } else {
                node.right.insert(data)
            },
            None => self.data = Some(Box::new(Node::new(data))),
        }
    }

    fn new() -> Tree<T> {
        Tree { data: None }
    }
}

impl<T: PartialOrd + Debug> Node<T> {
    pub fn new(t: T) -> Self {
        Node {
            elem: t,
            left: Tree::new(),
            right: Tree::new(),
        }
    }
}
