# Token
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI1MzYwNTYsInVzZXJfaWQiOjMsInJvbGVfaWQiOjN9.QPcyFt9lRMW5KttOUBvJUPl8kMbCS4zwJn0vquT1UMI
```
# **Admin**
## Faculty

### add faculty
```bash
curl -X POST http://localhost:PORT/admin/faculty \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Faculty Name",
  "code": "FAC001"
}'
```

### update faculty
```bash
curl -X PUT http://localhost:PORT/admin/faculty/{fid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Updated Faculty Name",
  "code": "FAC002"
}'
```

### delete faculty
```bash
curl -X DELETE http://localhost:PORT/admin/faculty/{fid} \
-H "Authorization: {token}" 
```

### get by id
```bash
curl -X GET http://localhost:PORT/admin/faculty/{fid} \
-H "Authorization: {token}" 
```

### get faculties
```bash
curl -X GET "http://localhost:PORT/admin/faculty?id=0&name=&code=&limit=10&page=1"
-H "Authorization: {token}" 
```

## Department

### add department
```bash
curl -X POST http://localhost:PORT/admin/department \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Department Name",
  "code": "DEP001",
  "faculty_id": 1
}'
```

### update department
```bash
curl -X PUT http://localhost:PORT/admin/department/{did} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Updated Department Name",
  "code": "DEP002",
  "faculty_id": 2
}'
```

### delete department
```bash
curl -X DELETE http://localhost:PORT/admin/department/{did} \
-H "Authorization: {token}" 
```

### get by id
```bash
curl -X GET http://localhost:PORT/admin/department/{did} \
-H "Authorization: {token}" 
```

### get departments
```bash
curl -X GET "http://localhost:PORT/admin/department?id=0&name=&code=&faculty_id=0&limit=10&page=1" \
-H "Authorization: {token}" 
```

## Profession

### add profession
```bash
curl -X POST http://localhost:PORT/admin/profession \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Profession Name",
  "code": "PRO001",
  "department_id": 1
}'
```

### update profession
```bash
curl -X PUT http://localhost:PORT/admin/profession/{pid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Updated Profession Name",
  "code": "PRO002",
  "department_id": 2
}'
```

### delete profession
```bash
curl -X DELETE http://localhost:PORT/admin/profession/{pid} \
-H "Authorization: {token}" 
```

### get by id
```bash
curl -X GET http://localhost:PORT/admin/profession/{pid} \
-H "Authorization: {token}" 
```

### get professions
```bash
curl -X GET "http://localhost:PORT/admin/profession?id=0&name=&code=&department_id=0&limit=10&page=1" \
-H "Authorization: {token}" 
```

## Group

### add group
```bash
curl -X POST http://localhost:PORT/admin/group \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Group Name",
  "code": "GRP001",
  "year": 2024,
  "profession_id": 1
}'
```

### update group
```bash
curl -X PUT http://localhost:PORT/admin/group/{gid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Updated Group Name",
  "code": "GRP002",
  "year": 2025,
  "profession_id": 2
}'
```

### delete group
```bash
curl -X DELETE http://localhost:PORT/admin/group/{gid} \
-H "Authorization: {token}" 
```

### get by id
```bash
curl -X GET http://localhost:PORT/admin/group/{gid} \
-H "Authorization: {token}" 
```

### get groups
```bash
curl -X GET "http://localhost:PORT/admin/group?id=0&name=&code=&year=0&profession_id=0&limit=10&page=1" \
-H "Authorization: {token}" 
```

## Lesson

### add lesson
```bash
curl -X POST http://localhost:PORT/admin/lesson \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Lesson Name",
  "code": "LES001"
}'
```

### update lesson
```bash
curl -X PUT http://localhost:PORT/admin/lesson/{lid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "name": "Updated Lesson Name",
  "code": "LES002"
}'
```

### delete lesson
```bash
curl -X DELETE http://localhost:PORT/admin/lesson/{lid} \
-H "Authorization: {token}" 
```

### get by id
```bash
curl -X GET http://localhost:PORT/admin/lesson/{lid} \
-H "Authorization: {token}" 
```

### get lessons
```bash
curl -X GET "http://localhost:PORT/admin/lesson?id=0&name=&code=&limit=10&page=1" \
-H "Authorization: {token}" 
```

## Time
### add time
```bash
curl -X POST http://localhost:PORT/admin/time \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "start_time": "09:00",
  "end_time": "17:00"
}'
```

### update time
```bash
curl -X PUT http://localhost:PORT/admin/time/{tid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "start_time": "10:00",
  "end_time": "18:00"
}'

