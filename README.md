# task-management-app

Before start execute next commands:

1. make postgres - for pulling and set up postgress container
2. make psql - for star postgress container
3. make createdb - for crearing db
4. make migrateup - for runnig migration(https://github.com/golang-migrate/migrate should be installed)


5. prepare app.toml config file in accordance with example
5. make for building app



# Endpoints description


## Project actions

### "POST"    /project/new - creates new project
#### req:
```json
content-type: application/json

{
    "name": string,
    "description": string
}
```
### "GET"     /projects    - returns all projects list
### "GET"     /project/id  - returns project with tasks 
### "PUT"     /project/id  - updates project name and description
#### req:
```json
content-type: application/json

{
    "name":string,
    "description":string
}
```
### "DELETE"  /project/id  - delete project


## Column actions

### "POST"    /column/new      - creates new column
#### req:
```json
content-type: application/json

{
    "name":string,
    "project_id": uint,
    "position":uint
}
```
### "PUT"     /column/id       - updates column name
#### req:
```json
content-type: application/json

{
    "name":string,
    "project_id": uint
    
}
```
### "PUT"     /column/position - updates columns positions
#### req:
```json
content-type: application/json

{   
    "positions":[
        {
            "id": columnID - uint,
            "position": new position - uint 

        },
        {
            "id": columnID - uint,
            "position": new position - uint 
        }
    ]
 
}
```
### "DELETE"  /column/id       - delete column 

## Task actions

### "POST"    /task/new      - creates new task
#### req:
```json
content-type: application/json

{
    
    "name":string,
    "column_id": uint,
    "priority":uint,
    "description":string

}
```
### "GET"     /task/id       - returns task with comments
### "DELETE"  /task/id       - delete task with comments
### "PUT"     /task/id       - updates task description and name
#### req:
```json
content-type: application/json
{
    
    "name":string,
    "column_id": uint,
    "description":string

}
```
### "PUT"     /task/priority - updates tasks pririty
#### req:
```json
content-type: application/json

{   
    "priorities":[
        {
            "id": task id - uint,
            "priority": new task priority - uint

        },
        {
            "id": task id - uint,
            "priority": new task priority - uint

        }
    ]
 
}
```

## Comment actions

### "POST"    /comment/new - creates new comment
#### req:
```json
content-type: application/json

{
    
    "text":string,
    "task_id": uint

}
```
### "PUT"     /comment/id  - updates comment text
#### req:
```json
content-type: application/json

{
    
   "text":string,
    "id": uint

}
```
### "DELETE"  /comment/id  - delete comment