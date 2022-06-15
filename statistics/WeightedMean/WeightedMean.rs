use std::io;

fn weighted_mean(n: Vec<i32>, w: Vec<i32>) -> f32 {
    let l: usize = n.len();
    let mut sum_w: i32 = 0;
    let mut sum_n: i32 = 0;

    for i in 0..l {
        sum_w = sum_w + w[i];
        sum_n = sum_n + (w[i] * n[i])
    }

    return sum_n as f32 / sum_w as f32;
}

fn main() {
    let mut vec_size: String = String::new();
    io::stdin().read_line(&mut vec_size).unwrap();
    let mut _tc: i32 = vec_size.trim().parse().unwrap();

    let mut numbers: String = String::new();
    io::stdin().read_line(&mut numbers).unwrap();
    let mut numbers_arr: Vec<i32> = numbers
        .split(" ")
        .map(|n| n.trim().parse().unwrap())
        .collect();

    let mut w: String = String::new();
    io::stdin().read_line(&mut w).unwrap();
    let mut ws_arr: Vec<i32> = w.split(" ").map(|n| n.trim().parse().unwrap()).collect();

    println!("{:.1}", weighted_mean(numbers_arr, ws_arr));
}

