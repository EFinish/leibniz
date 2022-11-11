db = db.getSiblingDB('argument');

db.arguments.insertMany(
    [
        {
            "_id":ObjectId("52ffc33cd85242f436000001"),
            "title": "Argument 1",
        },
    ]
)

db.premises.insertMany(
    [
        {
            "_id":ObjectId("52ffc33cd85242f436000002"),
            "title": "Premise 1",
            "subject": {
                "$ref": "subjects",
                "$id": ObjectId("52ffc33cd85242f436000003"),
                "$db": "argument"
             },
             "predicate": {
                "$ref": "predicate",
                "$id": ObjectId("52ffc33cd85242f436000004"),
                "$db": "argument"
             },
        },
    ]
)

db.subjects.insertMany(
    [
        {
            "_id":ObjectId("52ffc33cd85242f436000003"),
            "title": "potatoes",
        },
    ]
)

db.predicates.insertMany(
    [
        {
            "_id":ObjectId("52ffc33cd85242f436000004"),
            "title": "grow in the earth",
        },
    ]
)