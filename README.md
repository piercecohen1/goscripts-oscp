# goscripts-oscp
#### A collections of tiny but useful Go programs. Helpful for Penetration Testing and the OSCP Exam.

#### I reccommend creating aliases for each of them.

## How To Set Up

### Ensure you have Golang installed

To check if Golang is installed, run the following:
```
go version
```

If you get an error, you must install Go:

On Linux:
```
sudo apt install golang -y
```

On macOS:
```
brew install go
```

### Build

#### Option 1: Automatic Install (recommended)
Use the following to automatically build and create aliases:
```
./full_install.sh
```

#### Option 2: Compile only (no aliases)
```
./build_all.sh
```

(Optional) Manually create aliases for each program

Modify your `.bash_rc` or `.zsh_rc` to add the following lines, replacing the path with the actual path for each:
```
alias urlencode='/path/to/goscripts-oscp/url_encode/url_encode'
alias escape_double_quotes='/path/to/goscripts-oscp/escape_quotes/escape_double_quotes'
alias psencode='/path/to/goscripts-oscp/base64_powershell/b64'
alias path='/path/to/goscripts-oscp/path/copy_path'
```
Note that if you you relocate the binaries, you will need to update the path for each alias.
