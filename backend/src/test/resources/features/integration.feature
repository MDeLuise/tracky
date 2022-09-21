Feature: Test the correct interchange of information between the components.

  Scenario: Create a user, create a tracker, then retrieve all trackers.
    Given the following users
      | username | password  |
      | newUser1 | password1 |
    And the authenticated user with username "newUser1" and password "password1"
    And create the following trackers
      | name     | description | unit |
      | tracker1 |             |      |
      | tracker2 | aaa         | bb   |
    Then get all the trackers
    And the count of tracker is 2
    And cleanup the environment

  Scenario: Create two user, some trackers, then retrieve all trackers from the second user.
    Given the following users
      | username | password  |
      | newUser1 | password1 |
      | newUser2 | password2 |
    And the authenticated user with username "newUser1" and password "password1"
    And create the following trackers
      | name     | description | unit |
      | tracker1 |             |      |
      | tracker2 | aaa         | bb   |
    And login with username "newUser2" and password "password2"
    And create the following trackers
      | name     | description | unit |
      | tracker1 |             |      |
    Then get all the trackers
    And the count of tracker is 1
    And cleanup the environment

  Scenario: Create one user, one trackers and observations, then retrieve all observations.
    Given the following users
      | username | password  |
      | newUser1 | password1 |
    And the authenticated user with username "newUser1" and password "password1"
    And create the following trackers
      | name     | description | unit |
      | tracker1 |             |      |
    And add the following observations
      | instant | value | tracker  |
      |         | 42    | tracker1 |
      |         | 42.42 | tracker1 |
    Then the count of observation for tracker "tracker1" is 2
    And cleanup the environment