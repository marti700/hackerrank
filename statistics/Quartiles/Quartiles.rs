use std::io;

fn quartiles(mut data: Vec<i32>) -> (i32, i32, i32) {
    data.sort();
    let data_len = data.len();

    if data_len % 2 == 0 {
      return (median(&data[1..data_len/2].to_vec()), median(&data), median(&data[(data_len/2)..].to_vec()))
    } else {
      return (median(&data[1..(data_len/2)-1].to_vec()), median(&data), median(&data[(data_len/2)+1..].to_vec()))
    }
}

fn median(data: &Vec<i32>) -> i32 {
  let elems_no = data.len();
    if elems_no % 2 == 0 {
        let middle_n = (elems_no / 2) - 1;
        return ((data[middle_n]) + (data[middle_n + 1])) / 2
    } else {
        return data[elems_no / 2]
    }
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

    let (q1, q2,q3) = quartiles(numbers_arr);
    println!("{}", q1);
    println!("{}", q2);
    println!("{}", q3);
}
