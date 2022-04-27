# Database Documentation

From version 2.0 onward we switched from golang and sqlite to using mongodb in combination with flask on our backend. 
As before, both of these as well as our frontend run locally on the rapsberry pi in a container environment.


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

