-- name: FindUserByMssv :one
SELECT 
    id, name, age, mssv 
FROM users 
WHERE mssv = $1;
