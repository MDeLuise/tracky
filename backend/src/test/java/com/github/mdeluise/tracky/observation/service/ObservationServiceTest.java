package com.github.mdeluise.tracky.observation.service;


import com.github.mdeluise.tracky.TestEnvironment;
import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import com.github.mdeluise.tracky.common.AuthenticatedUserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
import com.github.mdeluise.tracky.observation.Observation;
import com.github.mdeluise.tracky.observation.ObservationRepository;
import com.github.mdeluise.tracky.observation.ObservationService;
import com.github.mdeluise.tracky.tracker.Tracker;
import com.github.mdeluise.tracky.tracker.TrackerRepository;
import com.github.mdeluise.tracky.tracker.TrackerService;
import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.springframework.context.annotation.Import;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.Pageable;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.util.Date;
import java.util.List;
import java.util.Optional;

@ExtendWith(SpringExtension.class)
@WithMockUser(username = "user")
@Import(TestEnvironment.class)
public class ObservationServiceTest {
    @Mock
    ObservationRepository observationRepository;
    @Mock
    TrackerRepository trackerRepository;
    @Mock
    TrackerService trackerService;
    @Mock
    UserService userService;
    @Mock
    AuthenticatedUserService authenticatedUserService;
    @InjectMocks
    ObservationService observationService;


    @Test
    void whenSaveObservation_thenReturnObservation() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Observation toSave = new Observation();
        toSave.setId(0L);
        toSave.setTracker(tracker);
        toSave.setUser(user);
        Mockito.when(observationRepository.save(toSave)).thenReturn(toSave);
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);
        Mockito.when(trackerRepository.save(tracker)).thenReturn(tracker);

        Assertions.assertThat(observationService.save(toSave)).isSameAs(toSave);
    }


    @Test
    void whenGetAllObservations_thenReturnAllObservations() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        tracker.setUser(user);
        Observation toGet1 = new Observation();
        toGet1.setId(0L);
        toGet1.setTracker(tracker);
        Observation toGet2 = new Observation();
        toGet2.setId(1L);
        toGet2.setTracker(tracker);

        List<Observation> allObservations = List.of(toGet1, toGet2);
        Date now = new Date();

        Mockito.when(
                   observationRepository.findAllByTrackerAndInstantBetween(tracker, new Date(0L), now,
                                                                           Pageable.unpaged()))
               .thenReturn(new PageImpl<>(allObservations));
        Mockito.when(trackerService.get(0L)).thenReturn(tracker);
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);

        Page<Observation> result =
            observationService.getAll(Optional.of(tracker.getId()), Optional.empty(), Optional.of(now),
                                      Pageable.unpaged()
            );
        Assertions.assertThat(result.getContent()).containsAll(allObservations);
        Assertions.assertThat(allObservations.containsAll(result.getContent()));
    }


    @Test
    void givenObservation_whenDeleteObservation_thenDeleteObservation() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Observation observation = new Observation();
        observation.setTracker(tracker);
        observation.setUser(user);
        Mockito.when(observationRepository.findById(0L)).thenReturn(Optional.of(observation));
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);
        observationService.remove(0L);
        Mockito.verify(observationRepository, Mockito.times(1)).delete(observation);
    }


    @Test
    void whenDeleteNonExistingObservation_thenError() {
        User user = new User();
        user.setId(0L);
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);
        Mockito.when(observationRepository.findById(0L)).thenReturn(Optional.empty());

        Assertions.assertThatThrownBy(() -> observationService.remove(0L))
                  .isInstanceOf(ResourceNotFoundException.class);
    }
}