```

### delete time
```bash
curl -X DELETE http://localhost:PORT/admin/time/{tid} \
-H "Authorization: {token}" 
```

### get time by id
```bash
curl -X GET http://localhost:PORT/admin/time/{tid} \
-H "Authorization: {token}" 
```

### get times
```bash
curl -X GET "http://localhost:PORT/admin/time?id=0&start_time=&end_time=&limit=10&page=1" \
-H "Authorization: {token}" 
```

## Lesson Teacher Student Binding 

### add lesson teacher student binding
```bash
curl -X POST http://localhost:PORT/admin/lesson-teacher-student-binding \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "student_id": 123,
  "teacher_id": 456,
  "group_id": 1,
  "lesson_id": 101,
  "type_id": 1
}'
```

### delete lesson teacher student binding
```bash
curl -X DELETE http://localhost:PORT/admin/lesson-teacher-student-binding \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "student_id": 123,
  "teacher_id": 456,
  "lesson_id": 101
}'
```

### get lesson teacher student binding
```bash
curl -X GET "http://localhost:PORT/admin/lesson-teacher-student-binding?teacherId=456&lessonId=101" \
-H "Authorization: {token}" 
```

## Timetable

### add Timetable
```bash
curl -X POST http://localhost:PORT/admin/timetable \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "weekday": 1,
  "group_id": 1,
  "lesson_id": 101,
  "time_id": 1,
  "auditory_id": 1,
  "alt_lesson_id": 102,
  "alt_auditory_id": 2,
  "type_id": 1
}'
```

### delete timetable
```bash
curl -X DELETE http://localhost:PORT/admin/timetable/{timetableID} \
-H "Authorization: {token}" 
```

### get timetable
```bash
curl -X GET http://localhost:PORT/admin/timetable/{group_id} \
-H "Authorization: {token}" 
```

## Student 

### add Student
```bash
curl -X POST http://localhost:PORT/admin/student \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "first_name": "John",
  "last_name": "Doe",
  "code": "S12345",
  "gender": true,
  "birth_date": "2000-01-01",
  "group_id": 1,
  "group_name": "Group A",
  "username": "johndoe",
  "password": "password123"
}'
```

### update student
```bash
curl -X PUT http://localhost:PORT/admin/student/{sid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "first_name": "John",
  "last_name": "Doe",
  "code": "S12345",
  "gender": true,
  "birth_date": "2000-01-01",
  "group_id": 1,
  "group_name": "Group A",
  "username": "johndoe_updated",
  "password": "newpassword123"
}'
```

### delete student
```bash
curl -X DELETE http://localhost:PORT/admin/student/{sid} \
-H "Authorization: {token}" 
```

### get student by id
```bash
curl -X GET http://localhost:PORT/admin/student/{sid} \
-H "Authorization: {token}" 
```

### get students
```bash
curl -X GET "http://localhost:PORT/admin/student?id=0&first_name=John&last_name=Doe&username=johndoe&code=S12345&limit=10&page=1" \
-H "Authorization: {token}" 
```

### upload student photo
```bash
curl -X POST http://localhost:PORT/admin/student/{sid}/image \
-H "Authorization: {token}" \
-F "image=@/path/to/image.jpg"
```

## Absence
### get absents
```bash
curl -X POST http://localhost:8080/admin/absences \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "faculty_id": 1,
  "department_id": 2,
  "group_id": 3,
  "lesson_id": 4,
  "type_id": 5,
  "teacher_id": 6,
  "student_id": 7,
  "student_first_name": "John",
  "student_last_name": "Doe",
  "from": "2024-01-01",
  "to": "2024-12-31"
}'
```

## Teacher 

### add Teacher
```bash
curl -X POST http://localhost:PORT/admin/teacher \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "first_name": "Jane",
  "last_name": "Smith",
  "code": "T12345",
  "gender": true,
  "username": "janesmith",
  "password": "password123",
  "department_id": 1,
  "department_name": "Mathematics"
}'
```

### update teacher
```bash
curl -X PUT http://localhost:PORT/admin/teacher/{tid} \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "first_name": "Jane",
  "last_name": "Smith",
  "code": "T12345",
  "gender": true,
  "username": "janesmith_updated",
  "password": "newpassword123",
  "department_id": 1,
  "department_name": "Mathematics"
}'
```

### delete teacher
```bash
curl -X DELETE http://localhost:PORT/admin/teacher/{tid} \
-H "Authorization: {token}" 
```

### get teacher by id
```bash
curl -X GET http://localhost:PORT/admin/teacher/{tid} \
-H "Authorization: {token}" 
```

### get teachers
```bash
curl -X GET "http://localhost:PORT/admin/teacher?id=0&first_name=Jane&last_name=Smith&username=janesmith&code=T12345&department_id=1&limit=10&page=1" \
-H "Authorization: {token}" 
```

### upload teacher photo
```bash
curl -X POST http://localhost:PORT/admin/teacher/{tid}/image \
-H "Authorization: {token}" \
-F "image=@/path/to/image.jpg"
```

# **Client**

## Sign In

### sign In
```bash
curl -X POST http://localhost:PORT/v1/signin \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
  "username": "user1",
  "password": "password123"
}'
```


### get students
```bash
curl -X GET "http://localhost:PORT/v1/students?group_id={groupID}" \
-H "Authorization: {token}"
```

### get faculties
```bash
curl -X GET "http://localhost:PORT/v1/faculties" \
-H "Authorization: {token}"
```

### get departments
```bash
curl -X GET "http://localhost:PORT/v1/departments?faculty_id={groupID}" \
-H "Authorization: {token}"
```

### get groups
```bash
curl -X GET "http://localhost:PORT/v1/groups?department_id={groupID}" \
-H "Authorization: {token}"
```

### get lessons
```bash
curl -X GET "http://localhost:PORT/v1/lessons" \
-H "Authorization: {token}"
```

### get types
```bash
curl -X GET "http://localhost:PORT/v1/types" \
-H "Authorization: {token}"
```

### get times
```bash
curl -X GET "http://localhost:PORT/v1/times" \
-H "Authorization: {token}"
```

### check students for existence
```bash
curl -X POST "http://localhost:PORT/v1/check-students-for-existence" \
-H "Content-Type: application/json" \
-H "Authorization: {token}" \
-d '{
    "group_id": {groupID},
    "lesson_id": {lessonID},
    "time_id": {timeID},
    "type_id": {typeID},
    "student_id": {studentID}
}'
```


