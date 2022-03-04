# Coffice New Normal Tracking App

<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <a href="">
    <img src="images/Logo-black.png" alt="Logo" height="200" width="200">
  </a>
  <p align="center">
    Pembangunan RESTful API menggunakan Gorilla Mux(router) & RS Cors(cors)
  </p>
</div>

### 🛠 &nbsp;Build App & Database

![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;

<!-- ABOUT THE PROJECT -->

## 💻 &nbsp;About The Project

<details>
<summary>ERD</summary>
<img src="images/capstone-project.jpg">
</details>
<details>
<summary>🧑‍💼🙎 &nbsp;Users</summary>
  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| ------------ | ---------| ----------- | ------------ | --------- | ------ |
| POST         | /users/login | - | email, nik, & password | NO |  |
| POST         | /users/register  | - | - | NO |  |
| POST         | /users/avatar  | - | - | YES |  |
| GET          | /users/ | - | - | YES |  |
| GET          | /users/profile | - | - | YES |  |
| PUT          | /users/ | - | - | YES |  | 
| DEL          | /users/ | - | - | YES |  |

</details>

<details>
<summary>🏤📦 &nbsp;Office</summary>
  
| Feature Office | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| --- | --- | --- | --- | --- | --- |
| GET | /offices/ | - | - | YES |  |

</details>

<details>
<summary>📓 &nbsp;Request</summary>

| Feature Request | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| --------------- | -------- | ----------- | ------------ | --------- | ------ |
| -               | -        | -           | -            | -         | -      |

</details>

<details>
<summary>🗓&nbsp;Day</summary>

| Feature Day | Endpoint | Query Param | Request Body | JWT Token       | Fungsi |
| ----------- | -------- | ----------- | ------------ | --------------- | ------ |
| GET         | -        | -           | -            | YES             | -      |
| POST        | -        | -           | -            | YES(only admin) | -      |

</details>

<details>
<summary>💉&nbsp;Vaccine Certificates</summary>

| Feature Certificates | Endpoint           | Query Param | Request Body | JWT Token | Fungsi |
| -------------------- | ------------------ | ----------- | ------------ | --------- | ------ |
| GET                  | /certificates/     | -           | -            | YES       | -      |
| GET                  | /certificates/user | -           | -            | YES       | -      |
| POST                 | /certificates/     | -           | -            | YES       | -      |
| PUT                  | /certificates/{id} | -           | -            | YES       | -      |

</details>

<details>
<summary>⏱️&nbsp;Checkins</summary>

| Feature Checkins | Endpoint    | Query Param | Request Body | JWT Token | Fungsi |
| ---------------- | ----------- | ----------- | ------------ | --------- | ------ |
| GET              | /check/     | -           | -            | YES       | -      |
| GET              | /check/user | -           | -            | YES       | -      |
| POST             | /check/ins  | -           | -            | YES       | -      |
| POST             | /check/outs | -           | -            | YES       | -      |

</details>

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

# How to Use

### 1. install

```bash
git clone https://github.com/HamzahAA15/group2-project-capstone
```

### 2.1 congif .env

```bash
touch .env
```

### 2.2 isi .env

```bash
isi .env nya
```

### 3. run main.go

```bash
go run main.go
```

# Contact

[![Linkedin](https://img.shields.io/badge/-Jundana-white?style=flat&logo=linkedin&logoColor=blue)](https://www.linkedin.com/in/jundanaalbasyir/)
[![Linkedin](https://img.shields.io/badge/-Hamzah-white?style=flat&logo=linkedin&logoColor=blue)](https://www.linkedin.com/in/hamzahaalfauzi/)
[![GitHub](https://img.shields.io/badge/-Jundana-white?style=flat&logo=github&logoColor=black)](https://github.com/justjundana)
[![GitHub](https://img.shields.io/badge/-Hamzah-white?style=flat&logo=github&logoColor=black)](https://github.com/HamzahAA15)

<p align="center">:copyright JH Top Golang Dev</p>
</h3>

<p align="right">(<a href="#top">back to top</a>)</p>
