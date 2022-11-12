# candidate-compatibility
This project is a sample project to determine compatibility amongst an existing team and potential new members

# Business Process
- Resources are evaluated against an existing team by the following attributes: Intelligence, Strength, Endurance, and Spicy Food Tolerance
- The goal is to determine compatibility, which will be defined as the difference between the skill level of the candidate and against the average of the teams.
- Scores closer to 1 indicate the candidate skill level is close to or exceeds the teams average level across all attributes.
- Scores closer to 0 indicate the candidate skill level may have definciencies in 1 or more attributes. 

## Determination of Score
The applicant score and the team average score are measured as a percentage, to determine how far the applicant is from the team average. Once that percentage is calculated, it is weighted equally amoungst the other attributes at 25% of the total score. If a percentage exceeds 100%, it is rounded back down to 100%. The weighted percentage from all of the attributes is then added to determine the score 

## Example Calulation
```
John Doe        
=========
Intelligence         4
Strength             5
Endurance            2
Spicy Food Tolerance 1

Team Average
============
Intelligence         4.33 
Strength             3.66
Endurance            4.65
Spicy Food Tolerance 4

Score
=====
Note: Figures are rounded for simplicity

Intelligence         4 / 4.33 = 0.92 * 0.25 = 0.23  
Strength             5 / 3.66 = 1    * 0.25 = 0.25 (Note: Rounded down to 1)
Endurance            2 / 4.65 = 0.43 * 0.25 = 0.11 
Spicy Food Tolerance 1 / 4    = 0.25 * 0.25 = 0.06
                                              ----
                                      Score = 0.65


```

# Assumptions 
- Attributes are score from 0 - 10
- Scores are weighted equally accross 4 attributes (e.g., 0.25 for Intelligence, 0.25 for Strength, etc)

# Build Information
🚨 These instuctions are based on a OSX workstation using [homebrew](https://brew.sh/). If you already have docker installed you may need to check if the service is running before the `docker build` step

Clone the repo at your preferred workspace location

`git clone https://github.com/rmortal/candidate-compatibility.git`

This project will be using `Docker` to build and run the application. Docker can be installed via the following command:

`brew install --cask docker`

From the root directory of the project run the following command to create the `Docker` image to run:

`docker build --tag candidate-compatibility .`

Start the application container via the following command:

`docker run --publish 9000:9000 -v candidate-compatibility`

Press `Ctrl-C` to stop the container. 

# Documentation
This section outlines how generate the scores output.

1. Start the application container as described in the build information section.
2. Download [Postman](https://www.postman.com/) or tool of preferance to generate a GET request.
3. Set the URL for the GET request to http://localhost:9000/analyze-compatibility.
4. Within the body of the GET request use the example payload(adjust score values as desired). Set the body to raw and the type to JSON.
5. Click Send to view the response. 

### Postman configuration
![postmanConfig](https://drive.google.com/file/d/1syVdim9_uVuLGJP2AC__7hkVhwVKO6DK/view?usp=share_link)
### Sample payload
```
{
    "team": [
        {
            "name": "Ritch",
            "attributes": {
                "intelligence": 1,
                "strength": 5,
                "endurance": 3,
                "spicyFoodTolerance": 1
            }
        },
        {
            "name": "Alvin",
            "attributes": {
                "intelligence": 9,
                "strength": 4,
                "endurance": 1,
                "spicyFoodTolerance": 6
            }
        },
        {
            "name": "Albert",
            "attributes": {
                "intelligence": 3,
                "strength": 2,
                "endurance": 9,
                "spicyFoodTolerance": 5
            }
        }
    ],
    "applicants": [
        {
            "name": "Gene",
            "attributes": {
                "intelligence": 4,
                "strength": 5,
                "endurance": 2,
                "spicyFoodTolerance": 1
            }
        },
        {
            "name": "Kaitlin",
            "attributes": {
                "intelligence": 7,
                "strength": 4,
                "endurance": 3,
                "spicyFoodTolerance": 2
            }
        },
                {
            "name": "Alan",
            "attributes": {
                "intelligence": 10,
                "strength": 10,
                "endurance": 10,
                "spicyFoodTolerance": 10
            }
        }
    ]
}
```