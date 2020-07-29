# simple-ecommerce-mongodb
Make Ecommerce with MongoDB 

This project already filled with the vendors

To run the project:
    "go run server.go"

----------------------------------------------------------------------------------------
To See the API Document: https://documenter.getpostman.com/view/2916835/Szzn7cKL
To See Sistem and Database Design: https://1drv.ms/u/s!Aos2oMKHXLjOtEU_kC7ZuK0p_IfT?e=zo6TK1
Available API : user, store, product, category, courier, account, search

There are 2 User (Admin & User)

Admin can do CRUD for Master DATA (category, courier, account)

User can be a regular user & store owner -> if user.havestore == 1

user who havestore == 1 can do CRUD for product

----------------------------------------------------------------------------------------
You can register and login with 2 method:
    --Github
    --Manual Register and Login

After Register User/Store, copy the verification token, and go to api verify and paste the token, to activate the account

----------------------------------------------------------------------------------------
Don't Forget to Import the Database (in folder db)

database name: simple_ecommerce

----------------------------------------------------------------------------------------
Account for admin :
    user: admin@test.com
    pass: testtest

Account for User(HaveStore) :
    user: user@test.com
    pass: testtest

----------------------------------------------------------------------------------------
Don't forget to add the token in header (example in API Docs)
you can get the token after login

----------------------------------------------------------------------------------------
