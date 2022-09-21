package com.github.mdeluise.tracky.tracker;

import com.github.mdeluise.tracky.authentication.User;
import com.github.mdeluise.tracky.common.AuthenticatedUserService;
import com.github.mdeluise.tracky.exception.ResourceNotFoundException;
import com.github.mdeluise.tracky.exception.UnauthorizedException;
import jakarta.transaction.Transactional;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TrackerService {
    private final TrackerRepository trackerRepository;
    private final AuthenticatedUserService authenticatedUserService;


    public TrackerService(TrackerRepository trackerRepository, AuthenticatedUserService authenticatedUserService) {
        this.trackerRepository = trackerRepository;
        this.authenticatedUserService = authenticatedUserService;
    }


    public Page<Tracker> getAll(Pageable pageable) {
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        return trackerRepository.findAllByUser(authenticatedUser, pageable);
    }


    @Transactional
    public void remove(Long id) {
        Tracker toDelete = trackerRepository.findById(id).orElseThrow(() -> new ResourceNotFoundException(id));
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        if (!authenticatedUser.equals(toDelete.getUser())) {
            throw new UnauthorizedException();
        }
        trackerRepository.delete(toDelete);
    }


    @Transactional
    public void remove(List<Long> ids) {
        ids.forEach(this::remove);
    }


    @Transactional
    public Tracker save(Tracker entityToSave) {
        entityToSave.setUser(authenticatedUserService.getAuthenticatedUser());
        return trackerRepository.save(entityToSave);
    }


    public Tracker update(Long trackerId, Tracker tracker) {
        Tracker savedTracker = get(trackerId);
        User authenticatedUser = authenticatedUserService.getAuthenticatedUser();
        if (!authenticatedUser.equals(savedTracker.getUser())) {
            throw new UnauthorizedException();
        }
        savedTracker.setName(tracker.getName());
        savedTracker.setUnit(tracker.getUnit());
        savedTracker.setDescription(tracker.getDescription());
        return trackerRepository.save(savedTracker);
    }


    public void removeAll() {
        trackerRepository.deleteAll();
    }


    public Tracker get(Long trackerId) {
        return trackerRepository.findById(trackerId).orElseThrow(() -> new ResourceNotFoundException(trackerId));
    }


    public Tracker getByName(String name) {
        return trackerRepository.findByName(name).orElseThrow(() -> new ResourceNotFoundException("name", name));
    }


    public long count() {
        return trackerRepository.countByUser(authenticatedUserService.getAuthenticatedUser());
    }
}
