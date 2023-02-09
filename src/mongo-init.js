db = db.getSiblingDB('admin');
// move to the admin db - always created in Mongo
db.auth("root", "root123");
// log as root admin if you decided to authenticate in your docker-compose file...
db = db.getSiblingDB('yt');
db.createCollection('videoinfo');
db.videoinfo.createIndex({description: "text"})