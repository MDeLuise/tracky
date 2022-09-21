package com.github.mdeluise.tracky.tracker;

import com.github.mdeluise.tracky.common.AbstractDTOConverter;
import com.github.mdeluise.tracky.observation.Observation;
import com.github.mdeluise.tracky.observation.ObservationDTOConverter;
import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Component;

import java.util.Comparator;

@Component
public class TrackerDTOConverter extends AbstractDTOConverter<Tracker, TrackerDTO> {
    private final ObservationDTOConverter observationDtoConverter;


    public TrackerDTOConverter(ModelMapper modelMapper, ObservationDTOConverter observationDtoConverter) {
        super(modelMapper);
        this.observationDtoConverter = observationDtoConverter;
    }


    @Override
    public Tracker convertFromDTO(TrackerDTO dto) {
        return modelMapper.map(dto, Tracker.class);
    }


    @Override
    public TrackerDTO convertToDTO(Tracker data) {
        TrackerDTO result = modelMapper.map(data, TrackerDTO.class);
        if (data.getObservations().size() > 0) {
            Observation lastObservation = data.getObservations().stream()
                                              .max(Comparator.comparing(Observation::getInstant)).orElse(null);
            result.setLastObservation(observationDtoConverter.convertToDTO(lastObservation));
        }
        return result;
    }
}
