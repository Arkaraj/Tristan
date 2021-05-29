# Go Portfolio Rest API

# Tech Stack

- Go/Golang
- Gin module (https://github.com/gin-gonic/gin)
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

- [ ] User can add details about themselves
- [x] User should be able to create skills
- [ ] User should be able to add/remove sills to their profile
- [ ] User should be able to create Projects
- [ ] User should be able to add/remove projects to their profile

### Routes

- GET /api/user/ Show User full details
- PUT /api/user/ Update User details
- POST /api/user User can add details about themselves
- GET /api/skills Get all skills
- POST /api/skills User can create skills
- DELETE /api/skills User can delete skills
- GET /api/projects/ Get all projects
- POST /api/projects User can create projects
- DELETE /api/projects User can delete projects
- POST /api/user/skill/:skillId add skill to User profile
- POST /api/user/project/:projectId add project to User profile

### Route Testing

- Postman

### To Run it Locally

Create a .env file:

```
MONGO_URL=
```
