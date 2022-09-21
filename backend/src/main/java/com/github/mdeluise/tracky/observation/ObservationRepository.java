package com.github.mdeluise.tracky.observation;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.tracker.Tracker;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Date;
import java.util.Optional;

public interface ObservationRepository extends JpaRepository<Observation, Long> {
    void deleteAllByTracker(Tracker tracker);

    Page<Observation> findAllByTrackerAndInstantBetween(Tracker tracker, Date start, Date end, Pageable pageable);

    long countByTracker(Tracker tracker);

    long countByUser(User authenticatedUser);

    Optional<Observation> findFirst1ByUserOrderByInstantDesc(User user);

    Optional<Observation> findFirst1ByTrackerOrderByInstantDesc(Tracker tracker);

    Page<Observation> findAllByUserAndInstantBetween(User authenticatedUser, Date orElse, Date orElse1,
                                                     Pageable pageable);
}
