# Database Documentation

From version 2.0 of our software onward we switched from golang and sqlite to using mongodb in combination with flask on our backend. 
As before, both of these as well as our frontend run locally on the rapsberry pi in a container environment.
The change in database technology was mainly driven by the fact that our data structure did change frequently. While our data is relatively structured, we also do not require a very high level of integrity and were happy about less of an overhead and a smoother development process.

## NoSQL Schema design

First of all, here is relational model that we had to translate (with changes):

![enter image description here](https://github.com/Emil9999/ROWA/blob/feature/backendRewrite/documentation/Database%20Rowa.png?raw=true)
As we can see there are multiple one-to-many relationships. One module has multiple plants and one plant type can have multiple modules.
If straightforward translate this to the bsjon/json format we get something like this:

    {
    "planttype": {
        "name": "basil",
        "growth_time": 42,
        ....

    },

    "plant": {
		    "_id": "156dfef5",
        "variety": "basil",
        "plant_date": "01/01/2022",
        "user_name": "Emil",
        ...
    },
    
    "module": {
        "modulenum": 1,
        "plants": [plant_1{..}, plant_2{}, ...],
        "plantable_varieties": ["basil", "oregano", "mint"]
        ...
    }

While tempting, I resisted the urge to simply keep thinking in tables and together with my colleagues (of which some had no idea about databases, which was really helpful) rethought our structure from zero, hence the many changes. 


In mongo we have two options to model one-to-many relationships: **Embedded documents** and **document references**. 

## Embedded documents vs. document references

### Embedded documents:

I decided to use embedded documents for the relationship between plants and modules. This means that every plant document is embedded in the respective module. 
As I often have to fetch all plants within a module, using references would have meant multiple queries for this one operation. Therefore embedding the plants is actually faster as they can all be fetched in a single query.

One problem with this approach is that it can lead to too big documents, however our field is bound, i.e the number of plants per module is capped to a max.

### Document references:

For our second example, consider every plant having a plant type and every plant type having many plants. In this case embedding the plant type within the plant would lead to repetition. 
One way of using references could be by embedding the plant _id in the plant type:

     {
    "planttype": {
        "name": "basil",
        "growth_time": 42,
        "plants":["156dfef5", ...]
        ....
    },

    "plant": {
        "_id": "156dfef5",
        "plant_date": "01/01/2022",
        "user_name": "Emil",
        ...
    }


The downside here is we now have a mutable and unbound array which can significantly impact performance and even lead to exceeding the BSON document size limit.

Therefore I chose this approach:

     {
    "planttype": {
        "name": "basil",
        "growth_time": 42,
        ....
    },

    "plant": {
        "_id": "156dfef5",
        "variety": "basil",
        "plant_date": "01/01/2022",
        "user_name": "Emil",
        ...
    }

## Replication
Once again I want to note that our DB is running locally on a raspberry pi that is in every farm. The amount of data is currently limited and frankly it's loss would be a minor problem for our business. However in the future payment history along with a phone app to access farm data are very much within the possible and therefore the data cannot only live locally anymore. 

The natural way of going about this in mongo is to use a replica set. This basically signifies a group of two or more instances of the same database running in the same or in geographically separate systems. This graphic what we are currently trying to do:

![enter image description here](https://github.com/Emil9999/ROWA/blob/feature/backendRewrite/documentation/replication.png?raw=true)
As we can see our replica set has three members. Our primary is by default the one receiving all reads and writes. The oplog (operation log) is then asynchronously transferred to the secondaries. In case our primary becomes unavailable and election is started between the secondaries to define a new primary and we can continue reading and writing our data. 
Elections are also the reason we need at least three members, otherwise votes do not make sense. In case we cannot replicate our data twice we could also have an arbiter instance that doesn't hold any data but takes part in elections. However in our case such a limitation doesn't exist and therefore it makes sense to use 3 members.
