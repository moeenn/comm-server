-- name: GetPendingNotificationById :one
SELECT * FROM "PendingNotification"
WHERE "id" = $1 LIMIT 1;


-- name: GetUserPendingNotifications :many
SELECT * FROM "PendingNotification"
WHERE "userId" = $1
ORDER BY "createdAt" ASC;


-- name: CreatePendingNotification :exec
INSERT INTO "PendingNotification" ("id", "userId", "payload")
VALUES ($1, $2, $3);


-- name: DeletePendingNotificationById :exec
DELETE FROM "PendingNotification"
WHERE "id" = $1;


-- name: DeletePendingNotificationByUserId :exec
DELETE FROM "PendingNotification"
WHERE "userId" = $1;
