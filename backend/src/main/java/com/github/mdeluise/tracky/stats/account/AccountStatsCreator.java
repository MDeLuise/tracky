package com.github.mdeluise.tracky.stats.account;

import com.github.mdeluise.tracky.observation.ObservationService;
import com.github.mdeluise.tracky.tracker.TrackerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AccountStatsCreator {
    private final ObservationService observationService;
    private final TrackerService trackerService;


    @Autowired
    public AccountStatsCreator(ObservationService observationService, TrackerService trackerService) {
        this.observationService = observationService;
        this.trackerService = trackerService;
    }


    public AccountStats getData() {
        AccountStats dataStats = new AccountStats();
        dataStats.setTotalObservations(observationService.count());
        dataStats.setTotalTrackers(trackerService.count());
        return dataStats;
    }
}
