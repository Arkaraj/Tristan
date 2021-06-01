# Go Portfolio Rest API

# Tech Stack

- Go/Golang
- Gin module (https://github.com/gin-gonic/gin)
- mongo Go Driver (https://github.com/mongodb/mongo-go-driver)
- NoSQL - MongoDB

Gin - Web Framework, like express for node

## About the Rest API

Personal Portfolio Rest api

Schema (in typescript):

```ts
interface IUser {
  name: string;
  description: string;
  tag: string;
  skills: ISkill[];
  projects: IProject[];
}

interface ISkill {
  id: string; // unique id, if i am using a db i.e
  sill: string;
  levelOfMastery: number; // Kindda like experience, out ot 5
}

interface IProject {
  id: string; // unique id, if i am using a db i.e
  title: string;
  description: string;
  skills: ISkill[];
  link: string;
  startTime: Date | string;
  endTime: Date | string;
}
```

- [x] User can add details about themselves
- [x] User should be able to create skills
- [x] User should be able to add/remove sills to their profile
- [x] User should be able to create Projects
- [x] User should be able to add/remove projects to their profile

### Routes

- GET /api/user/ Show User full details
- PUT /api/user/ Update User details
- POST /api/user User can add details about themselves
- GET /api/skill Get all skills
- POST /api/skill User can create skills
- DELETE /api/skill/:skillId User can delete skills
- GET /api/project/ Get all projects
- POST /api/project User can create projects
- PUT /api/project/:projectId/:skillId User can add skills to projects
- DELETE /api/project/:projectId User can delete projects
- POST /api/user/skill/:skillId add skill to User profile
- POST /api/user/project/:projectId add project to User profile

### Route Testing

- Postman

Link: https://documenter.getpostman.com/view/8802598/TzY1iH55

### To Run it Locally

Create a .env file:

```
MONGO_URL=
```
