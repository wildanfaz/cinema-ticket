# Cinema Ticket

Cinema Ticket is an app used for booking movie ticket easily.

## Table of Contents

- [About Me](#about-me)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Documentations](#documentations)

## About Me

Hello! I'm Muhamad Wildan Faz, a Junior Backend Developer who's deeply passionate about software development.

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' if not already installed. 

    * On Debian/Ubuntu, you can use

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/)

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/)

    ```bash
    choco install make
    ```

3. Clone the repository

    ```bash
    git clone https://github.com/wildanfaz/cinema-ticket.git
    ```

4. Change to the project directory

    ```bash
    cd cinema-ticket
    ```

## Usage

1. Start the application using docker

    ```bash
    docker-compose up
    ```

2. Open this [postman documentation](https://documenter.getpostman.com/view/22978251/2sA35Ba43C) to test the endpoints

3. Check commands if there's a need to set the role to admin

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```

3. Set role to admin
    ```bash
    make admin $(email) $(database-url)
    ```

    Example
    ```bash
    make admin email=example@mail.com database-url=postgres://postgres:mysecretpassword@localhost:5432/cinema-ticket
    ```

4. Add user balance
    ```bash
    make balance $(email) $(amount) $(database-url)
    ```

     Example
    ```bash
    make balance email=example@mail.com amount=1000000 database-url=postgres://postgres:mysecretpassword@localhost:5432/cinema-ticket
    ```

## Documentations

1. [Postman](https://documenter.getpostman.com/view/22978251/2sA35Ba43C)

2. [Database](https://dbdiagram.io/d/Cinema-Ticket-65fd5676ae072629ceb32d68)

3. [Flowchart](https://drive.google.com/file/d/1TzVBYXffa7qyMmKkGvhGquNRMudZRAJ6/view?usp=sharing)