package com.github.mdeluise.tracky.observation.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.mdeluise.tracky.TestEnvironment;
import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
import com.github.mdeluise.tracky.observation.Observation;
import com.github.mdeluise.tracky.observation.ObservationController;
import com.github.mdeluise.tracky.observation.ObservationDTO;
import com.github.mdeluise.tracky.observation.ObservationDTOConverter;
import com.github.mdeluise.tracky.observation.ObservationService;
import com.github.mdeluise.tracky.security.apikey.ApiKeyFilter;
import com.github.mdeluise.tracky.security.apikey.ApiKeyRepository;
import com.github.mdeluise.tracky.security.apikey.ApiKeyService;
import com.github.mdeluise.tracky.security.jwt.JwtTokenFilter;
import com.github.mdeluise.tracky.security.jwt.JwtTokenUtil;
import com.github.mdeluise.tracky.security.jwt.JwtWebUtil;
import com.github.mdeluise.tracky.tracker.Tracker;
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
import java.util.Optional;

@WebMvcTest(ObservationController.class)
@AutoConfigureMockMvc(addFilters = false)
@WithMockUser(username = "admin", roles = "ADMIN")
@Import(TestEnvironment.class)
public class ObservationControllerTest {
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
    private ObservationService observationService;
    @MockBean
    private ObservationDTOConverter observationDTOConverter;
    @MockBean
    private UserService userService;
    @Autowired
    private ObjectMapper objectMapper;
    @Autowired
    private MockMvc mockMvc;


    @Test
    void whenGetObservations_ShouldReturnObservations() throws Exception {
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Observation observation1 = new Observation();
        observation1.setId(1L);
        observation1.setTracker(tracker);
        ObservationDTO observationDTO1 = new ObservationDTO();
        observationDTO1.setId(1L);
        observationDTO1.setTrackerId(0L);
        Observation observation2 = new Observation();
        observation2.setId(2L);
        observation2.setTracker(tracker);
        ObservationDTO observationDTO2 = new ObservationDTO();
        observationDTO2.setId(2L);
        observationDTO2.setTrackerId(0L);
        Mockito.when(observationService.getAll(Optional.of(0L), Optional.empty(), Optional.empty(),
                                               PageRequest.of(0, 10, Sort.Direction.DESC, "instant")
        )).thenReturn(new PageImpl<>(List.of(observation1, observation2)));
        Mockito.when(observationDTOConverter.convertToDTO(observation1)).thenReturn(observationDTO1);
        Mockito.when(observationDTOConverter.convertToDTO(observation2)).thenReturn(observationDTO2);

        mockMvc.perform(MockMvcRequestBuilders.get("/observation?trackerId=0"))
               .andExpect(MockMvcResultMatchers.status().isOk())
               .andExpect(MockMvcResultMatchers.jsonPath("$.numberOfElements").value(2));
    }


    @Test
    void whenDeleteObservation_shouldReturnOk() throws Exception {
        Mockito.doNothing().when(observationService).remove(0L);

        mockMvc.perform(MockMvcRequestBuilders.delete("/observation/0"))
               .andExpect(MockMvcResultMatchers.status().isOk());
    }


    @Test
    void whenDeleteNonExistingObservation_shouldError() throws Exception {
        Mockito.doThrow(ResourceNotFoundException.class).when(observationService).remove(List.of(0L));

        mockMvc.perform(MockMvcRequestBuilders.delete("/observation/0"))
               .andExpect(MockMvcResultMatchers.status().is4xxClientError());
    }


    @Test
    void whenCreateObservation_shouldReturnObservation() throws Exception {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Observation created = new Observation();
        created.setId(0L);
        created.setTracker(tracker);
        created.setUser(user);
        ObservationDTO createdDTO = new ObservationDTO();
        createdDTO.setId(0L);
        createdDTO.setTrackerId(0L);
        Mockito.when(observationService.save(created)).thenReturn(created);
        Mockito.when(observationDTOConverter.convertToDTO(created)).thenReturn(createdDTO);
        Mockito.when(observationDTOConverter.convertFromDTO(createdDTO)).thenReturn(created);
        Mockito.when(userService.get("admin")).thenReturn(user);

        mockMvc.perform(MockMvcRequestBuilders.post("/observation").content(
                                                  objectMapper.writeValueAsString(observationDTOConverter.convertToDTO(created)))
                                              .contentType(MediaType.APPLICATION_JSON))
               .andExpect(MockMvcResultMatchers.status().isOk())
               .andExpect(MockMvcResultMatchers.jsonPath("$.trackerId").value(0L))
               .andExpect(MockMvcResultMatchers.jsonPath("$.id").value(0));
    }
}
