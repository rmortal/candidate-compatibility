# candidate-compatibility
This project is a sample project to determine compatibility amongst an existing team and potential new members

# Business Process
- Resources are evaluated against an existing team by the following attributes: Intelligence, Strength, Endurance, and Spicy Food Tolerance
- The goal is to determine compatibility, which will be defined as the difference between the skill level of the candidate and against the average of the teams.
- Scores closer to 1 indicate the candidate skill level is close to or exceeds the teams average level across all attributes.
- Scores closer to 0 indicate the candidate skill level may have definciencies in 1 or more attributes. 

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