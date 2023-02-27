FROM mongo

COPY subjects-data.json /subjects-data.json
COPY predicates-data.json /predicates-data.json
CMD mongoimport --host mongo --db argument --collection subjects --file /subjects-data.json --jsonArray \
    && mongoimport --host mongo --db argument --collection predicates --file /predicates-data.json --jsonArray
