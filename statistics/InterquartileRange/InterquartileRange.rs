use std::io;

fn get_final_data(numbers: Vec<f32>, freq: Vec<f32>) -> Vec<f32> {
    let mut data: Vec<f32> = Vec::new();
    let mut idx = 0;
    for i in numbers {
        for _j in 0..(freq[idx] as i32) {
            data.push(i);
        }
        idx = idx + 1;
    }
    return data;
}

fn quartiles(mut data: Vec<f32>) -> (f32, f32, f32) {
    data.sort_by(|a, b| a.partial_cmp(b).unwrap());
    let data_len = data.len();

    if data_len % 2 == 0 {
        return (
            median(&data[0..data_len / 2].to_vec()),
            median(&data),
            median(&data[(data_len / 2)..].to_vec()),
        );
    } else {
        return (
            median(&data[0..(data_len / 2) - 1].to_vec()),
            median(&data),
            median(&data[(data_len / 2) + 1..].to_vec()),
        );
    }
}

fn median(data: &Vec<f32>) -> f32 {
    let elems_no = data.len();
    if elems_no % 2 == 0 {
        let middle_n = (elems_no / 2) - 1;
        return ((data[middle_n] as f32) + (data[middle_n + 1] as f32)) / 2 as f32;
    } else {
        return data[elems_no / 2];
    }
}

fn main() {
    let mut vec_size: String = String::new();
    io::stdin().read_line(&mut vec_size).unwrap();
    let mut _tc: i32 = vec_size.trim().parse().unwrap();

    let mut numbers: String = String::new();
    io::stdin().read_line(&mut numbers).unwrap();
    let numbers_arr: Vec<f32> = numbers
        .split(" ")
        .map(|n| n.trim().parse().unwrap())
        .collect();

    let mut w: String = String::new();
    io::stdin().read_line(&mut w).unwrap();
    let ws_arr: Vec<f32> = w.split(" ").map(|n| n.trim().parse().unwrap()).collect();

    let r = get_final_data(numbers_arr, ws_arr);
    let (q1, q2, q3) = quartiles(r);

    println!("{:.1}", (q3-q1) as f32);
}