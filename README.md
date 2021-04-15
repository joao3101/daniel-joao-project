<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

The goal of this project is to create a fantasy app for the Brazilian Soccer championship, but based on a draft, such as the NFL fantasy.


### Built With

* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)
* [PostgreSQL](https://www.postgresql.org/)
* [go-migrate](https://github.com/golang-migrate/migrate)



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* Go stable version
* Docker

### Installation

1. Clone the repo
  ```sh
   git clone git@github.com:joao3101/daniel-joao-project.git
   ```
2. Install Docker PostgreSQL image (cmd available from Makefile at root)
   ```sh
   make docker_postgres
   ```
3. Start Postgres image
   ```sh
   make docker_compose
   ```
4. Create the DB
   ```sh
   make createdb
   ```
5. Run SQL Boiler to generate the code
   ```sh
   make sqlboiler
   ```
6. Start the server
   ```sh
   make server
   ```



<!-- USAGE EXAMPLES -->
## Usage


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Daniel Rocha Lopes - [LinkedIn](https://www.linkedin.com/in/daniel-rocha-lopes-42a55b149/)

João Victor de Carvalho Silva - [LinkedIn](https://www.linkedin.com/in/joao-victor-de-carvalho-silva/)

Project Link: [https://github.com/joao3101/daniel-joao-project](https://github.com/joao3101/daniel-joao-project)