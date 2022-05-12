use std::io;

fn cumulative(day: i32) -> i32 {
    let mut counter: i32 = 5;
    let mut result: i32 = 0;
    for _ in 1..(day+1) {
        counter = counter/2;
        result += counter;
        counter = counter*3;
    }
    result
}

fn main() {
    let mut input: String = String::new();
    io::stdin().read_line(&mut input).expect("failed to read");
    let r: i32 = cumulative(input.trim().parse().expect("Expecting a number"));

    println!("{}", r);
}
