package com.github.mdeluise.tracky.observation;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.format.annotation.DateTimeFormat;
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

import java.util.Date;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/observation")
@Tag(name = "Observations", description = "Endpoints for operations on observations.")
public class ObservationController {
    private final ObservationService observationService;
    private final ObservationDTOConverter observationDtoConverter;


    @Autowired
    public ObservationController(ObservationDTOConverter observationDTOConverter,
                                 ObservationService observationService) {
        this.observationService = observationService;
        this.observationDtoConverter = observationDTOConverter;
    }


    @GetMapping
    @Operation(summary = "Get all the Observations", description = "Get all the Observations.")
    @SuppressWarnings("ParameterNumber") //TODO refactor
    public ResponseEntity<Page<ObservationDTO>> findAll(@RequestParam(required = false) Optional<Long> trackerId,
                                                        @DateTimeFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSSXXX")
                                                        @RequestParam(required = false) Optional<Date> betweenStart,
                                                        @DateTimeFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSSXXX")
                                                        @RequestParam(required = false) Optional<Date> betweenEnd,
                                                        @RequestParam(defaultValue = "0", required = false)
                                                        Integer pageNo,
                                                        @RequestParam(defaultValue = "10", required = false)
                                                        Integer pageSize,
                                                        @RequestParam(defaultValue = "instant", required = false)
                                                        String sortBy,
                                                        @RequestParam(defaultValue = "DESC", required = false)
                                                        Sort.Direction sortDir) {
        Pageable pageable = PageRequest.of(pageNo, pageSize, sortDir, sortBy);
        Page<ObservationDTO> result = observationService.getAll(trackerId, betweenStart, betweenEnd, pageable)
                                                        .map(observationDtoConverter::convertToDTO);
        return ResponseEntity.ok(result);
    }


    @Operation(
        summary = "Delete Observations",
        description = "Delete the given Observations, according to the `ids` parameter."
    )
    @DeleteMapping("/{ids}")
    public void remove(
        @Parameter(description = "The ID of the Observation on which to perform the operation") @PathVariable
        List<Long> ids) {
        observationService.remove(ids);
    }


    @Operation(
        summary = "Create a new Observation", description = "Create a new Observation."
    )
    @PostMapping
    public ResponseEntity<ObservationDTO> save(@RequestBody ObservationDTO entityToSave) {
        ObservationDTO result = observationDtoConverter.convertToDTO(
            observationService.save(observationDtoConverter.convertFromDTO(entityToSave)));
        return ResponseEntity.ok(result);
    }


    @Operation(
        summary = "Update an Observation", description = "Update an Observation."
    )
    @PutMapping("/{id}")
    public ResponseEntity<ObservationDTO> update(@PathVariable Long id, @RequestBody ObservationDTO entityToSave) {
        ObservationDTO result = observationDtoConverter.convertToDTO(
            observationService.update(id, observationDtoConverter.convertFromDTO(entityToSave)));
        return ResponseEntity.ok(result);
    }


    @Operation(
        summary = "Count the observations",
        description = "Count all the inserted observations, or all related to the Tracker with specified `trackerId`."
    )
    @GetMapping("/count")
    public ResponseEntity<Long> count(@RequestParam(required = false) Long trackerId) {
        Long result = trackerId == null ? observationService.count() : observationService.count(trackerId);
        return ResponseEntity.ok(result);
    }
}
