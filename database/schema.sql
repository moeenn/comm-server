CREATE TABLE PendingNotification (
    id TEXT NOT NULL,
    userId TEXT NOT NULL,
    payload JSONB NOT NULL,
    createdAt TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "PendingNotification_pkey" PRIMARY KEY ("id")
);