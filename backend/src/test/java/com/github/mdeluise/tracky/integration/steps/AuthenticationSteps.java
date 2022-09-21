package com.github.mdeluise.tracky.integration.steps;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.mdeluise.tracky.authentication.payload.request.LoginRequest;
import com.github.mdeluise.tracky.authentication.payload.response.UserInfoResponse;
import io.cucumber.java.en.Given;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;


public class AuthenticationSteps {
    private final String authPath = "/authentication";
    private final MockMvc mockMvc;
    private final StepData stepData;
    private final int port;
    private final ObjectMapper objectMapper;


    public AuthenticationSteps(@Value("${server.port}") int port, MockMvc mockMvc, StepData stepData,
                               ObjectMapper objectMapper) {
        this.port = port;
        this.mockMvc = mockMvc;
        this.stepData = stepData;
        this.objectMapper = objectMapper;
    }


    @Given("login with username {string} and password {string}")
    public void theClientLoginWithUsernameAndPassword(String username, String password) throws Exception {
        LoginRequest loginRequest = new LoginRequest(username, password);
        MvcResult result = mockMvc.perform(
            MockMvcRequestBuilders.post(String.format("http://localhost:%s%s/login", port, authPath))
                                  .contentType(MediaType.APPLICATION_JSON)
                                  .content(objectMapper.writeValueAsString(loginRequest))).andReturn();
        stepData.setResponse(result);
        if (result.getResponse().getStatus() == 200) {
            UserInfoResponse loginResponse =
                objectMapper.readValue(result.getResponse().getContentAsString(), UserInfoResponse.class);
            stepData.setJwt(loginResponse.jwt().value());
        }
    }
}
