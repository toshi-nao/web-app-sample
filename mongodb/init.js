db = db.getSiblingDB("tutorial");
db.createCollection('tutorial_collection');

db.tutorial_collection.insertMany([
    {
        "title": "Tutorial 1",
        "description": "This is tutorial 1 from MongoDB",
        "published": false,
        "createdAt": "2023-05-01T02:53:48.690Z",
        "updatedAt": "2023-05-01T02:53:48.690Z",
        "id": "644f29bc21eb646280e59c84"
    },
    {
        "title": "Tutorial 2",
        "description": "This is tutorial 2 MongoDB",
        "published": false,
        "createdAt": "2023-05-01T02:53:48.690Z",
        "updatedAt": "2023-05-01T02:53:48.690Z",
        "id": "644f29c2acefe3e62a866731"
    },
    {
        "title": "Tutorial 3",
        "description": "This is tutorial 3 MongoDB",
        "published": false,
        "createdAt": "2023-05-01T02:53:48.690Z",
        "updatedAt": "2023-05-01T02:53:48.690Z",
        "id": "644f2a0eacefe3e62a866735"
    }
])