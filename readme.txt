On windown operating system, cd into folder and run start main.exe
I'm still not finish unit tests because I know just a little about it and learning.

API example

GET http://localhost:8000/v1/user/getall
GET http://localhost:8000/v1/user/get/6280c64d646de4bbab6be5b6
POST http://localhost:8000/v1/user/create
{
    "name": "Linh",
    "age": 12,
    "task": 5
}
PATCH http://localhost:8000/v1/user/update
{   
    "_id": "62865b9314cc59afc6ae2d43",
    "nane": "Linh",
    "age": 32,
    "task": 5
}
DELETE http://localhost:8000/v1/user/delete/62865b9314cc59afc6ae2d43



GET http://localhost:8000/v1/task/getall
GET http://localhost:8000/v1/task/get/62884e2a5af93badb7f29d82
POST http://localhost:8000/v1/task/create
{
    "taskname": "test18",
    "owner": "6280c64d646de4bbab6be5b6",
    "status": "done"
}
PATCH http://localhost:8000/v1/task/update
{   
    "_id": "628a64d941acce2b53a97f7a",
    "taskname": "test19",
    "owner": "6280c64d646de4bbab6be5b6",
    "status": "done",
    "date": "20220521"
}
DELETE http://localhost:8000/v1/task/delete/628a64d741acce2b53a97f79