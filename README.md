# sqdif
This CLI tool, named `sqdif`, is designed to parse a specified SQL file and generate dummy data based on its content. Users have the ability to specify the path to the output file, whether to overwrite an existing output file, and the prompt to be passed to the GPT-4 API.


## Prerequisites
To run this program, you will need the following prerequisites:
- Go 1.18 or higher
- OpenAI API key (set in the environment variable OPENAI_API_KEY)
- Dependencies specified in the go.mod file
## Usage
The program can be executed from the command line with the following options:

- `--sql-file`: This flag specifies the file to be passed to GPT-4.
- `--output-file`: Path to the output file (optional)
- `--override`: Flag to indicate whether to overwrite the existing output file (optional)
- `--prompt`: Flag to specify the prompt to be passed to GPT-4 (optional)
- `--api-key`: The API key for the GPT-4 API

## Example
```sh
sqdif  --sql-file=./testdata/table.sql --output-file=./testdata/test.sql --api-key={api_key} --override=true
```
## License
This project is released under the MIT license.
