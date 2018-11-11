---
publishedOnWebsite: true
title: GraphQL Query API
subject: OSS
---

In this section, some example queries will be clarified. A full schema definition can be found HERE or at the GraphiQL IDE when you are running the prototype at `http://localhost:8081/graphql`.

##### Get local Things
Fetch information about all Things from a specified class.
In this case all airports will be returned, with information about properties references to cities and countries nested in these Things. Note that when the property refers to a class (like `InCity` and `InCountry`), these fields start with a capital (although this is not the case in the original schema where all property names start with small letters). These fields are filled with an object mentioning the Class as an object like `...on <Class> {<property>}`.

``` graphql
{
  Local{
    Get{
      Things{
        Airport{
          code
          name
          InCity{
            ...on City{
              name
              population
              coordinates
              isCapital
              InCountry{
                ...on Country{
                  name
                  population
                }
              }
            }
          }
        }
      }
    }
  }
}
```

##### Get local Things with filters
The query below returns information about all airports which lie in a city with a population bigger than 1,000,000, AND which also lies in the Netherlands.

``` graphql
{
  Local{
    Get(where:{
      operator: And,
      operands: [{
        path: ["Things", "Airport", "inCity", "City", "inCountry", "Country", "name"],
        operator:Equal,
        valueString: "Netherlands"
      },
      {
        path: ["Things", "Airport", "inCity", "City", "population"],
        operator: GreaterThan
        valueInt: 1000000
      }]
    }){
      Things{
        Airport{
          code
          name
          InCity{
            ...on City{
              name
              population
              coordinates
              isCapital
              InCountry{
                ...on Country{
                  name
                  population
                }
              }
            }
          }
        }
      }
    }
  }
}
```

##### Fetch Things in a converted way on a local Weaviate instance with arbitratry AND an OR filters
For the filter design, look [here](https://github.com/bobvanluijt/weaviate-graphql-prototype/wiki/Filter-design-pattern-for-Local-queries)
The query below returns all cities where either (at least one of the following conditions should hold):
- The population is between 1,000,000 and 10,000,000
- The city is a capital

``` graphql
{
  Local{
    Get(where:{
      operator: Or,
      operands: [{
        operator: And,
        operands: [{
          path: ["Things", "Airport", "inCity", "City", "population"],
          operator:LessThan,
          valueInt: 10000000
        },
        {
          path: ["Things", "Airport", "inCity", "City", "population"],
          operator: GreaterThanEqual
          valueInt: 1000000
        }],
      },{
        path: ["Things", "City", "isCapital"],
        valueBoolean:true
    }]
    }){
      Things{
        City{
        name
        population
        coordinates
        isCapital
        }
      }
    }
  }
}

```

##### Query generic meta data of Thing or Action classes
Generic meta data about classes and its properties can be queried. This does not count for individial nodes.
The query below returns metadata of the nodes in the class `City`.

``` graphql
{
  Local {
    GetMeta{
      Things {
        City {
          meta{
            count
          }
          inCountry{
            type
            count
            pointingTo
          }
          isCapital{
            type
            count
            totalTrue
            percentageTrue
          }
          population {
            type
            count
            lowest
            highest
            average
            sum
          }
          name{
            type
            count
            topOccurrences(first:2){
              value
              occurs
            }
          }
        }
      }
    }
  }
}

```

The same filters as the converted fetch can be used to filter the data. The following query returns the meta data of all cities of the Netherlands. 

``` graphql
{
  Local {
    GetMeta(where:{
      path: ["Things", "City", "inCountry", "name"],
      operator: Equal
      valueString: "Netherlands"
    }) {
      Things {
        City {
          meta{
            count
          }
          inCountry{
            type
            count
            pointingTo
          }
          isCapital{
            type
            count
            totalTrue
            percentageTrue
          }
          population {
            type
            count
            lowest
            highest
            average
            sum
          }
          name{
            type
            count
            topOccurrences(first:2){
              value
              occurs
            }
          }
        }
      }
    }
  }
}
```

##### Pagination
Pagination allows to request a certain amount of Things or Actions at one query. The arguments `first` and `after` can be combined in the query for classes of Things and Actions, where
- `first` is an integer with the maximum amount of returned nodes.
- `after` is an integer representing how many nodes should be skipped in the returned data.

``` graphql
{
  Local{
    Get{
      Things{
        City(first:10, after:2){
          name
          population
          coordinates
          isCapital
        }
      }
    }
  }
}
```