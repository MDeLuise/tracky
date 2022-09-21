Feature: Test the correct functionality of the authentication.

  Scenario: call GET /observation without authentication. Request should be denied.
    When call GET "/observation"
    Then receive status code of 403

  Scenario: call GET /observation/0 without authentication. Request should be denied.
    When call GET "/observation?trackerId=0"
    Then receive status code of 403

  Scenario: call POST /observation without authentication. Request should be denied.
    When call POST "/observation" with body "{}"
    Then receive status code of 403

  Scenario: call PUT /observation/0 without authentication. Request should be denied.
    When call PUT "/observation/0" with body "{}"
    Then receive status code of 403

  Scenario: call DELETE /observation/0 without authentication. Request should be denied.
    When call DELETE "/observation/0" with body "{}"
    Then receive status code of 403

  Scenario: call GET /observation without authentication. Request should be denied.
    When call GET "/tracker"
    Then receive status code of 403

  Scenario: call GET /observation/0 without authentication. Request should be denied.
    When call GET "/tracker/0"
    Then receive status code of 403

  Scenario: call POST /observation without authentication. Request should be denied.
    When call POST "/tracker" with body "{}"
    Then receive status code of 403

  Scenario: call PUT /observation/0 without authentication. Request should be denied.
    When call PUT "/tracker/0" with body "{}"
    Then receive status code of 403

  Scenario: call DELETE /observation/0 without authentication. Request should be denied.
    When call DELETE "/tracker/0" with body "{}"
    Then receive status code of 403


  Scenario: Login correctly.
    Given the following users
      | id | username | password |
      | 1  | user1    | user1111 |
    And login with username "user1" and password "user1111"
    Then receive status code of 200

  Scenario: Login denied with username correct and wrong password.
    Given the following users
      | id | username | password |
      | 2  | user2    | user2222 |
    And login with username "user2" and password "wrongPassword"
    Then receive status code of 500

  Scenario: Login denied with username and password wrong.
    Given login with username "userDoesNotExist" and password "user"
    Then receive status code of 500