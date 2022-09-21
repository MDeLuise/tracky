package com.github.mdeluise.tracky.observation;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.common.AuthenticatedUserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
import com.github.mdeluise.tracky.exception.UnauthorizedException;
import com.github.mdeluise.tracky.tracker.Tracker;
import com.github.mdeluise.tracky.tracker.TrackerService;
import jakarta.transaction.Transactional;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.Calendar;
import java.util.Date;
import java.util.List;
import java.util.Optional;

@Service
public class ObservationService {
    private final ObservationRepository observationRepository;
    private final TrackerService trackerService;
    private final AuthenticatedUserService authenticatedUserService;


    public ObservationService(ObservationRepository observationRepository, TrackerService trackerService,
                              AuthenticatedUserService authenticatedUserService) {
        this.observationRepository = observationRepository;
        this.trackerService = trackerService;
        this.authenticatedUserService = authenticatedUserService;
    }


    public Page<Observation> getAll(Optional<Long> trackerId, Optional<Date> betweenStart, Optional<Date> betweenEnd,
                                    Pageable pageable) {
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        Calendar calendar = Calendar.getInstance();
        calendar.setTime(new Date());
        calendar.add(Calendar.YEAR, 100);

        if (trackerId.isEmpty()) {
            return observationRepository.findAllByUserAndInstantBetween(
                authenticatedUser, betweenStart.orElse(new Date(0L)), betweenEnd.orElse(calendar.getTime()), pageable);
        }

        Tracker tracker = trackerService.get(trackerId.get());
        if (!authenticatedUser.equals(tracker.getUser())) {
            throw new UnauthorizedException();
        }
        return observationRepository.findAllByTrackerAndInstantBetween(
            tracker, betweenStart.orElse(new Date(0L)), betweenEnd.orElse(calendar.getTime()), pageable);
    }


    @Transactional
    public void remove(Long id) {
        Observation toDelete = observationRepository.findById(id).orElseThrow(() -> new ResourceNotFoundException(id));
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        if (!authenticatedUser.equals(toDelete.getUser())) {
            throw new UnauthorizedException();
        }
        observationRepository.delete(toDelete);
    }


    @Transactional
    public void remove(List<Long> ids) {
        ids.forEach(this::remove);
    }


    @Transactional
    public Observation save(Observation entityToSave) {
        entityToSave.setUser(authenticatedUserService.getAuthenticatedUser());
        entityToSave.getTracker().setLastObservationOn(entityToSave.getInstant());
        return observationRepository.save(entityToSave);
    }


    @Transactional
    public Observation update(Long observationId, Observation updatedEntity) {
        Observation savedObservation = get(observationId);
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        if (!authenticatedUser.equals(savedObservation.getUser())) {
            throw new UnauthorizedException();
        }
        if (!savedObservation.getInstant().equals(updatedEntity.getInstant())) {
            savedObservation.getTracker().setLastObservationOn(updatedEntity.getInstant());
            savedObservation.setInstant(updatedEntity.getInstant());
        }
        savedObservation.setValue(updatedEntity.getValue());
        savedObservation.setNote(updatedEntity.getNote());
        return observationRepository.save(savedObservation);
    }


    public Observation get(Long id) {
        return observationRepository.findById(id).orElseThrow(() -> new ResourceNotFoundException(id));
    }


    public void removeAll() {
        observationRepository.deleteAll();
    }


    public void removeAllInTracker(Long trackerId) {
        Tracker tracker = trackerService.get(trackerId);
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        if (!authenticatedUser.equals(tracker.getUser())) {
            throw new UnauthorizedException();
        }
        observationRepository.deleteAllByTracker(tracker);
    }


    public long count(Long trackerId) {
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        Tracker tracker = trackerService.get(trackerId);
        if (!authenticatedUser.equals(tracker.getUser())) {
            throw new UnauthorizedException();
        }
        return observationRepository.countByTracker(tracker);
    }


    public long count() {
        return observationRepository.countByUser(authenticatedUserService.getAuthenticatedUser());
    }
}
