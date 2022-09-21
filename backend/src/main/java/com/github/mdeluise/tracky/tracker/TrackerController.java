package com.github.mdeluise.tracky.tracker;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/tracker")
@Tag(name = "Tracker", description = "Endpoints for operations on trackers.")
public class TrackerController {
    private final TrackerService trackerService;
    private final TrackerDTOConverter trackerDtoConverter;


    @Autowired
    public TrackerController(TrackerService trackerService, TrackerDTOConverter trackerDtoConverter) {
        this.trackerService = trackerService;
        this.trackerDtoConverter = trackerDtoConverter;
    }


    @Operation(summary = "Get all the Trackers", description = "Get all the Trackers.")
    @GetMapping
    public ResponseEntity<Page<TrackerDTO>> findAll(@RequestParam(defaultValue = "0", required = false) Integer pageNo,
                                                    @RequestParam(defaultValue = "10", required = false)
                                                    Integer pageSize,
                                                    @RequestParam(defaultValue = "id", required = false) String sortBy,
                                                    @RequestParam(defaultValue = "ASC", required = false)
                                                    Sort.Direction sortDir) {
        Pageable pageable = PageRequest.of(pageNo, pageSize, sortDir, sortBy);
        Page<TrackerDTO> allTrackers = trackerService.getAll(pageable).map(trackerDtoConverter::convertToDTO);
        return ResponseEntity.ok(allTrackers);
    }


    @Operation(summary = "Get a Tracker", description = "Get one Tracker.")
    @GetMapping("/{id}")
    public ResponseEntity<TrackerDTO> get(@PathVariable Long id) {
        TrackerDTO tracker = trackerDtoConverter.convertToDTO(trackerService.get(id));
        return ResponseEntity.ok(tracker);
    }


    @Operation(
        summary = "Delete Trackers", description = "Delete the given Trackers, according to the `ids` parameter."
    )
    @DeleteMapping("/{ids}")
    public void remove(@Parameter(description = "The ID of the Tracker on which to perform the operation") @PathVariable
                       List<Long> ids) {
        trackerService.remove(ids);
    }


    @Operation(
        summary = "Create a new Tracker", description = "Create a new Tracker."
    )
    @PostMapping
    public ResponseEntity<TrackerDTO> save(@RequestBody TrackerDTO entityToSave) {
        TrackerDTO result =
            trackerDtoConverter.convertToDTO(trackerService.save(trackerDtoConverter.convertFromDTO(entityToSave)));
        return ResponseEntity.ok(result);
    }


    @Operation(
        summary = "Update a Tracker", description = "Update a Tracker."
    )
    @PutMapping("{id}")
    public ResponseEntity<TrackerDTO> update(@PathVariable Long id, @RequestBody TrackerDTO entityToSave) {
        TrackerDTO result = trackerDtoConverter.convertToDTO(
            trackerService.update(id, trackerDtoConverter.convertFromDTO(entityToSave)));
        return ResponseEntity.ok(result);
    }


    @Operation(
        summary = "Count the trackers", description = "Count all the inserted trackers."
    )
    @GetMapping("/count")
    public ResponseEntity<Long> count() {
        Long result = trackerService.count();
        return ResponseEntity.ok(result);
    }
}
