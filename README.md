<img src="https://lh6.ggpht.com/gNy40q6S_519oQZ_AE9sGypZ-Z94zDy2Xpm5Tg5mYf8yVOSLAxAhEatKLn0vJDyFErE=w300" width="80"/>

# GUC API

REST API wrapper for the German University in Cairo (GUC) ~~private~~ API.

## Why?

* The GUC API is only exclusively used by the official GUC mobile application
* The GUC API is altogether poorly designed _(e.g. JSON embedded within XML responses)_

## API

### Authentication

All API calls require [basic authentication](https://en.wikipedia.org/wiki/Basic_access_authentication#Client_side).
Example: if your username is `john.doe` & your password is `12345`, then your HTTP `Authorization` header should look like this: `Basic Z3VjaWFuOjEyMzQ1`.

### API Calls

#### Login

<pre><b>GET</b>http://guc-api.herokuapp.com/api/<b><i>login</i></b></pre>

Response:
```javascript
{
    "authorized": true
}
```
or
```javascript
{
    "authorized": false
}
```

***

#### Coursework

<pre><b>GET</b>http://guc-api.herokuapp.com/api/<b><i>coursework</i></b></pre>

Response:
```javascript
{  
   "error": null,
   "data": [  
      {  
         "code": "CSEN701",
         "name": "Embedded System Architecture",
         "grades": [  
            {  
               "module": "Assignment 1",
               "point": "9.75",
               "maxPoint": "10"
            },
            ...
         ]
      },
      ...
   ]
}
```

***

#### Midterms

<pre><b>GET</b>http://guc-api.herokuapp.com/api/<b><i>midterms</i></b></pre>

Response:
```javascript
{  
   "error": null,
   "data": [  
      {  
         "name": "MET Computer Science 7th Semester - Analysis and Design of Algorithms CSEN703",
         "percentage": "41.25"
      },
      ...
   ]
}
```

***

#### Attendance

<pre><b>GET</b>http://guc-api.herokuapp.com/api/<b><i>attendance</i></b></pre>

Response:
```javascript
{  
   "error": null,
   "data": [  
      {  
         "name": "Computer Graphics",
         "level": "1"
      },
      ...
   ]
}
```

## Limitations

The GUC servers go down quite often. Transitively, our API cannot serve anything during that time.

## License

This project is licensed under the MIT License.
