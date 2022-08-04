# Clothing Online Store Group 2

<!-- ABOUT THE PROJECT -->

## 👗 About The Project 👕

An online retailer called Clothing Online Store carries a selection of clothing from different nations. To serve the demands of Indonesians looking to purchase foreign clothing online, this online retailer was established.

</details>   
       
### 🛠 &nbsp;Build App & Database
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)
![Cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

## 👚 ERD

<img src="ERD4.drawio.png">

## Run Locally

Clone the project

```bash
  git clone https://github.com/faqihassyfa/PROJECT-III.git
```

Go to the project directory

```bash
  cd PROJECT-III
```

## Open Api

If you're interested in using our Open Api, this is an example of how to do so.

```bash
https://virtserver.swaggerhub.com/vaniliacahya/E-Store/1.0.0
```

<div>
      <details>
<summary>👶 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
This is an explanation of the Users section's CRUD method.
 
<div>
  
| Feature User | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| GET | /users  | - | YES | Users obtain their account information in this form.  |
| POST | /users | - | NO | This is how users register their account |
| Delete | /users | - | YES | Delete user account |
| PUT | /users | - | YES | Update user account |
| GET | /login | - | NO | This is how users log in. |

</details>

       
<div>
      <details>
<summary>🦊 Admin</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
Several commands make use of admin features, as shown below.
 
<div>
  
| Feature Admin | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| PUT | /admins/:productid  | ID Product | YES | create new restaurant |
| DELETE| /admins/:productid | ID Product | YES | edit the restaurant information |
| POST | /admins | - | YES | Create product |
| GET | /admins | - | YES | Displaying recently posted products |
| GET | /admins/history | - | YES | show the order history for existing orders |
