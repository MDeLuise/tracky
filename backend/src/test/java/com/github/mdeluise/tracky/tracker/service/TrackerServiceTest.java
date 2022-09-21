package com.github.mdeluise.tracky.tracker.service;


import com.github.mdeluise.tracky.TestEnvironment;
import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.authentication.UserService;
import com.github.mdeluise.tracky.common.AuthenticatedUserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
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
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.Pageable;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.util.List;
import java.util.Optional;

@ExtendWith(SpringExtension.class)
@WithMockUser(username = "user")
@Import(TestEnvironment.class)
public class TrackerServiceTest {
    @Mock
    TrackerRepository trackerRepository;
    @Mock
    UserService userService;
    @Mock
    AuthenticatedUserService authenticatedUserService;
    @InjectMocks
    TrackerService trackerService;


    @Test
    void whenSaveTracker_thenReturnTracker() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        Tracker toSave = new Tracker();
        toSave.setId(0L);
        toSave.setUser(user);
        Mockito.when(trackerRepository.save(toSave)).thenReturn(toSave);
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);

        Assertions.assertThat(trackerService.save(toSave)).isSameAs(toSave);
    }


    @Test
    void whenGetAllTrackers_thenReturnAllTrackers() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        tracker.setUser(user);
        Tracker toGet1 = new Tracker();
        toGet1.setId(0L);
        Tracker toGet2 = new Tracker();
        toGet2.setId(1L);

        List<Tracker> allTrackers = List.of(toGet1, toGet2);
        Mockito.when(trackerRepository.findAllByUser(user, Pageable.unpaged())).thenReturn(new PageImpl<>(allTrackers));
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);

        Assertions.assertThat(trackerService.getAll(Pageable.unpaged()).getContent()).containsAll(allTrackers);
        Assertions.assertThat(allTrackers).containsAll(trackerService.getAll(Pageable.unpaged()).getContent());
    }


    @Test
    void givenTracker_whenDeleteTracker_thenDeleteTracker() {
        User user = new User();
        user.setId(0L);
        Tracker tracker = new Tracker();
        tracker.setId(0L);
        tracker.setUser(user);
        Mockito.when(trackerRepository.findById(0L)).thenReturn(Optional.of(tracker));
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);
        trackerService.remove(0L);
        Mockito.verify(trackerRepository, Mockito.times(1)).delete(tracker);
    }


    @Test
    void whenDeleteNonExistingTracker_thenError() {
        User user = new User();
        user.setId(0L);
        Mockito.when(userService.get("user")).thenReturn(user);
        Mockito.when(authenticatedUserService.getAuthenticatedUser()).thenReturn(user);
        Mockito.when(trackerRepository.findById(0L)).thenReturn(Optional.empty());

        Assertions.assertThatThrownBy(() -> trackerService.remove(0L)).isInstanceOf(ResourceNotFoundException.class);
    }
}
