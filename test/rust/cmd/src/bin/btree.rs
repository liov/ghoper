use std::collections::BTreeSet;

fn main() {
    let mut set: BTreeSet<usize> = BTreeSet::new();
    set.insert(3);
    set.insert(5);
    set.insert(6);
    set.insert(2);
    set.insert(7);
    set.insert(0);
    let mut set_iter = set.iter();

    while let Some(ref mut t) = set_iter.next() {
        *t =&5usize;
        println!("{}", t);
        set_iter = set.iter();
    }
}
