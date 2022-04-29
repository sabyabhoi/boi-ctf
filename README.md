<div id="top"></div>

<div align="center">

<h3 align="center">boi-ctf</h3>

  <p align="center">
	A Capture-The-Flag Server for improving your cyber-security skills
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->
### How does it work?
1. First a vulnerable program - like an unsafe C program making use of `gets` 
function call - is compiled. 
2. This program is then hosted on a given port of the server using [socat](https://linux.die.net/man/1/socat).
3. Users can interact with this and find their unique *flag*, which they can 
then post onto the server by sending a simple post request to the API endpoint
provided. They can view their rankings on it as well.

### Built With

* [Fiber](https://gofiber.io/)
* [GORM](https://gorm.io/)


### Prerequisites

* [GoLang](https://go.dev/)
* [PostgreSQL](https://www.postgresql.org/)
* [Socat](https://linux.die.net/man/1/socat)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/sabyabhoi/boi-ctf
   ```
2. Download the required GoLang libraries
   ```sh
   go mod tidy
   ```
3. Populate your `.env` file with the following variables
   ```env
	 DB_USER=your_postgres_username
	 DB_PASS=your_postgres_user_password
	 FLAG=boiCTF{the-actual-flag}
   ```
4. Launch the application using 
   ```sh
	 go build && ./boi-ctf

<!-- USAGE EXAMPLES -->
## Usage

The server currently has two major API endpoints on port `8080`:
1. `/` for getting general help about the CTF, and posting your flags.
2. `/leaderboard` for viewing the current rankings

The actual vulnerable program is hosted using `socat` on port `8081`. 
Connect to it using `netcat`:
```sh
nc <IP> 8081 
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<!-- CONTACT -->
## Contact

Sabyasachi Bhoi - [sabyabhoi](https://www.linkedin.com/in/sabyabhoi/) - sabyabhoi@gmail.com

