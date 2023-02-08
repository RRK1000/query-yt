db.createUser(
    {
      user: "poll-svc",
      pwd: "password",
      roles: [
        {
          role: "readWrite",
          db: "yt"
        }
      ]
    }
);
db.createCollection('videoinfo');