package com.github.mdeluise.tracky.integration.steps;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import com.github.mdeluise.tracky.authentication.payload.request.LoginRequest;
import com.github.mdeluise.tracky.authentication.payload.request.SignupRequest;
import com.github.mdeluise.tracky.observation.ObservationDTO;
import com.github.mdeluise.tracky.observation.ObservationService;
import com.github.mdeluise.tracky.tracker.Tracker;
import com.github.mdeluise.tracky.tracker.TrackerDTO;
import com.github.mdeluise.tracky.tracker.TrackerService;
import io.cucumber.java.DataTableType;
import io.cucumber.java.ParameterType;
import io.cucumber.java.en.And;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import io.cucumber.java.en.When;
import org.assertj.core.util.Strings;
import org.junit.Assert;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

import java.io.UnsupportedEncodingException;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Arrays;
import java.util.Date;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class IntegrationSteps {
    private final TrackerService trackerService;
    private final ObservationService observationService;
    private final UserService userService;
    private final MockMvc mockMvc;
    private final StepData stepData;
    private final ObjectMapper objectMapper;
    private final String signupEndpoint = "/authentication/signup";
    private final String loginEndpoint = "/authentication/login";
    private final String trackerEndpoint = "/tracker";
    private final String observationEndpoint = "/observation";


    public IntegrationSteps(TrackerService trackerService, ObservationService observationService,
                            UserService userService, MockMvc mockMvc, StepData stepData, ObjectMapper objectMapper) {
        this.trackerService = trackerService;
        this.observationService = observationService;
        this.userService = userService;
        this.mockMvc = mockMvc;
        this.stepData = stepData;
        this.objectMapper = objectMapper;
    }


    @When("call GET {string}")
    public void callGet(String url) throws Exception {
        stepData.setResponse(mockMvc.perform(MockMvcRequestBuilders.get(url).contentType(MediaType.APPLICATION_JSON)
                                                                   .header(HttpHeaders.AUTHORIZATION,
                                                                           stepData.getJwt().orElse("")
                                                                   )).andReturn());
    }


    @When("call POST {string} with body {string}")
    public void callPostWithBody(String url, String body) throws Exception {
        stepData.setResponse(mockMvc.perform(MockMvcRequestBuilders.post(url).contentType(MediaType.APPLICATION_JSON)
                                                                   .header(HttpHeaders.AUTHORIZATION,
                                                                           stepData.getJwt().orElse("")
                                                                   ).content(body)).andReturn());
    }


    @When("call PUT {string} with body {string}")
    public void callPutWithBody(String url, String body) throws Exception {
        stepData.setResponse(mockMvc.perform(MockMvcRequestBuilders.get(url).contentType(MediaType.APPLICATION_JSON)
                                                                   .header(HttpHeaders.AUTHORIZATION,
                                                                           stepData.getJwt().orElse("")
                                                                   ).content(body)).andReturn());
    }


    @When("call DELETE {string} with body {string}")
    public void callDeleteWithBody(String url, String body) throws Exception {
        stepData.setResponse(mockMvc.perform(MockMvcRequestBuilders.delete(url).contentType(MediaType.APPLICATION_JSON)
                                                                   .header(HttpHeaders.AUTHORIZATION,
                                                                           stepData.getJwt().orElse("")
                                                                   ).content(body)).andReturn());
    }


    @Then("receive status code of {int}")
    public void theClientReceivesStatusCode(int expectedStatus) throws UnsupportedEncodingException {
        Assert.assertEquals(stepData.getResponse(), expectedStatus, stepData.getResponseCode());
    }


    @Then("the count of observation for tracker {string} is {int}")
    public void countObservationsForUsers(String trackerName, int expectedCount) throws Exception {
        getAllTheObservations(trackerName);
        stepData.getResultAction().andExpect(MockMvcResultMatchers.jsonPath("$.numberOfElements").value(expectedCount));
    }


    @Then("the count of tracker is {int}")
    public void countTrackerForUsers(int expectedCount) throws Exception {
        stepData.setResultActions(mockMvc.perform(
            MockMvcRequestBuilders.get(trackerEndpoint).contentType(MediaType.APPLICATION_JSON)
                                  .header(HttpHeaders.AUTHORIZATION, stepData.getJwt().orElse(""))));
        Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
        stepData.getResultAction().andExpect(MockMvcResultMatchers.jsonPath("$.numberOfElements").value(expectedCount));
    }


    @ParameterType("(?:\\d+,\\s*)*\\d+")
    public List<Long> listOfLongs(String arg) {
        return Arrays.stream(arg.split(",\\s?")).sequential().map(Long::parseLong).collect(Collectors.toList());
    }


    @And("cleanup the environment")
    public void cleanupTheEnvironment() {
        observationService.removeAll();
        trackerService.removeAll();
        userService.removeAll();
        stepData.cleanup();
    }


    @When("remove observation(s) {listOfLongs}")
    public void removeObservations(List<Long> observationIds) throws Exception {
        String joinedIds = Strings.join(observationIds).with(",");
        callDeleteWithBody(String.format("%s/%s", observationEndpoint, joinedIds), "{}");
        Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
    }


    @When("remove tracker(s) {listOfLongs}")
    public void removeTrackers(List<Long> trackerIds) throws Exception {
        String joinedIds = Strings.join(trackerIds).with(",");
        callDeleteWithBody(String.format("%s/%s", trackerEndpoint, joinedIds), "{}");
        Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
    }


    @Given("the following users")
    public void theFollowingUsers(List<User> userList) throws Exception {
        for (User user : userList) {
            SignupRequest signupRequest = new SignupRequest(user.getUsername(), user.getPassword());
            callPostWithBody(signupEndpoint, objectMapper.writeValueAsString(signupRequest));
            Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
        }
    }


    @DataTableType
    public User userEntry(Map<String, String> entry) {
        User user = new User();
        user.setUsername(entry.get("username"));
        user.setPassword(entry.get("password"));
        return user;
    }


    @Given("the authenticated user with username {string} and password {string}")
    public void login(String username, String password) throws Exception {
        LoginRequest loginRequest = new LoginRequest(username, password);
        callPostWithBody(loginEndpoint, objectMapper.writeValueAsString(loginRequest));
        Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());

        // This should be done without parsing the response, using ModelMapper as:
        // modelMapper.map(stepData.getResponse(), UserInfoResponse.class)
        // but, the UserInfoResponse.class is a record, and ModelMapper does not support they yet
        Pattern pattern = Pattern.compile("(value\":)(\"[^\"]+\")");
        Matcher matcher = pattern.matcher(stepData.getResponse());
        if (matcher.find()) {
            String jwt = matcher.group(2).replaceAll("^\"+", "").replaceAll("\"+$", "");
            stepData.setJwt(jwt);
        }
    }


    @When("get all the trackers")
    public void getAllTheTrackers() throws Exception {
        stepData.setResultActions(mockMvc.perform(
            MockMvcRequestBuilders.get(trackerEndpoint).contentType(MediaType.APPLICATION_JSON)
                                  .header(HttpHeaders.AUTHORIZATION, stepData.getJwt().orElse(""))));
    }


    @When("get all the observation for tracker {string}")
    public void getAllTheObservations(String trackerName) throws Exception {
        Tracker tracker = trackerService.getByName(trackerName);
        stepData.setResultActions(mockMvc.perform(
            MockMvcRequestBuilders.get(String.format("%s?trackerId=%s", observationEndpoint, tracker.getId()))
                                  .contentType(MediaType.APPLICATION_JSON)
                                  .header(HttpHeaders.AUTHORIZATION, stepData.getJwt().orElse(""))));
    }


    @When("add the following observation(s)")
    public void addObservations(List<ObservationDTO> observations) throws Exception {
        for (ObservationDTO observation : observations) {
            callPostWithBody(observationEndpoint, objectMapper.writeValueAsString(observation));
            Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
        }
    }


    @DataTableType
    public ObservationDTO observationEntry(Map<String, String> entry) throws ParseException {
        ObservationDTO observation = new ObservationDTO();
        observation.setValue(Float.valueOf(entry.get("value")));
        if (entry.get("instant") != null) {
            SimpleDateFormat formatter = new SimpleDateFormat("dd-MM-yyyy hh:mm:ss", Locale.ENGLISH);
            observation.setInstant(formatter.parse(entry.get("instant")));
        } else {
            observation.setInstant(new Date());
        }
        Tracker tracker = trackerService.getByName(entry.get("tracker"));
        observation.setTrackerId(tracker.getId());
        return observation;
    }


    @When("create the following tracker(s)")
    public void createTrackers(List<TrackerDTO> trackers) throws Exception {
        for (TrackerDTO tracker : trackers) {
            callPostWithBody(trackerEndpoint, objectMapper.writeValueAsString(tracker));
            Assert.assertEquals(stepData.getResponse(), HttpStatus.OK.value(), stepData.getResponseCode());
        }
    }


    @DataTableType
    public TrackerDTO trackerEntry(Map<String, String> entry) throws ParseException {
        TrackerDTO tracker = new TrackerDTO();
        tracker.setName(entry.get("name"));
        tracker.setUnit(entry.get("unit"));
        tracker.setDescription(entry.get("description"));
        return tracker;
    }
}
