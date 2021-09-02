# challenge-yofio

# Database
The database required is with this schema:

Table name:   transaction_yf
Columns:      id  INT NOT NULL AUTO_INCREMENT
              success BIT NOT NULL

Required enviroment:
      CONNECTION_STRING = root:password1@tcp(127.0.0.1:3306)/test

# To generate build
    use .\cmd\buildaws.bat
    
# Tests
  Endpoint:  https://npir30dy5k.execute-api.us-east-1.amazonaws.com/test/

  Resources:
      POST: /credit-assignment
        - json Body Example:  { "Investment": 3000 }
      POST: /credit-assignments
        - json Body Example:  { "Investment": 3000 }
      POST: /statistics
