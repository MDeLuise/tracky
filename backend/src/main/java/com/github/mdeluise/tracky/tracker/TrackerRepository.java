package com.github.mdeluise.tracky.tracker;

import com.github.mdeluise.tracky.authentication.User;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface TrackerRepository extends JpaRepository<Tracker, Long> {
    Page<Tracker> findAllByUser(User user, Pageable pageable);

    Optional<Tracker> findByName(String name);

    long countByUser(User user);
}
