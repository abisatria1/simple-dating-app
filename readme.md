### Simple Bumble Application
Simple dating app backend application. Include functionality to view, like and match each others.

#### How to run the application
1. Make sure mysql already installed in computer
2. Run mysql database
3. Create .env file or just copy from .env.example and rename it to .env 
4. run command ```make migration``` to run migration
5. run command ```make seed``` to run seeder
6. run command ```make run``` to start the backend application


#### Folder structure
* `cmd` folder 
This folder is to store all command related to our application. For example: if application run cron, the main file we can store in cmd folder 

* `pkg` folder 
This folder is to store all common customization for the entire application. 

* `src` folder 
This folder will contains all file related to our backend application logic. 
    - In this src folder there was 3 layer. The first layer is `handler` layer -> `usecase` layer -> `domain` layer 
    - Put all the API handler, CRON handler or message queue handler in `handler` folder. 
    - Put all the business logic inside usecase function. This layer is connecting several `domain` layer
    - Put all the behavior of an entity inside the `domain` layer. 
* `service`
This folder is used to store all module initation in the application. 
* `model` This folder is to used to store all type, request , response and all typing related to application. 
* `middleware` This folder is to used to store all middleware function in application. 
* `constant` This folder used to store all constant in application
* `config` This folder used to store all configuration in application


#### How to run unit test 
Simply run command `make test` to run unit test in this application 