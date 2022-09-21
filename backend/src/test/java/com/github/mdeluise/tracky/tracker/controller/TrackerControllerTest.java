package com.github.mdeluise.tracky.tracker.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.mdeluise.tracky.TestEnvironment;
import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
import com.github.mdeluise.tracky.security.apikey.ApiKeyFilter;
import com.github.mdeluise.tracky.security.apikey.ApiKeyRepository;
import com.github.mdeluise.tracky.security.apikey.ApiKeyService;
import com.github.mdeluise.tracky.security.jwt.JwtTokenFilter;
import com.github.mdeluise.tracky.security.jwt.JwtTokenUtil;
import com.github.mdeluise.tracky.security.jwt.JwtWebUtil;
import com.github.mdeluise.tracky.tracker.Tracker;
import com.github.mdeluise.tracky.tracker.TrackerController;
import com.github.mdeluise.tracky.tracker.TrackerDTO;
import com.github.mdeluise.tracky.tracker.TrackerDTOConverter;
import com.github.mdeluise.tracky.tracker.TrackerService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.annotation.Import;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.http.MediaType;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

import java.util.List;

@WebMvcTest(TrackerController.class)
@AutoConfigureMockMvc(addFilters = false)
@WithMockUser(username = "admin", roles = "ADMIN")
@Import(TestEnvironment.class)
public class TrackerControllerTest {
    @MockBean
    private JwtTokenFilter jwtTokenFilter;
    @MockBean
    private JwtTokenUtil jwtTokenUtil;
    @MockBean
    private JwtWebUtil jwtWebUtil;
    @MockBean
    private ApiKeyFilter apiKeyFilter;
    @MockBean
    private ApiKeyService apiKeyService;
    @MockBean
    private ApiKeyRepository apiKeyRepository;
    @MockBean
    private TrackerService trackerService;
    @MockBean
    private TrackerDTOConverter trackerDTOConverter;
    @MockBean
    private UserService userService;
    @Autowired
    private ObjectMapper objectMapper;
    @Autowired
    private MockMvc mockMvc;


    @Test
    void whenGetTrackers_ShouldReturnTrackers() throws Exception {
        Tracker tracker1 = new Tracker();
        tracker1.setId(1L);
        TrackerDTO trackerDTO1 = new TrackerDTO();
        trackerDTO1.setId(1L);
        Tracker tracker2 = new Tracker();
        tracker2.setId(2L);
        TrackerDTO trackerDTO2 = new TrackerDTO();
        trackerDTO2.setId(2L);
        Mockito.when(trackerService.getAll(PageRequest.of(0, 10, Sort.Direction.ASC, "id")))
               .thenReturn(new PageImpl<>(List.of(tracker1, tracker2)));
        Mockito.when(trackerDTOConverter.convertToDTO(tracker1)).thenReturn(trackerDTO1);
        Mockito.when(trackerDTOConverter.convertToDTO(tracker2)).thenReturn(trackerDTO2);

        mockMvc.perform(MockMvcRequestBuilders.get("/tracker")).andExpect(MockMvcResultMatchers.status().isOk())
               .andExpect(MockMvcResultMatchers.jsonPath("$.content", org.hamcrest.Matchers.hasSize(2)));
    }


    @Test
    void whenDeleteTracker_shouldReturnOk() throws Exception {
        Mockito.doNothing().when(trackerService).remove(0L);

        mockMvc.perform(MockMvcRequestBuilders.delete("/tracker/0")).andExpect(MockMvcResultMatchers.status().isOk());
    }


    @Test
    void whenDeleteNonExistingTracker_shouldError() throws Exception {
        Mockito.doThrow(ResourceNotFoundException.class).when(trackerService).remove(List.of(0L));

        mockMvc.perform(MockMvcRequestBuilders.delete("/tracker/0"))
               .andExpect(MockMvcResultMatchers.status().is4xxClientError());
    }


    @Test
    void whenCreateTracker_shouldReturnTracker() throws Exception {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Tracker created = new Tracker();
        created.setId(0L);
        created.setUser(user);
        TrackerDTO createdDTO = new TrackerDTO();
        createdDTO.setId(0L);
        Mockito.when(trackerService.save(created)).thenReturn(created);
        Mockito.when(trackerDTOConverter.convertToDTO(created)).thenReturn(createdDTO);
        Mockito.when(trackerDTOConverter.convertFromDTO(createdDTO)).thenReturn(created);
        Mockito.when(userService.get("admin")).thenReturn(user);

        mockMvc.perform(MockMvcRequestBuilders.post("/tracker").content(
                                                  objectMapper.writeValueAsString(trackerDTOConverter.convertToDTO(created)))
                                              .contentType(MediaType.APPLICATION_JSON))
               .andExpect(MockMvcResultMatchers.status().isOk())
               .andExpect(MockMvcResultMatchers.jsonPath("$.id").value(0));
    }
}
