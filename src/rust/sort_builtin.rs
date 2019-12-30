use std::env;
use std::fs;
use std::fs::File;
use std::io::Write;

fn main() -> Result<(), Box<dyn std::error::Error + 'static>> {
    // setup
    let args: Vec<String> = env::args().collect();
    let input_file_path = &args[1];
    let output_file_path = &args[2];

    /////////////////////
    // read input file //
    /////////////////////

    let input_data_string: String = fs::read_to_string(input_file_path)?;
    let mut input_data_vec: Vec<&str> = input_data_string.lines().collect();

    ////////////////
    // sort input //
    ////////////////

    input_data_vec.sort();

    ///////////////////////
    // write output file //
    ///////////////////////

    let mut output_data_string: String = String::new();
    for input_data_item in input_data_vec {
        output_data_string.push_str(input_data_item);
        output_data_string.push_str("\n");
    }
    let output_data_bytes = output_data_string.as_bytes();
    let mut output_file = File::create(output_file_path)?;
    output_file.write_all(output_data_bytes)?;

    // teardown
    Ok(())
}