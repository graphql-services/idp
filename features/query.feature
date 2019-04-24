Feature: Managing users

    Scenario: Add user
        When I send query:
            """
            mutation { createUser(input:{email:"john.doe@example.com",password:"top_secret"}) { email } }
            """
        Then the response should be:
            """
            {
            "createUser" : { "email":"john.doe@example.com" }
            }
            """
    Scenario: Login
        Given I send query:
            """
            mutation { createUser(input:{email:"john.doe2@example.com",password:"top_secret2"}) { email } }
            """
        When I send query:
            """
            query { login(email:"john.doe2@example.com",password:"top_secret2") { email email_verified } }
            """
        Then the response should be:
            """
            {
            "login" : { "email":"john.doe2@example.com","email_verified": false }
            }
            """
